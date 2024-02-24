package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

const defaultDbLogLevel = gorm_logger.Warn

type config struct {
	// DB config
	DbUserName             string `mapstructure:"db_username"`
	DbPassword             string `mapstructure:"db_password"`
	DbHost                 string `mapstructure:"db_host"`
	DbPort                 int    `mapstructure:"db_port"`
	DbName                 string `mapstructure:"db_name"`
	DbSchema               string `mapstructure:"db_schema"`
	DbPreferSimpleProtocol bool   `mapstructure:"db_prefer_simple_protocol"`
	DbPrepareStatement     bool   `mapstructure:"db_prepare_statement"`
	DbMigrate              bool   `mapstructure:"db_migrate"`
	DbLogLevel             string `mapstructure:"db_log_level"`
}

type User struct {
	Id       uuid.UUID `gorm:"primarykey"`
	FullName string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	// Init config
	cfg := &config{
		DbUserName:             "postgres",
		DbPassword:             "postgres",
		DbHost:                 "localhost",
		DbPort:                 5432,
		DbName:                 "go-samples",
		DbSchema:               "go-samples",
		DbPreferSimpleProtocol: false,
		DbPrepareStatement:     true,
		DbMigrate:              true,
		DbLogLevel:             "warning",
	}

	// Init database
	db, err := initDB(cfg)
	if err != nil {
		log.Fatalf("database init error: %s", err)
	}

	log.Printf("db: %v", db)
}

func initDB(cfg *config) (*gorm.DB, error) {
	dbLogLevel := dbLogLevelConvertor(cfg.DbLogLevel)

	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		cfg.DbHost,
		cfg.DbPort,
		cfg.DbUserName,
		cfg.DbName,
		cfg.DbPassword,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: cfg.DbPreferSimpleProtocol,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: dbSchema(cfg),
		},
		PrepareStmt: cfg.DbPrepareStatement,
		Logger:      gorm_logger.Default.LogMode(dbLogLevel),
	})
	if err != nil {
		return nil, err
	}

	if cfg.DbMigrate {
		err := migrateDb(db)
		if err != nil {
			log.Fatalf("database migration error: %v", err)
		}
	}

	return db, nil
}

func dbLogLevelConvertor(l string) gorm_logger.LogLevel {
	l = strings.ToLower(l)
	switch l {
	case "":
		return defaultDbLogLevel
	case "info":
		return gorm_logger.Info
	case "warning":
		return gorm_logger.Warn
	case "error":
		return gorm_logger.Error
	case "silent":
		return gorm_logger.Silent
	default:
		log.Fatal("incorrect db log level")
		return defaultDbLogLevel
	}
}

func dbSchema(cfg *config) string {
	var dbSchema string
	if len(cfg.DbSchema) > 0 {
		dbSchema = cfg.DbSchema + "."
	}
	return dbSchema
}

func migrateDb(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return fmt.Errorf("schema User migrate error: %s", err)
	}

	return nil
}
