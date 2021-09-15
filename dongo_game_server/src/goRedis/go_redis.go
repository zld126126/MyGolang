package goRedis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/sirupsen/logrus"
	"log"
)

func RedisTest() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		logrus.WithField("err", fmt.Sprintf("%+v", err)).Errorln(`Connect to redis error`)
		log.Fatal(err)
	}
	defer c.Close()

	err = setGetRedis(err, c)

	err = deleteRedis(err, c)

	redisJson(err, c)
}

func setGetRedis(err error, c redis.Conn) error {
	_, err = c.Do("SET", "mykey", "123456")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	is_key_exit, err := redis.Bool(c.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit)
	}
	_, err = c.Do("SET", "mykey1", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	is_key_exit2, err := redis.Bool(c.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit2)
	}
	return err
}

func deleteRedis(err error, c redis.Conn) error {
	_, err = c.Do("DEL", "mykey1")
	if err != nil {
		fmt.Println("redis delelte failed:", err)
	}
	is_key_exit3, err := redis.Bool(c.Do("EXISTS", "mykey1"))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("exists or not: %v \n", is_key_exit3)
	}
	return err
}

func redisJson(err error, c redis.Conn) {
	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)
	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}
	var imapGet map[string]string
	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}
	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}
