package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/user"
	"todo-iam/app"
	"todo-iam/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterHandler(app *app.App, c echo.Context) error {

	l := app.Echo.Logger

	l.Info("POST /register called")
	username := c.FormValue("username")
	if username == "" {
		l.Error("No username supplied")
		return c.JSON(http.StatusBadRequest, RegisterResponse{LogLevel: "WARN", Message: "No username supplied"})
	}

	l.Info("Checking if user exists")

	// Set Mongodb collection
	coll := app.Mongo.Database("todo").Collection("users")

	// Attempt to get user from DB
	var userResult user.User
	err := coll.FindOne(context.TODO(), bson.D{{Key: "username", Value: username}}).Decode(&userResult)

	// If user is not found - create user record in mongoDB (happy path)
	if err == mongo.ErrNoDocuments {
		insertResult, err := coll.InsertOne(context.TODO(), bson.D{{Key: "username", Value: username}})
		if err != nil {
			l.Errorf("failed to insert user record to mongodb: %v", err)
			return c.JSON(http.StatusInternalServerError, RegisterResponse{LogLevel: "ERROR", Message: "An error occurred when creating user account"})
		}

		jsonUser, err := json.Marshal(insertResult)
		if err != nil {
			l.Errorf("ERROR: Could not marshall insertResult to JSON: %v", err)
			return c.JSON(http.StatusInternalServerError, RegisterResponse{LogLevel: "ERROR", Message: "Failed to serialise db RegisterResponse"})
		}
		l.Infof("User created successfully: %s", string(jsonUser))

		// Create session ID and set sesion active with POST request to session service
		sessionId := utils.GenerateSessionId()
		err = utils.SetSession(fmt.Sprintf("http://%s:3003/set-session", os.Getenv("TODO_SESSION_HOST")), sessionId)
		if err != nil {
			l.Errorf("set session call failed: %v", err)
		}

		return c.JSON(http.StatusOK, RegisterResponse{LogLevel: "Info", Message: string(jsonUser)})

	}

	// If some other error
	if err != nil && err != mongo.ErrNoDocuments {
		l.Errorf("Unable to retrieve user record from db: %s", err)
		return c.JSON(http.StatusInternalServerError, RegisterResponse{LogLevel: "ERROR", Message: "An error occurred during user validation"})
	}

	// If user record is found in MongoDB, user already exists error
	l.Infof("User already exists for: Submitted Username: %s - Found Username: %s", username, userResult.Username)
	return c.JSON(http.StatusConflict, RegisterResponse{LogLevel: "ERROR", Message: "User already exists"})
}
