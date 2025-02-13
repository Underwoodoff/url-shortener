package main

import (
	"fmt"
	config "url-shortener/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	//todo: init logger: slog

	//todo: init storage: sqlite

	//todo: init router: chi, chi render

	//todo: run server:

}
