package controllers

import (
	"encoding/json"
	"hi-redis/models"
	"log"

	"github.com/go-redis/redis"
)

type Data struct {
	Type string
	Val  string
}

func (c *MainController) Data() {
	client := redis.NewClient(&redis.Options{
		Addr:     models.AppConfig.DBIp + ":" + models.AppConfig.DBPort,
		Password: models.AppConfig.DBAuth,
		DB:       0, // use default DB
	})

	var data = make([]Data, 0)
	keys, _, err := client.Scan(0, "*", 3000).Result()
	if err != nil {
		panic(err)
	}

	for _, v := range keys {
		t, _ := client.Type(v).Result()
		d := Data{Type: t, Val: v}
		data = append(data, d)
	}

	json, _ := json.Marshal(data)
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	c.Ctx.WriteString(string(json))
}

func (c *MainController) Key() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	key := c.GetString("key")
	if key == "" {
		c.Ctx.WriteString("key is empty")
		return
	}
	log.Println(key)

	client := redis.NewClient(&redis.Options{
		Addr:     models.AppConfig.DBIp + ":" + models.AppConfig.DBPort,
		Password: models.AppConfig.DBAuth,
		DB:       0, // use default DB
	})

	var ret []byte
	t, _ := client.Type(key).Result()
	switch t {
	case "string":
		result, err := client.Get(key).Result()
		if err != nil {
			log.Println(err)
		}
		ret, _ = json.Marshal(result)

	case "hash":
		result, err := client.HGetAll(key).Result()
		if err != nil {
			log.Println(err)
		}
		ret, _ = json.Marshal(result)
	case "set":
		result, err := client.SMembers(key).Result()
		if err != nil {
			log.Println(err)
		}
		ret, _ = json.Marshal(result)
	case "list":
		var ret_tmp []string
		_len, err := client.LLen(key).Result()
		if err != nil {
			log.Println(err)
		}
		for index := 0; index < int(_len); index++ {
			_ret, err := client.LIndex(key, int64(index)).Result()
			if err != nil {
				log.Println(err)
			}
			ret_tmp = append(ret_tmp, _ret)
		}
		ret, _ = json.Marshal(ret_tmp)

	}

	c.Ctx.WriteString(string(ret))
}
