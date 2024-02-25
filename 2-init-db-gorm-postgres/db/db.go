package db

import (
	"fmt"
	"github.com/google/uuid"
	"go-samples/2-init-db-gorm-postgres/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"strings"
	"time"
)

const defaultDbLogLevel = gorm_logger.Warn

type User struct {
	Id       uuid.UUID `gorm:"primarykey"`
	FullName string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func InitDB(cfg *config.Config) (*gorm.DB, error) {
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

func dbSchema(cfg *config.Config) string {
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
