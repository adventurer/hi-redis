package controllers

import (
	"log"
)

func (c *MainController) Welcome() {

	c.Layout = "layout.tpl"
	c.TplName = "welcome.html"
}

func (c *MainController) Echo() {
	log.Println("run echo")
	ws, err := Upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	Clients[ws] = true
	c.StopRun()
}
