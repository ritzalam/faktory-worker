package main

import (
  "fmt"

  worker "github.com/contribsys/faktory_worker_go"

  client "faktory-worker/db"

  "github.com/go-redis/redis"

)





func someFunc(ctx worker.Context, args ...interface{}) error {
  fmt.Println("Working on job", ctx.Jid())

  connStru := client.GetSessionCache()
  conn := connStru.Conn

  err := conn.Set("key", "value", 0).Err()
  if err != nil {
    panic(err)
  }

  val, err := connStru.Conn.Get("key").Result()
  if err != nil {
    panic(err)
  }
  fmt.Println("key", val)

  val2, err := connStru.Conn.Get("key2").Result()
  if err == redis.Nil {
    fmt.Println("key2 does not exist")
  } else if err != nil {
    panic(err)
  } else {
    fmt.Println("key2", val2)
  }

  return nil
}

func main() {
  connStru := client.GetSessionCache()
  if connStru == nil {
    fmt.Println("redis conn is NIL")
  }
  //conn := *connStru.Conn  
  //pong, err := conn.Ping().Result()
  //fmt.Println(pong, err)
  // Output: PONG <nil>

  mgr := worker.NewManager()

  // register job types and the function to execute them
  mgr.Register("SomeJob", someFunc)
  //mgr.Register("AnotherJob", anotherFunc)

  // use up to N goroutines to execute jobs
  mgr.Concurrency = 20

  // pull jobs from these queues, in this order of precedence
  mgr.ProcessStrictPriorityQueues("critical", "default", "bulk")

  // alternatively you can use weights to avoid starvation
  //mgr.ProcessWeightedPriorityQueues(map[string]int{"critical":3, "default":2, "bulk":1})

  // Start processing jobs, this method does not return
  mgr.Run()
}
