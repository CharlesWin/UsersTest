package main

import (
	"UsersTest/internal/config"
	"UsersTest/internal/handler"
)

func main() {
	config.GetInstance()
	handler.Start()
}
