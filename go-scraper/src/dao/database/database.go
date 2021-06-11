package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go-scraper/src/errorcode"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Account        string `mapstructure:"account"`         // 登入帳號
	Password       string `mapstructure:"password"`        // 登入密碼
	Address        string `mapstructure:"address"`         // MongoDB位址
	Port           int    `mapstructure:"port"`            // MongoDB埠號
	MaxPoolSize    uint64 `mapstructure:"max_pool_size"`   // MaxPoolSize 連線時最大數量
	RequestTimeout int64  `mapstructure:"request_timeout"` // RequestTimeout 請求結束時間
}

/*
	database initialize
*/
func Initialize() (int, error) {
	d := &Database{
		Account:        os.Getenv("MONGO_ACCOUNT"),
		Password:       os.Getenv("MONGO_PASSWORD"),
		Address:        os.Getenv("MONGO_ADDR"),
		MaxPoolSize:    1024,
		Port:           27017,
		RequestTimeout: 30,
	}

	return d.NewDatabase()
}

/*
	new database
*/
func (d *Database) NewDatabase() (int, error) {
	conn := GetInstance()

	conn.SetCreate(func(d *Database) (c *mongo.Client, code int, err error) {
		account, password, address, port, maxPoolSize, requestTimeout := d.Account, d.Password, d.Address, d.Port, d.MaxPoolSize, d.RequestTimeout
		if address == "" {
			d.Address = conn.database.Address
		}
		if port == 0 {
			d.Port = conn.database.Port
		}
		if maxPoolSize == 0 {
			d.MaxPoolSize = conn.database.MaxPoolSize
		}
		if requestTimeout == 0 {
			d.RequestTimeout = conn.database.RequestTimeout
		}

		// set client options
		opts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", d.Address, d.Port))
		if account != "" && password != "" {
			auth := options.Credential{
				AuthSource:  "admin",
				Username:    d.Account,
				Password:    d.Password,
				PasswordSet: true,
			}
			opts.SetAuth(auth)
		}

		// set max pool size
		opts.SetMaxPoolSize(maxPoolSize)

		// new client
		c, err = mongo.NewClient(opts)
		if err != nil {
			code = errorcode.DatabaseConnectClientError
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(requestTimeout)*time.Second)
		defer cancel()
		if err = c.Connect(ctx); err != nil {
			code = errorcode.DatabaseConnectFailed
			return
		}

		conn.SetClient(c)
		conn.SetDatabase(d)

		// check connection
		if code, err = conn.Ping(ctx); err != nil {
			return
		}

		return
	})

	conn.SetDestroy(func(m *mongo.Client) (int, error) {
		return conn.Close()
	})

	return conn.initialize(d)
}
