package main

import (
	"log"
	"popular/app/controllers"
)

func main() {
	CmdFlag()
	log.Println(controllers.StartWebServer())
}
