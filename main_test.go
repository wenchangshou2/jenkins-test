package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/mitchellh/mapstructure"
	gocb "gopkg.in/couchbase/gocb.v1"
)

type MockBucket struct{}

var testdatabase Database

func convert(start interface{}, end interface{}) error {
	bytes, err := json.Marshal(start)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, end)
	if err != nil {
		return err
	}
	return nil
}

func (b MockBucket) Get(key string, value interface{}) (gocb.Cas, error) {
	switch key {
	case "nraboy":
		err := convert(Person{Type: "person", Firstname: "Nic", Lastname: "Raboy"}, value)
		if err != nil {
			return 0, err
		}
	default:
		return 0, gocb.ErrKeyNotFound
	}
	return 1, nil
}

func (b MockBucket) Insert(key string, value interface{}, expiry uint32) (gocb.Cas, error) {
	switch key {
	case "nraboy":
		return 0, gocb.ErrKeyExists
	}
	return 1, nil
}

func TestMain(m *testing.M) {
	testdatabase.bucket = &MockBucket{}
	os.Exit(m.Run())
}

func TestGetPersonDocument(t *testing.T) {
	data, err := testdatabase.GetPersonDocument("nraboy")
	if err != nil {
		t.Fatalf("Expected `err` to be `%s`, but got `%s`", "nil", err)
	}
	var person Person
	mapstructure.Decode(data, &person)
	if person.Type != "person" {
		t.Fatalf("Expected `type` to be %s, but got %s", "person", person.Type)
	}
}

func TestCreatePersonDocument(t *testing.T) {
	_, err := testdatabase.CreatePersonDocument("blawson", Person{Type: "person", Firstname: "Brett", Lastname: "Lawson"})
	if err != nil {
		t.Fatalf("Expected `err` to be `%s`, but got `%s`", "nil", err)
	}
}
