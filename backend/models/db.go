package models

import (
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
	"time"
)

var (
	DB        *sqlx.DB
	RedisConn *redis.Pool
	ESConn    *elasticsearch.Client
	err       error
)

func init() {
	// setup mysql
	DB = sqlx.MustConnect(setting.DBInfo.Mode, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		setting.DBInfo.User,
		setting.DBInfo.Password,
		setting.DBInfo.Host,
		setting.DBInfo.Port,
		setting.DBInfo.DBName))
	
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	
	// setup redis
	RedisConn = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			conn, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", setting.RedisInfo.Host, setting.RedisInfo.Port))
			if err != nil {
				return nil, err
			}
			
			if setting.RedisInfo.Password != "" {
				if _, err = conn.Do("AUTH", setting.RedisInfo.Password); err != nil {
					_ = conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     setting.RedisInfo.MaxIdle,
		MaxActive:   setting.RedisInfo.MaxActive,
		IdleTimeout: time.Second * time.Duration(setting.RedisInfo.IdleTimeout),
	}
	
	// setup es
	esAddr := fmt.Sprintf("http://%s:%s", setting.ESInfo.Host, setting.ESInfo.Port)
	ESConn, _ = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			esAddr,
		},
	})
}
