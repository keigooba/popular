package main

import (
	"fashion/app/controllers"
	"log"
)

func main() {
	CmdFlag()
	log.Println(controllers.StartWebServer())
}
