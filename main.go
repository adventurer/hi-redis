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
	beego.SetStaticPath("/", "dist")

	// go func() {
	// 	var handel *exec.Cmd
	// 	cmd := "serve -s dist"
	// 	handel = exec.Command("/bin/sh", "-c", cmd)
	// 	output, err := handel.Output()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println(string(output))
	// }()
	beego.SetLevel(beego.LevelDebug)
	// beego.BeeLogger.DelLogger("console")

	beego.Run()
}

func websockTask() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.3.208:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	tk1 := toolbox.NewTask("tk1", "0/3 * * * * *", func() error {
		info, err := client.Info().Result()
		if err != nil {
			log.Fatal(err)
		}
		info_result := models.UnixConfigToMap(info)
		Json, _ := json.Marshal(info_result)
		controllers.Broadcast <- string(Json)
		return nil
	})
	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()
	// defer toolbox.StopTask()
}
