package dao

import (
	"go-scraper/src/dao/database"
	"go-scraper/src/dao/repository"
	"go-scraper/src/errorcode"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao struct {
	alias string
}

/*
	new single dao
*/
func newSingleDao(alias string) (d repository.Dao) {
	d = &Dao{
		alias: alias,
	}

	return
}

func (c *Dao) Insert(t string, d []interface{}, o ...*options.InsertManyOptions) ([]string, int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	_ids := make([]string, 0)

	r, err := co.InsertMany(ctx, d, o...)
	if err != nil {
		return _ids, errorcode.DatabaseInsertFailed, err
	}

	for _, v := range r.InsertedIDs {
		_ids = append(_ids, v.(primitive.ObjectID).Hex())
	}

	return _ids, errorcode.None, nil
}

func (c *Dao) InsertOne(t string, d interface{}, o ...*options.InsertOneOptions) (string, int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	r, err := co.InsertOne(ctx, d, o...)
	if err != nil {
		isDup := false
		if we, ok := err.(mongo.WriteException); ok {
			for _, v := range we.WriteErrors {
				if v.Code == 11000 {
					isDup = true
					break
				}
			}
		}

		if isDup {
			return "", errorcode.DatabaseDuplicateKey, err
		} else {
			return "", errorcode.DatabaseInsertFailed, err
		}
	}

	_id := r.InsertedID.(primitive.ObjectID).Hex()

	return _id, errorcode.None, nil
}

func (c *Dao) Query(t string, r interface{}, f interface{}, o ...*options.FindOptions) (int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	rr, err := co.Find(ctx, f, o...)
	if err != nil {
		return errorcode.DatabaseQueryFailed, err
	}
	defer rr.Close(ctx)

	if r == nil {
		r = bson.M{}
	}

	if err = rr.All(ctx, r); err != nil {
		return errorcode.DatabaseParseDataError, err
	}

	return errorcode.None, nil
}

func (c *Dao) QueryOne(t string, r interface{}, f interface{}, o ...*options.FindOneOptions) (int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	rr := co.FindOne(ctx, f, o...)
	if err := rr.Err(); err != nil {
		return errorcode.DatabaseQueryFailed, err
	}

	if r == nil {
		r = bson.M{}
	}

	if err := rr.Decode(r); err != nil {
		return errorcode.DatabaseParseDataError, err
	}

	return errorcode.None, nil
}

func (c *Dao) QueryAndUpdate(t string, r interface{}, f interface{}, u interface{}, o ...*options.FindOneAndUpdateOptions) (int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	rr := co.FindOneAndUpdate(ctx, f, u, o...)
	if err := rr.Err(); err != nil {
		return errorcode.DatabaseQueryFailed, err
	}

	if r == nil {
		r = bson.M{}
	}

	if err := rr.Decode(r); err != nil {
		return errorcode.DatabaseParseDataError, err
	}

	return errorcode.None, nil
}

func (c *Dao) QueryCount(t string, f interface{}, o ...*options.CountOptions) (int64, int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	count, err := co.CountDocuments(ctx, f, o...)
	if err != nil {
		return 0, errorcode.DatabaseQueryFailed, err
	}

	return count, errorcode.None, nil
}

func (c *Dao) Update(t string, f interface{}, u interface{}, o ...*options.UpdateOptions) (int64, int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	r, err := co.UpdateMany(ctx, f, u, o...)
	if err != nil {
		return 0, errorcode.DatabaseUpdateFailed, err
	}
	count := r.ModifiedCount
	return count, errorcode.None, nil
}

func (c *Dao) Delete(t string, f interface{}, o ...*options.DeleteOptions) (int64, int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	r, err := co.DeleteMany(ctx, f, o...)
	if err != nil {
		return 0, errorcode.DatabaseDeleteFailed, err
	}
	count := r.DeletedCount
	return count, errorcode.None, nil
}

func (c *Dao) Aggregate(t string, r interface{}, p interface{}, o ...*options.AggregateOptions) (int, error) {
	conn := database.GetInstance()
	ctx, cancel, co := conn.Open(c.alias, t)
	defer cancel()

	rr, err := co.Aggregate(ctx, p, o...)
	if err != nil {
		return errorcode.DatabaseQueryFailed, err
	}
	defer rr.Close(ctx)

	if r == nil {
		r = bson.M{}
	}

	if err = rr.All(ctx, r); err != nil {
		return errorcode.DatabaseParseDataError, err
	}

	return errorcode.None, nil
}
