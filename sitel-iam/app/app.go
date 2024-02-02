package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	elog "github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Echo  *echo.Echo
	Mongo *mongo.Client
}

func stringToLoggingLevel(levelStr string) (elog.Lvl, error) {
	switch levelStr {
	case "":
		return elog.ERROR, nil
	case "DEBUG":
		return elog.DEBUG, nil
	case "INFO":
		return elog.INFO, nil
	case "WARN":
		return elog.WARN, nil
	case "ERROR":
		return elog.ERROR, nil
	case "OFF":
		return elog.OFF, nil
	default:
		return elog.OFF, fmt.Errorf("unsupported log level: %s", levelStr)
	}
}

func NewApp() (*App, error) {
	/*
		Initialise godotenv
	*/
	log.Println("Initialising godotenv")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	/*
		Load logging level
	*/
	log.Println("Loading logging level from env")
	levelStr := os.Getenv("LOGGING_LEVEL")
	if levelStr == "" {
		log.Fatal("Unable to get LOGGING_LEVEL from env")
	}
	loggingLevel, err := stringToLoggingLevel(levelStr)
	if err != nil {
		log.Fatal(err)
	}

	/*
		Load Mongodb connection string
	*/
	log.Println("Building mongodb connection string")
	mongoDbHost := os.Getenv("MONGO_DB_HOST")
	if mongoDbHost == "" {
		log.Fatal("unable to get mongodb host from env vars")
	}
	mongoDbPort := os.Getenv("MONGO_DB_PORT")
	if mongoDbPort == "" {
		log.Fatal("unable to get mongodb port from env vars")
	}

	mongoConnectionString := fmt.Sprintf("mongodb://%s:%s/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.1.1", mongoDbHost, mongoDbPort)

	/*
		Set up Mongodb client
	*/
	log.Println("Setting mongodb client options")
	mongoClientOptions := options.Client().ApplyURI(mongoConnectionString)

	/*
		Create Mongodb client and connect to Mongodb
	*/
	log.Println("Creating mongodb client and connecting to mongodb")
	client, err := mongo.Connect(context.TODO(), mongoClientOptions)
	if err != nil {
		return nil, err
	}

	/*
		Test Mongodb connection
	*/
	log.Println("Testing mongodb connection...")
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return nil, err
	}
	log.Println("Successfully connected to mongoDB")

	log.Println("App initialisation successful")

	/*
		Create Echo instance
	*/
	e := echo.New()

	/*
		Use Middleware
	*/
	e.Use(middleware.Logger())
	e.Logger.SetLevel(loggingLevel)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.POST},
	}))

	return &App{Echo: e, Mongo: client}, nil
}
