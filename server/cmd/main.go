package main

import (
	"github.com/MyLink_Server/MyLink_Server/internal/app"
)

func main() {
	serv := app.NewServer()
	serv.Init()

}
