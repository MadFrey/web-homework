package main

import (
"fmt"
"github.com/gomodule/redigo/redis"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go subMessage("redis",&wg)
	pubMessage("redis","hello")
	wg.Wait()
}
func initRedis() (conn redis.Conn, err error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	return c, err
}
func subMessage(channel string,wg *sync.WaitGroup)  {
	defer wg.Done()
	conn,err := initRedis()
	if err != nil {
		panic(err)
		return
	}
	client := redis.PubSubConn{Conn:conn}
	err = client.Subscribe(channel)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("wait..")
	for {
		switch v := client.Receive().(type){
		case redis.Message:
			fmt.Println("Message", v.Channel, string(v.Data))
		case redis.Subscription:
			fmt.Println("Subscription", v.Channel, v.Kind, v.Count)
		}
	}
}

func pubMessage(channel, msg string) {
	conn,err := initRedis()
	if err != nil {
		panic(err)
		return
	}
	_, err = conn.Do("Publish", channel, msg)
	if err != nil {
		log.Println(err)
		return
	}
}
