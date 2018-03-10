package main

import (
	"Seeds/router"
	"Seeds/utils"
	"fmt"
)

func main() {
	c := utils.GetConfig()
	fmt.Println(c.Get("database.host"))
	e := router.Load()
	e.Run()
}
