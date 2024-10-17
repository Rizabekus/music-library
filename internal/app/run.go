package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Rizabekus/music-library/internal/handlers"
	"github.com/Rizabekus/music-library/internal/services"
	"github.com/Rizabekus/music-library/internal/storage"
	"github.com/Rizabekus/music-library/pkg/loggers"
	"github.com/Rizabekus/music-library/pkg/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	loggers.InitLoggers()
	file, line, _ := utils.GetCallerInfo()
	loggers.InfoLog(file, line, "Loaded the configuration data from .env")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	loggers.InfoLog(file, line+7, "Successfully connected to database")

	m, err := migrate.New(
		"file://migrations",
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Force(1); err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			loggers.InfoLog(file, line+21, "No new migrations to apply.")
		} else {
			log.Fatal(err)
		}
	} else {
		version, _, err := m.Version()
		if err != nil {
			log.Fatal(err)
		}
		loggers.InfoLog(file, line, fmt.Sprintf("Successfully applied migrations. Current version: %d", version))
	}
	loggers.InfoLog(file, line+32, "Successfully applied migrations")

	storage := storage.StorageInstance(db)
	service := services.ServiceInstance(storage)
	handler := handlers.HandlersInstance(service)

	Routes(handler)

	defer db.Close()
}
