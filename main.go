package main

import (
	"encoding/json"
	"hi-redis/controllers"
	"hi-redis/models"
	_ "hi-redis/routers"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/go-redis/redis"
)

func main() {
	go websockTask()
	beego.SetStaticPath("/assets", "static")
	beego.Run()
}

func websockTask() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.208:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	log.Println("run task")
	// println("__________________")
	// log.Printf("%#v", info_result)
	tk1 := toolbox.NewTask("tk1", "0/3 * * * * *", func() error {
		info, err := client.Info().Result()
		if err != nil {
			log.Fatal(err)
		}
		info_result := models.UnixConfigToMap(info)
		Json, _ := json.Marshal(info_result)
		// msg := Message{Message: string(Json)}
		controllers.Broadcast <- string(Json)
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	// defer toolbox.StopTask()
}
