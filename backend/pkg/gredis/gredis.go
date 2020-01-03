package gredis

import (
	"github.com/lawtech0902/go_gin_blog/backend/models"
	"github.com/lawtech0902/go_gin_blog/backend/pkg/setting"
)

type Options struct {
	Timeout bool // 是否设置过期时间
}

var defaultOptions = Options{Timeout: false}

type Option func(*Options)

func NewOptions(opts ...Option) Options {
	// 初始化默认值
	opt := defaultOptions
	
	for _, o := range opts {
		o(&opt) // 依次调用opts函数列表中的函数，为服务选项（opt变量）赋值
	}
	
	return opt
}

func SetTimeout(timeout bool) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

// crud实现
func SetKey(key string, value interface{}, opts ...Option) error {
	conn := models.RedisConn.Get()
	defer conn.Close()
	
	options := NewOptions(opts...)
	if options.Timeout {
		_, err := conn.Do("SET", key, value, "EX", setting.RedisInfo.CacheTime)
		return err
	}
	_, err := conn.Do("SET", key, value)
	return err
}

func GetKey(key string) (data interface{}, err error) {
	conn := models.RedisConn.Get()
	defer conn.Close()
	
	data, err = conn.Do("GET", key)
	return
}

func DelKey(key string) error {
	conn := models.RedisConn.Get()
	defer conn.Close()
	
	_, err := conn.Do("DEL", key)
	return err
}

func INCRKey(key string) error {
	conn := models.RedisConn.Get()
	defer conn.Close()
	
	_, err := conn.Do("INCR", key)
	return err
}


