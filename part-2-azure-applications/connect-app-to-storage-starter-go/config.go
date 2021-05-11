package main

import "net/url"

// Config holds all the configuration variables that are needed for this app.
// Structure is direct copy from original Udacity, Python implementation.
type Config struct {
	secretKey                    string
	SQLServer                    string
	SQLDatabase                  string
	SQLUsername                  string
	SQLPassword                  string
	SQLAlchemyDatabaseURI        url.URL
	SQLAlchemyTrackModifications bool
	blobAccount                  string
	blobStorageKey               string
	blobContainer                string
}
