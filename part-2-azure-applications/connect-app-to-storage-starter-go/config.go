package main

import (
	"log"
	"net/url"
	"os"
)

// Config holds all the configuration variables that are needed for this app.
// Structure is direct copy from original Udacity, Python implementation.
type Config struct {
	secretKey                    string
	SQLServer                    string
	SQLDatabase                  string
	SQLUsername                  string
	SQLPassword                  string
	SQLAlchemyDatabaseURI        *url.URL
	SQLAlchemyTrackModifications bool
	blobAccount                  string
	blobStorageKey               string
	blobContainer                string
	port						 int
}

func NewConfig() Config {
	config := Config{
		secretKey:                    os.Getenv("SECRET_KEY"),
		SQLDatabase:                  os.Getenv("SQL_DATABASE"),
		SQLPassword:                  os.Getenv("SQL_PASSWORD"),
		SQLUsername:                  os.Getenv("SQL_SERVER"),
		SQLServer:                    os.Getenv("SQL_USER_NAME"),
		SQLAlchemyTrackModifications: false,
		blobAccount:                  os.Getenv("BLOB_ACCOUNT"),
		blobContainer:                os.Getenv("BLOB_CONTAINER"),
		blobStorageKey:               os.Getenv("BLOB_STORAGE_KEY"),
		port:                         5555,
	}

	var err error
	config.SQLAlchemyDatabaseURI, err = url.Parse("mssql+pyodbc://" +
		config.SQLUsername + "@" + config.SQLServer + ":" + config.SQLPassword 	+ "@" +
		config.SQLServer + ":1433/" + config.SQLDatabase + "?driver=ODBC+Driver+17+for+SQL+Server")

	if err != nil {
		log.Fatalf("Invalid SQLAlchemyDatabaseURI | %v", err)
	}

	return config
}
