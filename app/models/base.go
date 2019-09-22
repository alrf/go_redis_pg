package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"fmt"
	"os"
	"strconv"
)

var db *gorm.DB
var rdb *redis.Client

func init() {

	dbUsername := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		dbPort = 5432
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		redisPort = 6379
	}

	redisUri := fmt.Sprintf("%s:%d", redisHost, redisPort)
	client := redis.NewClient(&redis.Options{
		Addr: redisUri,
		Password: "",
		DB: 0,  //use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		fmt.Println(redisUri)
	}
	rdb = client

	dbUri := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbName, dbPassword)
	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(dbUri)
		fmt.Println("DB error", err)
	}
	db = conn
	db.Debug().AutoMigrate(&Inventory{}).AddUniqueIndex("idx_dep_sect_eq", "department", "section", "equipment")

	fmt.Println("DB init was completed")
}

func GetDB() *gorm.DB {
	return db
}

func GetRDB() *redis.Client {
	return rdb
}
