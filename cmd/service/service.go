package main

import (
	"30.8.1/pkg/storage"
	"30.8.1/pkg/storage/postgres"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

var db storage.Interface

func main() {
	var err error
	constr, err := constr()
	if err != nil {
		log.Fatal(err)
	}
	db, err = postgres.New(constr)
	if err != nil {
		log.Fatal(err)
	}
	//db = memdb.DB{}
	tasks, err := db.Tasks(0, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tasks)
}

func constr() (string, error) {
	puser, exists := os.LookupEnv("POSTGRES_USER")
	if !exists {
		return "", errors.New("not exist postgres user")
	}
	pwd, exists := os.LookupEnv("POSTGRES_PASSWORD")
	if !exists {
		return "", errors.New("not exist postgres password")
	}
	pdb, exists := os.LookupEnv("POSTGRES_DB")
	if !exists {
		return "", errors.New("not exist postgres db")
	}
	port, exists := os.LookupEnv("POSTGRES_PORT")
	if !exists {
		return "", errors.New("not exist postgres port")
	}
	host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		return "", errors.New("not exist postgres host")
	}
	constr := "postgres://" + puser + ":" + pwd + "@" + host + ":" + port + "/" + pdb
	return constr, nil
}
