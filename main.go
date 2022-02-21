package main

import (
	"fmt"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/apiserver"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	host, user, password, dbname, addr string
	port                               int
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return
	}
	addr = os.Getenv("ADDR")
	host = os.Getenv("POSTGRES_HOST")
	port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_BD_NAME")

}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	server := apiserver.NewServer(addr, psqlconn)
	log.Fatal(http.ListenAndServe(server.Addr, server.Handler))

}
