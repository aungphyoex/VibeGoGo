package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "log"
    "os"
    "time"
)

func NewDBConnection() (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    // Custom logger configuration
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags),
        logger.Config{
            SlowThreshold:             time.Second,
            LogLevel:                  logger.Info,
            IgnoreRecordNotFoundError: true,
            Colorful:                  true,
        },
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: newLogger,
    })
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    // Get generic database object sql.DB to use its functions
    sqlDB, err := db.DB()
    if err != nil {
        return nil, fmt.Errorf("failed to get database instance: %v", err)
    }

    // SetMaxIdleConns sets the maximum number of connections in the idle connection pool
    sqlDB.SetMaxIdleConns(10)
    
    // SetMaxOpenConns sets the maximum number of open connections to the database
    sqlDB.SetMaxOpenConns(100)
    
    // SetConnMaxLifetime sets the maximum amount of time a connection may be reused
    sqlDB.SetConnMaxLifetime(time.Hour)

    return db, nil
}