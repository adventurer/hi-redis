package controllers

import (
	"encoding/json"
	"hi-redis/models"
	"log"
	"strings"

	"github.com/go-redis/redis"
)

type Data struct {
	Type string
	Val  string
	Ttl  float64
}

func (c *MainController) Data() {

	key := c.GetString("key")

	var data = make([]Data, 0)
	keys, _, err := models.Client.Scan(0, key+"*", 3000).Result()
	if err != nil {
		panic(err)
	}

	for _, v := range keys {
		_type, _ := models.Client.Type(v).Result()
		ttl, _ := models.Client.TTL(v).Result()
		d := Data{Type: _type, Val: v, Ttl: ttl.Seconds()}
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

	var ret []byte
	t, _ := models.Client.Type(key).Result()
	switch t {
	case "string":
		result, err := models.Client.Get(key).Result()
		if err != nil {
			log.Println(err)
		}
		ret, _ = json.Marshal(result)

	case "hash":
		result, err := models.Client.HGetAll(key).Result()
		if err != nil {
			log.Println(err)
		}
		ret, _ = json.Marshal(result)
	case "set":
		result, err := models.Client.SMembers(key).Result()
		if err != nil {
			log.Println(err)
		}
		ret, _ = json.Marshal(result)
	case "list":
		var ret_tmp []string
		_len, err := models.Client.LLen(key).Result()
		if err != nil {
			log.Println(err)
		}
		for index := 0; index < int(_len); index++ {
			_ret, err := models.Client.LIndex(key, int64(index)).Result()
			if err != nil {
				log.Println(err)
			}
			ret_tmp = append(ret_tmp, _ret)
		}
		ret, _ = json.Marshal(ret_tmp)

	}

	c.Ctx.WriteString(string(ret))
}

func (c *MainController) Command() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	input := c.GetString("cmd")
	if input == "" {
		c.Ctx.WriteString("input is empty")
		return
	}

	args := []interface{}{}
	inputArr := strings.Split(input, " ")
	for _, v := range inputArr {
		args = append(args, v)
	}
	log.Println(args...)
	cmd := redis.NewStringCmd(args...)
	log.Println(cmd)
	err := models.Client.Process(cmd)
	if err != nil {
		c.Ctx.WriteString("err on process:" + err.Error())
		return
	}
	t, err := cmd.Result()
	if err != nil {
		c.Ctx.WriteString("err on result:" + err.Error())
		return
	}

	ret, _ := json.Marshal(t)
	c.Ctx.WriteString(string(ret))

}
