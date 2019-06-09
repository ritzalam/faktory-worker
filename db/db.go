package db

import (
     "fmt"

	"sync"

    redis "github.com/go-redis/redis"
)

type cache struct {
    Conn *redis.Client
}

var singleCache *cache
var initOnce sync.Once

func GetSessionCache() *cache {
    initOnce.Do(func() {
        fmt.Println("initing singleCache")
        singleCache2 := &cache{}
        singleCache2.Conn =   redis.NewClient(&redis.Options{
    		Addr:     "localhost:6379",
    		Password: "", // no password set
    		DB:       0,  // use default DB
  		})
        pong, err := singleCache2.Conn.Ping().Result()
        fmt.Println(pong, err)

        if singleCache2 == nil {
            fmt.Println("INSIDE: singleCache is NIL")
        } else {
            fmt.Println("INSIDE: singleCache is NOT NIL")
        }

        singleCache = singleCache2
    })

    if singleCache == nil {
        fmt.Println("singleCache is NIL")
    }
    return singleCache
}
