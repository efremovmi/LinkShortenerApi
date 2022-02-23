package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func getPathToEnv(curPath string) string {
	count := 0
	for i := len(curPath) - 1; i >= 0; i-- {
		if curPath[i] == '/' {
			count += 1
			if count == 2 {
				return curPath[:i]
			}
		}
	}
	return ""
}

func GetSqlconnAndTableName() (string, string) {
	curDir, _ := os.Getwd()
	pathToEnv := getPathToEnv(curDir)
	if err := godotenv.Load(pathToEnv + "/.env"); err != nil {
		log.Fatal("No .env file found")
	}
	//addr := os.Getenv("ADDR")
	host := os.Getenv("POSTGRES_HOST_TEST")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT_TEST"))
	user := os.Getenv("POSTGRES_USER_TEST")
	password := os.Getenv("POSTGRES_PASSWORD_TEST")
	dbname := os.Getenv("POSTGRES_BD_NAME_TEST")
	tableName := os.Getenv("POSTGRES_TABLE_NAME_TEST")
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname), tableName
}

func GetAddrProtocolWithDomain() (string, string) {
	curDir, _ := os.Getwd()
	pathToEnv := getPathToEnv(curDir)
	if err := godotenv.Load(pathToEnv + "/.env"); err != nil {
		log.Fatal("No .env file found")
	}
	addr := os.Getenv("ADDR")
	protocolWithDomain := os.Getenv("PROTOCOL_With_DOMAIN")
	return addr, protocolWithDomain

}
