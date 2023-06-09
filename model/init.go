package model

import (
	"fmt"
	"log"
	"time"

	"subjectInformation/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
		config.Config.MysqlUser,
		config.Config.MysqlPass,
		config.Config.MysqlHost,
		config.Config.MysqlPort,
		config.Config.MysqlName)
	var dbLogger logger.Interface
	if config.DatabaseLogger == nil {
		dbLogger = logger.Default.LogMode(logger.Info)
	} else {
		dbLogger = logger.New(
			log.New(config.DatabaseLogger, "\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  false,
			},
		)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: dbLogger})
	if err != nil {
		panic(err)
	}
	DB = db
}

func InitModel() {
	DB.AutoMigrate(&Article{}, &Book{}, &Category{}, &Dissertation{}, &Institute{}, &Project{})
	if !DB.Migrator().HasTable(&Article{}) {
		DB.Migrator().CreateTable(&Article{})
	}
	if !DB.Migrator().HasTable(&Book{}) {
		DB.Migrator().CreateTable(&Book{})
	}
	if !DB.Migrator().HasTable(&Category{}) {
		DB.Migrator().CreateTable(&Category{})
	}
	if !DB.Migrator().HasTable(&Dissertation{}) {
		DB.Migrator().CreateTable(&Dissertation{})
	}
	if !DB.Migrator().HasTable(&Institute{}) {
		DB.Migrator().CreateTable(&Institute{})
	}
	if !DB.Migrator().HasTable(&Project{}) {
		DB.Migrator().CreateTable(&Project{})
	}
}
