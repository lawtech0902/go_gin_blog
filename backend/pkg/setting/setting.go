package setting

import (
	"github.com/spf13/viper"
	"time"
)

/*
配置文件解析
*/

type App struct {
	TimeFormat     string `json:"time_format"`
	JwtSecret      string `json:"jwt_secret"`
	TokenTimeout   int64  `json:"token_timeout"`
	RootBasePath   string `json:"root_base_path"`
	LogBasePath    string `json:"log_base_path"`
	UploadBasePath string `json:"upload_base_path"`
	ImageRelPath   string `json:"image_rel_path"`
	AvatarRelPath  string `json:"avatar_rel_path"`
	ApiBaseUrl     string `json:"api_base_url"`
}

type Server struct {
	RunMode      string        `json:"run_mode"`
	ServerAddr   string        `json:"server_addr"`
	ReadTimeout  time.Duration `json:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout"`
}

type DB struct {
	Mode     string `json:"mode"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type Redis struct {
	Host        string `json:"host"`
	Port        string `json:"port"`
	Password    string `json:"password"`
	DB          int    `json:"db"`
	CacheTime   int    `json:"cache_time"`
	MaxIdle     int    `json:"max_idle"`
	MaxActive   int    `json:"max_active"`
	IdleTimeout int    `json:"idle_timeout"`
}

type ES struct {
	Host  string `json:"host"`
	Port  string `json:"port"`
	Index string `json:"index"`
}

var (
	AppInfo    *App
	ServerInfo *Server
	DBInfo     *DB
	RedisInfo  *Redis
	ESInfo     *ES
)

func init() {
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	
	AppInfo = &App{
		TimeFormat:     viper.GetString("app.timeFormat"),
		JwtSecret:      viper.GetString("app.jwtSecret"),
		TokenTimeout:   viper.GetInt64("app.tokenTimeout"),
		RootBasePath:   viper.GetString("app.rootBasePath"),
		LogBasePath:    viper.GetString("app.logBasePath"),
		UploadBasePath: viper.GetString("app.uploadBasePath"),
		ImageRelPath:   viper.GetString("app.imageRelPath"),
		AvatarRelPath:  viper.GetString("app.avatarRelPath"),
		ApiBaseUrl:     viper.GetString("app.apiBaseUrl"),
	}
	
	ServerInfo = &Server{
		RunMode:      viper.GetString("server.runMode"),
		ServerAddr:   viper.GetString("server.serverAddr"),
		ReadTimeout:  time.Duration(viper.GetInt64("server.readTimeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt64("server.writeTimeout")) * time.Second,
	}
	
	DBInfo = &DB{
		Mode:     viper.GetString("db.mode"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbName"),
	}
	
	RedisInfo = &Redis{
		Host:        viper.GetString("redis.host"),
		Port:        viper.GetString("redis.port"),
		Password:    viper.GetString("redis.password"),
		DB:          viper.GetInt("redis.db"),
		CacheTime:   viper.GetInt("redis.cacheTime"),
		MaxIdle:     viper.GetInt("redis.maxIdle"),
		MaxActive:   viper.GetInt("redis.maxActive"),
		IdleTimeout: viper.GetInt("redis.idleTimeout"),
	}
	
	ESInfo = &ES{
		Host:  viper.GetString("es.host"),
		Port:  viper.GetString("es.port"),
		Index: viper.GetString("es.index"),
	}
}
