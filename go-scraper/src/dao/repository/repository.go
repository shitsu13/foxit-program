package repository

import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Dao interface {
	Insert(t string, d []interface{}, o ...*options.InsertManyOptions) ([]string, int, error)
	InsertOne(t string, d interface{}, o ...*options.InsertOneOptions) (string, int, error)
	Query(t string, r interface{}, f interface{}, o ...*options.FindOptions) (int, error)
	QueryOne(t string, r interface{}, f interface{}, o ...*options.FindOneOptions) (int, error)
	QueryAndUpdate(t string, r interface{}, f interface{}, u interface{}, o ...*options.FindOneAndUpdateOptions) (int, error)
	QueryCount(t string, f interface{}, o ...*options.CountOptions) (int64, int, error)
	Update(t string, f interface{}, u interface{}, o ...*options.UpdateOptions) (int64, int, error)
	Delete(t string, f interface{}, o ...*options.DeleteOptions) (int64, int, error)
	Aggregate(t string, r interface{}, p interface{}, o ...*options.AggregateOptions) (int, error)
}
