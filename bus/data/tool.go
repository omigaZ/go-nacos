package data

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"go-nacos/bus/config"
	"go-nacos/go-common/utils"
	"time"
	"xorm.io/xorm"
)

var (
	cache *redis.Client
	db *sql.DB
	xdb *xorm.Engine
)

func CacheOp() *redis.Client {
	if cache == nil {
		var cfg = config.Redis
		cache = redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr: cfg.Addr(),
			Password: cfg.Password,
			DB: cfg.Index,
			MaxRetries: 1,
			PoolSize: 1024,
			IdleTimeout: -1,
			DialTimeout: 100 * time.Microsecond,
		})
		if cache == nil {
			panic("redis connection unestablished")
		}
		fmt.Printf("%s redis connected ...(data)\n", utils.TimestampString())
		fmt.Printf("\taddr  : %s\n", cfg.Addr())
		fmt.Printf("\tindex : %d\n", cfg.Index)
	}
	return cache
}

func DbOp() *sql.DB {
	if db == nil {
		var err error
		var cfg = config.Mysql
		db, err = sql.Open(cfg.Flavor,cfg.DSN())
		if err != nil || db == nil {
			panic("mysql connection unestablished")
		}
		fmt.Printf("%s mysql connected ...\n", utils.TimestampString())
		fmt.Printf("\taddr  : %s\n", config.Mysql.Addr())
		fmt.Printf("\tdb    : %s\n", config.Mysql.Schema)

	}
	return db
}

func XormOp() *xorm.Engine {
	if xdb == nil {
		var err error
		var cfg = config.Mysql
		xdb, err = xorm.NewEngine(cfg.Flavor,cfg.DSN())
		if err != nil || xdb == nil {
			panic("mysql connection unestablished")
		}
		fmt.Printf("%s mysql connected ...\n", utils.TimestampString())
		fmt.Printf("\taddr  : %s\n", config.Mysql.Addr())
		fmt.Printf("\tdb    : %s\n", config.Mysql.Schema)
	}
	return xdb
}
