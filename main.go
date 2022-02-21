package main

import (
	"LinkShortenerApi/pkg/apiserver"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"net/http"
)

type Config struct {
	Addr string
}

func main() {
	var conf Config
	if _, err := toml.DecodeFile("configs/apiserverConfig.toml", &conf); err != nil {
		fmt.Println("Ошибка при считывании config файла.")
	}
	server := apiserver.NewServer(conf.Addr)
	log.Fatal(http.ListenAndServe(server.Addr, server.Handler))
}
