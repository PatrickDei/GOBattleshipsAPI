package app

import (
	"fmt"
	"github.com/PatrickDei/log-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"go-battleships/constants"
	"go-battleships/domain"
	"go-battleships/handlers"
	"go-battleships/service"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	dbClient := getDbClient()
	ph := handlers.PlayerHandler{
		Service: service.NewPlayerService(domain.NewPlayerRepository(dbClient)),
	}

	router.HandleFunc("/player", ph.CreatePlayer).Methods(http.MethodPost)

	address := os.Getenv(constants.ServerAddressEnv)
	port := os.Getenv(constants.ServerPortEnv)
	logger.Info("API listening on: " + address + ":" + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func sanityCheck() {
	if os.Getenv(constants.ServerAddressEnv) == "" {
		log.Fatal("Environment variable \"" + constants.ServerAddressEnv + "\" not defined...")
	}
	if os.Getenv(constants.ServerPortEnv) == "" {
		log.Fatal("Environment variable \"" + constants.ServerPortEnv + "\" not defined...")
	}
	if os.Getenv(constants.DbUserEnv) == "" {
		log.Fatal("Environment variable \"" + constants.DbUserEnv + "\" not defined...")
	}
	if os.Getenv(constants.DbPasswordEnv) == "" {
		log.Fatal("Environment variable \"" + constants.DbPasswordEnv + "\" not defined...")
	}
	if os.Getenv(constants.DbAddressEnv) == "" {
		log.Fatal("Environment variable \"" + constants.DbAddressEnv + "\" not defined...")
	}
	if os.Getenv(constants.DbPortEnv) == "" {
		log.Fatal("Environment variable \"" + constants.DbPortEnv + "\" not defined...")
	}
	if os.Getenv(constants.DbNameEnv) == "" {
		log.Fatal("Environment variable \"" + constants.DbNameEnv + "\" not defined...")
	}
}

func getDbClient() *sqlx.DB {
	user := os.Getenv(constants.DbUserEnv)
	password := os.Getenv(constants.DbPasswordEnv)
	address := os.Getenv(constants.DbAddressEnv)
	port := os.Getenv(constants.DbPortEnv)
	dbName := os.Getenv(constants.DbNameEnv)

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, address, port, dbName)

	db, err := sqlx.Open("mysql", datasource)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}