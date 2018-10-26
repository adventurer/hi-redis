package controllers

import (
	"encoding/json"
	"hi-redis/models"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

type MainController struct {
	beego.Controller
}

func init() {
	go func() {
		tk1 := toolbox.NewTask("tk1", "0/1 * * * * *", func() error {
			info, err := models.Client.Info().Result()
			if err != nil {
				log.Fatal(err)
			}
			info_result := models.UnixConfigToMap(info)
			Json, _ := json.Marshal(info_result)
			Broadcast <- string(Json)
			return nil
		})
		toolbox.AddTask("tk1", tk1)
		toolbox.StartTask()
	}()

}
