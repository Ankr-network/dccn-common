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
	err         error
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
		instance = mongodbconnect()
	})
	return instance
}

func getMgoConfig() *cfg.MgoConfig {
	config := GetConfig()
	mc, err := cfg.GetMgoConfig(config.VaultAddr, config.VaultRole, config.DataPath)
	if err != nil {
		panic(err)
	}
	return mc
}

func mongodbconnect() *mgo.Database {
	session, err = CreateDBSession()
	if err != nil {
		panic(err)
	}
	return session.DB(getMgoConfig().DbName)
}

func GetDB(database string) *mgo.Database {
	if session == nil {
		session, err = CreateDBSession()
		if err != nil {
			panic(err)
		}
	}
	return session.DB(database)
}

func CreateDBSession() (*mgo.Session, error) {
	config := GetConfig()
	mc, err := cfg.GetMgoConfig(config.VaultAddr, config.VaultRole, config.DataPath)
	if err != nil {
		panic(err)
	}
	logStr := fmt.Sprintf("mongodb hostname : %s", mc.Address)
	WriteLog(logStr)
	session, err := mgo.DialWithTimeout(
		fmt.Sprintf("mongodb://%s:%s@%s", mc.UserName, mc.PassWd, mc.Address),
		15*time.Second)
	if err != nil {
		WriteLog("can not connect to database")
		return nil, err
	}
	//defer session.Close()
	session.SetPoolLimit(int(mc.PoolSize))
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
