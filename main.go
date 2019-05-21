package main

import (
	"fmt"
	"os"

	gocb "gopkg.in/couchbase/gocb.v1"
)

type BucketInterface interface {
	Get(key string, value interface{}) (gocb.Cas, error)
	Insert(key string, value interface{}, expiry uint32) (gocb.Cas, error)
}

type Database struct {
	bucket BucketInterface
}

type Person struct {
	Type      string `json:"type"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (d Database) GetPersonDocument(key string) (interface{}, error) {
	var data interface{}
	_, err := d.bucket.Get(key, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d Database) CreatePersonDocument(key string, data interface{}) (interface{}, error) {
	_, err := d.bucket.Insert(key, data, 0)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	fmt.Println("Starting the application...")
	var database Database
	cluster, _ := gocb.Connect("couchbase://" + os.Getenv("DB_HOST"))
	cluster.Authenticate(gocb.PasswordAuthenticator{Username: os.Getenv("DB_USER"), Password: os.Getenv("DB_PASS")})
	database.bucket, _ = cluster.OpenBucket(os.Getenv("DB_BUCKET"), "")
	fmt.Println(database.GetPersonDocument("8eaf1065-5bc7-49b5-8f04-c6a33472d9d5"))
	database.CreatePersonDocument("blawson", Person{Type: "person", Firstname: "Brett", Lastname: "Lawson"})
}
