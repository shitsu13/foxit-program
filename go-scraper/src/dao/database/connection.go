package database

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go-scraper/src/errorcode"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	conn *connection
	once sync.Once
)

type connection struct {
	client   *mongo.Client
	database *Database
	mu       *sync.Mutex
	create   func(d *Database) (*mongo.Client, int, error)
	destroy  func(m *mongo.Client) (int, error)
}

func (c *connection) Client() *mongo.Client {
	return c.client
}

func (c *connection) SetClient(client *mongo.Client) {
	c.client = client
}

func (c *connection) Database() *Database {
	return c.database
}

/*
	set database
*/
func (c *connection) SetDatabase(database *Database) {
	c.database = database
}

/*
	set connection create callback
*/
func (c *connection) SetCreate(create func(d *Database) (*mongo.Client, int, error)) {
	c.create = create
}

/*
	set connection destroy callback
*/
func (c *connection) SetDestroy(destroy func(m *mongo.Client) (int, error)) {
	c.destroy = destroy
}

/*
	use singleton pattern to get instance
*/
func GetInstance() *connection {
	once.Do(func() {
		conn = newConnection()
	})
	return conn
}

/*
	new connection
*/
func newConnection() (c *connection) {
	c = &connection{
		client: &mongo.Client{},
		database: &Database{
			Account:        "",
			Password:       "",
			Address:        "",
			Port:           27017,
			MaxPoolSize:    1024,
			RequestTimeout: 30,
		},
		mu:      &sync.Mutex{},
		create:  func(d *Database) (client *mongo.Client, i int, e error) { return &mongo.Client{}, 0, nil },
		destroy: func(m *mongo.Client) (i int, e error) { return 0, nil },
	}

	return
}

/*
	connection initialize
*/
func (c *connection) initialize(d *Database) (code int, err error) {
	o, code, err := c.create(d)
	if err != nil {
		return
	}

	c.client = o
	return
}

/*
	open connection
*/
func (c *connection) Open(db, name string) (context.Context, context.CancelFunc, *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.database.RequestTimeout)*time.Second)
	co := c.client.Database(db).Collection(name)
	return ctx, cancel, co
}

/*
	check connection alive
*/
func (c *connection) Ping(ctx context.Context) (code int, err error) {
	if err = c.client.Ping(ctx, readpref.Primary()); err != nil {
		code = errorcode.DatabaseConnectPoolFailed
		return
	}

	return
}

/*
	connection destroy
*/
func (c *connection) Destroy() (int, error) {
	return c.destroy(c.client)
}

/*
	close database
*/
func (c *connection) Close() (code int, err error) {
	if c.client == nil {
		code = errorcode.DatabaseConnectClientError
		err = fmt.Errorf("mongo client not found")
		return
	}

	sync := &sync.WaitGroup{}
	sync.Add(1)
	go func() {
		defer sync.Done()

		client := c.client
		for client.NumberSessionsInProgress() > 0 {
			time.Sleep(1 * time.Second)
		}

		if err = c.client.Disconnect(context.Background()); err != nil {
			code = errorcode.DatabaseConnectFailed
		}
	}()

	sync.Wait()
	return
}
