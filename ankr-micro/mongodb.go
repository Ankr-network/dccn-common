package ankrmicro

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

// MongoDBHost saves the endpoint of mongo db
var MongoDBHost string
var instance *mgo.Database
var once sync.Once
var session *mgo.Session

// GetCollection return the mongo db collection instance
func GetCollection(collection string) *mgo.Collection {
	db := GetDBInstance()
	c := db.C(collection)
	return c
}

// GetDBInstance return an instance of mongo db instance
func GetDBInstance() *mgo.Database {
	once.Do(func() {
		instance = mongodbconnect()
	})
	return instance
}

func mongodbconnect() *mgo.Database {
	config := GetConfig()
	logStr := fmt.Sprintf("mongodb hostname : %s", config.DatabaseHost)
	WriteLog(logStr)
	session, err := mgo.Dial(config.DatabaseHost)
	if err != nil {
		WriteLog("can not connect to database")
		return nil
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(config.DatabaseName)
	return db
}


func GetDB(database string)*mgo.Database{
	if session == nil {
		localSession, _ := CreateDBSession()
		session = localSession
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(database)
	return db
}


func CreateDBSession() (*mgo.Session,  error) {
	config := GetConfig()
	logStr := fmt.Sprintf("mongodb hostname : %s", config.DatabaseHost)
	WriteLog(logStr)
	session, err := mgo.Dial(config.DatabaseHost)
	if err != nil {
		WriteLog("can not connect to database")
		return session, err
	}
	return session, err
}
