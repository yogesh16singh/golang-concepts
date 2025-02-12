package main

import (
	"log"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
)

var redisPool = &redis.Pool{
	MaxIdle:   5,
	MaxActive: 5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	},
}

var enqueuer = work.NewEnqueuer("demo_app", redisPool)

func main() {
	_, err := enqueuer.Enqueue("email",
		work.Q{"email": "a@b.com", "subject": "hello", "body": "world"})
	if err != nil {
		log.Fatal(err)
	}
}
