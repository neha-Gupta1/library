package utils

import (
	"fmt"
	"os"
)

// Database constants
var (
	Library = "library"
	Books   = "books"
)

func GetDBURL() string {
	mongoHost := os.Getenv("MONGOHOST")
	return fmt.Sprintf("mongodb://" + mongoHost + ":27017")

}
