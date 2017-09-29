package mongo

import (
	"errors"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

var (
	mongodb *mgo.Database
)

type MongoConfig struct {
	Addresses []string
	UserName  string
	Password  string
	Database  string
}

type MongoComponent struct{}

func (MongoComponent) Init(ops ...interface{}) (err error) {
	if len(ops) == 0 {
		return errors.New("param error")
	}
	cfg := ops[0].(*MongoConfig)

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    cfg.Addresses,
		Username: cfg.UserName,
		Password: cfg.Password,
		Timeout:  2 * time.Second,
	})
	if err != nil {
		log.Fatal("New mongo error:", err)
		return
	}
	mongodb = session.DB(cfg.Database)
	return nil
}

func Get() *mgo.Database {
	return mongodb
}
