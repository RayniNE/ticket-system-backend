package server

import (
	"database/sql"
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/raynine/ticket-system-backend/models"
	"github.com/raynine/ticket-system-backend/server/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func migrateSchemas(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil

}

func createDB() (*gorm.DB, error) {

	connectionString := os.Getenv("DB_CONNECTION")

	log.Println("-- Opening Database Connection -- ")
	postgresDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	log.Println("-- Pinging Database -- ")
	err = postgresDB.Ping()
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgresDB,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	log.Println("-- Initializing Schema Migration -- ")
	err = migrateSchemas(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateServer() *mux.Router {

	db, err := createDB()
	if err != nil {
		log.Panic(err.Error())
	}

	options := models.Config{
		DB: db,
	}

	router := mux.NewRouter()
	userHandler := handlers.NewUserHandler(options)
	InitializeUserHandler(router, userHandler)

	return router
}

func InitializeUserHandler(route *mux.Router, userHandler *handlers.UserHandler) {
	userRouter := route.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/", userHandler.CreateUser).Methods("POST")
}
