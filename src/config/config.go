package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionDbString = ""
	ApiPort            = 0
)

//initialize enviroment vars
func Load() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	ApiPort, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		ApiPort = 9000
	}
	// "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName";
	ConnectionDbString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
