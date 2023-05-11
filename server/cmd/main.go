package main

import "MyLink_Server/server/internal/app"

func main() {
	serv := app.NewServer()
	serv.Init()

}
