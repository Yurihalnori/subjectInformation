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
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func InitModel() {
	if !DB.Migrator().HasTable(&Project{}) {
		_ = DB.Migrator().CreateTable(&Project{})
		DB.Exec("ALTER TABLE projects ADD FULLTEXT (title) WITH PARSER ngram")
	}
	if !DB.Migrator().HasTable(&Institute{}) {
		_ = DB.Migrator().CreateTable(&Institute{})
	}
	if !DB.Migrator().HasTable(&Book{}) {
		_ = DB.Migrator().CreateTable(&Book{})
		DB.Exec("ALTER TABLE books ADD FULLTEXT (title) WITH PARSER ngram")
	}
	if !DB.Migrator().HasTable(&Dissertation{}) {
		_ = DB.Migrator().CreateTable(&Dissertation{})
		DB.Exec("ALTER TABLE dissertations ADD FULLTEXT (title) WITH PARSER ngram")
	}
	if !DB.Migrator().HasTable(&Article{}) {
		_ = DB.Migrator().CreateTable(&Article{})
		DB.Exec("ALTER TABLE articles ADD FULLTEXT (title) WITH PARSER ngram")
	}
	if !DB.Migrator().HasTable(&Tutor{}) {
		_ = DB.Migrator().CreateTable(&Tutor{})
	}
	if !DB.Migrator().HasTable(&Category{}) {
		_ = DB.Migrator().CreateTable(&Category{})
	}
	if !DB.Migrator().HasTable(&News{}) {
		_ = DB.Migrator().CreateTable(&News{})
		DB.Exec("ALTER TABLE news ADD FULLTEXT (text) WITH PARSER ngram")
		DB.Exec("ALTER TABLE news ADD FULLTEXT (title) WITH PARSER ngram")
		DB.Exec("ALTER TABLE news ADD FULLTEXT (title,text) WITH PARSER ngram")
	}
	if !DB.Migrator().HasTable(&User{}) {
		_ = DB.Migrator().CreateTable(&User{})
	}
	if !DB.Migrator().HasTable(&UniqueDatabase{}) {
		_ = DB.Migrator().CreateTable(&UniqueDatabase{})
		DB.Exec("ALTER TABLE unique_databases ADD FULLTEXT (name) WITH PARSER ngram")
	}
	if !DB.Migrator().HasTable(&Teamwork{}) {
		_ = DB.Migrator().CreateTable(&Teamwork{})
		DB.Exec("ALTER TABLE teamworks ADD FULLTEXT (name) WITH PARSER ngram")
	}
	_ = DB.AutoMigrate(&News{}, &User{}, &UniqueDatabase{}, &Teamwork{},
		&Project{}, &Book{}, &Dissertation{}, &Article{}, &Category{})
	_ = DB.AutoMigrate(&Institute{}, &Tutor{})
	// if err != nil {
	// 	panic(err)
	// }
}
