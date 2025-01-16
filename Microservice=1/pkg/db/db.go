package db

import (
	"database/sql"
	"fmt"
	"methodOne/pkg/config"
	"methodOne/pkg/model"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(config *config.DataBase) (*gorm.DB, error) {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPassword, config.DBPort)
	sqlDB, err := sql.Open("postgres", connectionString) 
	if err != nil {
		fmt.Println("Error connecting to Postgres:", err)
		return nil, err
	}
	defer sqlDB.Close()

	rows, err := sqlDB.Query("SELECT 1 FROM pg_database WHERE datname = $1", config.DBName)
	if err != nil {
		fmt.Println("Error querying for database existence:", err)
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		_, err = sqlDB.Exec("CREATE DATABASE " + config.DBName)
		if err != nil {
			fmt.Println("Error in creating database:", err)
			return nil, err
		}
		fmt.Println("Database created:", config.DBName)
	} else {
		fmt.Println("Database", config.DBName, "already exists")
	}

	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable", config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBPort)
	DB, dberr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

	if dberr != nil {
		fmt.Println("Error connecting to database with GORM:", dberr)
		return nil, dberr
	}

	if err := DB.AutoMigrate(&model.User{}); err != nil {
		fmt.Println("Error in migrating database:", err)
		return nil, err
	}

	return DB, nil
}
