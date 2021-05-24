package core

import (
	"github.com/erDong01/micro-kit/cache/I"
	"github.com/erDong01/micro-kit/config"
	"github.com/go-redis/redis/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
)

type info struct {
	port    int
	env     string
	Name    string
	version string
}
type Core struct {
	Db          *gorm.DB
	Transaction *gorm.DB
	Redis       *redis.Client
	Config      *config.Config
	Cache       *I.ICache
	Info        info
	opts        []Option
	once        sync.Once
	Mongo       *mongo.Client
}

var (
	c    *Core
	once sync.Once
)

func New() *Core {
	once.Do(func() {
		c = &Core{}
	})
	return c

}

func Set(newCore *Core) {
	once.Do(func() {
		c = newCore
	})
}

func Close() {
	if c.Redis != nil {
		defer c.Redis.Close()
	}

}
