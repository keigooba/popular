package main

import (
	"log"
	"タイトル/app/controllers"
)

func main() {
	CmdFlag()
	log.Println(controllers.StartWebServer())
}
