package ankrmicro

import (
	"fmt"
	"sync"
	"time"

	cfg "github.com/Ankr-network/dccn-common/config"
	"gopkg.in/mgo.v2"
)

// MongoDBHost saves the endpoint of mongo db
var (
	MongoDBHost string
	instance    *mgo.Database
	once        sync.Once
	session     *mgo.Session
)

// GetCollection return the mongo db collection instance
func GetCollection(collection string) *mgo.Collection {
	db := GetDBInstance()
	c := db.C(collection)
	return c
}

// GetDBInstance return an instance of mongo db instance
func GetDBInstance() *mgo.Database {
	once.Do(func() {
		instance = mongodbConnect()
	})
	return instance
}

func mongodbConnect() *mgo.Database {
	session, err := CreateDBSession()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	config := GetConfig()
	if config.DbAuth {
		mc, err := cfg.GetMgoConfig(config.VaultAddr, config.VaultRole, config.DataPath)
		if err != nil {
			panic(err)
		}
		return session.DB(mc.DbName)
	}

	return session.DB(config.DatabaseName)
}

func GetDB(database string) *mgo.Database {
	if session == nil {
		localSession, _ := CreateDBSession()
		session = localSession
	}
	session.SetMode(mgo.Monotonic, true)
	db := session.DB(database)
	return db
}

func CreateDBSession() (*mgo.Session, error) {
	config := GetConfig()

	if config.DbAuth {
		// support auth
		mc, err := cfg.GetMgoConfig(config.VaultAddr, config.VaultRole, config.DataPath)
		if err != nil {
			panic("get secret from kms failed: " + err.Error())
		}
		logStr := fmt.Sprintf("mongodb hostname : %s", mc.Address)
		WriteLog(logStr)
		session, err := mgo.DialWithTimeout(
			fmt.Sprintf("mongodb://%s:%s@%s", mc.UserName, mc.PassWd, mc.Address),
			15*time.Second)
		if err != nil {
			panic(err)
		}
		if session != nil {
			session.SetPoolLimit(int(mc.PoolSize))
			session.SetMode(mgo.Monotonic, true)
		} else {
			panic("session can not be nil, please check")
		}
		return session, err
	}

	// compatible old way
	logStr := fmt.Sprintf("mongodb hostname : %s", config.DatabaseHost)
	WriteLog(logStr)
	session, err := mgo.Dial(config.DatabaseHost)
	if err != nil {
		panic(err)
	}
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
