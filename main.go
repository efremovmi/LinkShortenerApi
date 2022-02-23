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
	host, user, password, dbname, addr, tableName string
	port                                          int
	isStoreWithDB                                 bool
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
	addr = os.Getenv("ADDR")
	host = os.Getenv("POSTGRES_HOST")
	port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_BD_NAME")
	tableName = os.Getenv("POSTGRES_TABLE_NAME")
	var err error
	//isStoreWithDB, err = strconv.ParseBool(os.Getenv("FLAG_STORE_WITH_DATABASE"))
	isStoreWithDBStr, isSet := os.LookupEnv("FLAG_STORE_WITH_DATABASE")
	if !isSet {
		log.Fatal("FLAG_STORE_WITH_DATABASE not set")
	}

	isStoreWithDB, err = strconv.ParseBool(isStoreWithDBStr)
	if err != nil {
		log.Fatal("FLAG_STORE_WITH_DATABASE is not a boolean variable")
	}

}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	server := apiserver.NewServer(addr, psqlconn, tableName, isStoreWithDB)
	log.Fatal(http.ListenAndServe(server.Addr, server.Handler))
}
