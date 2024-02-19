package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {
	// TODO:инициализировать объект конфига
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: инициализировать логгер

	// TODO: инициализировать приложение

	// TODO: запустить gRPC-сервер приложения

}
