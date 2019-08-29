package config

import (
	"fmt"
	"strconv"
)

type MgoConfig struct {
	UserName string
	PassWd   string
	Address  string
	DbName   string
	PoolSize int64
}

func GetMgoConfig(addr, saRole, keyPath string) (*MgoConfig, error) {
	whisper, err := getClient(addr)
	if err != nil {
		return nil, err
	}
	if secret, err := whisper.GetSecret(saRole, keyPath); err != nil {
		return nil, err
	} else {
		if secret.data != nil && secret.data.Data != nil {
			mc := &MgoConfig{}
			data := secret.data.Data
			if userName, ok := data["username"]; ok {
				mc.UserName, _ = userName.(string)
			}
			if passWd, ok := data["password"]; ok {
				mc.PassWd, _ = passWd.(string)
			}
			if addr, ok := data["addr"]; ok {
				mc.Address, _ = addr.(string)
			}
			if dbName, ok := data["name"]; ok {
				mc.DbName, _ = dbName.(string)
			}
			if size, ok := data["size"]; ok {
				ts, _ := size.(string)
				mc.PoolSize, err = strconv.ParseInt(ts, 10, 64)
				if err != nil {
					mc.PoolSize = 10
				}
			} else {
				mc.PoolSize = 10
			}
			return mc, nil
		} else {
			return nil, fmt.Errorf("config data null")
		}
	}
}
