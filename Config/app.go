package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// mysql://root:EEf334-BHbDBDCgCFh1d3adg6cb4b561@monorail.proxy.rlwy.net:16383/railway
	d, err := gorm.Open("mysql", "root:EEf334-BHbDBDCgCFh1d3adg6cb4b561@monorail.proxy.rlwy.net:16383/railway")
	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {

	return db

}
