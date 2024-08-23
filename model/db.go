package model

import (
	"fmt"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// Conn 所有的数据库操作放在这里
var Conn *gorm.DB
var Rdb *redis.Client
var MongoDB *mongo.Client

// Database 定义 MongoDB 连接的结构体
type Database struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func NewMysql() {
	my := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "book", "6zJYcP75TCMDCBrp", "192.168.30.38", "book")
	conn, err := gorm.Open(mysql.Open(my), &gorm.Config{})
	if err != nil {
		fmt.Printf("err:%s\n", err)
		panic(err)
	}
	Conn = conn
}

func NewRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.30.38:6379",
		Password: "",
		DB:       0,
	})
	Rdb = rdb

	//初始化session
	Store, _ = redisstore.NewRedisStore(context.TODO(), Rdb)
	return

}

// ConnectToDB 连接到 MongoDB
func ConnectToDB() (*Database, error) {
	// 设置 MongoDB 连接
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://192.168.30.38:27017"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db := client.Database("library")
	collection := db.Collection("avatars")

	return &Database{
		Client:     client,
		Database:   db,
		Collection: collection,
	}, nil
}

func Close() {
	db, _ := Conn.DB()
	_ = db.Close()
	_ = Rdb.Close()
	_ = MongoDB.Disconnect(context.TODO())
}
