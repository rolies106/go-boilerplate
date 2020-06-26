package models

import (
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Used by gorm
	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v9"

	u "mortred/utils"

	"fmt"
	"os"
	"reflect"
	"strings"
)

// Global instance
var DB *gorm.DB
var validate *validator.Validate
var ES *elasticsearch.Client

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	// Initiate validation
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

// GetDB get current DB connection
func GetDB() *gorm.DB {

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dbURI := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		u.Log("error", err)
	}

	DB = conn

	// Enable debug mod
	if os.Getenv("DB_DEBUG") == "true" {
		conn.LogMode(true)
	}

	return DB
}

// GetES get current ElasticSearch connection
func GetES() *elasticsearch.Client {

	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		u.Log("error", err)
	}

	ES = es

	return ES
}

// GetAdminDB Get connection into Admin DB
func GetAdminDB() *gorm.DB {

	username := os.Getenv("DB_ADMIN_USER")
	password := os.Getenv("DB_ADMIN_PASS")
	dbName := os.Getenv("DB_ADMIN_NAME")
	dbHost := os.Getenv("DB_ADMIN_HOST")
	dbPort := os.Getenv("DB_ADMIN_PORT")

	dbURI := username + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	return conn
}
