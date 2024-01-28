package handlers

import (
	"context"
	"net/http"
	"os/user"
	"todo-iam/app"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginHandler(app *app.App, c echo.Context) error {

	l := app.Echo.Logger

	l.Info("GET /login called")
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

	// User doesnt exist in DB
	if err == mongo.ErrNoDocuments {
		l.Warn("user doest not exist in database")
		return c.JSON(http.StatusNotFound, RegisterResponse{LogLevel: "ERROR", Message: "User account not found. The user doesn't exist"})
	}

	// Error communicating with the DB
	if err != nil && err != mongo.ErrNoDocuments {
		l.Errorf("Unable to retrieve user record from db: %s", err)
		return c.JSON(http.StatusInternalServerError, RegisterResponse{LogLevel: "ERROR", Message: "An error occurred during user validation"})
	}

	// User found
	l.Infof("retreived user from DB - Login successful - User: %s", userResult.Username)
	l.Info("Generating session token")
	sessionToken := uuid.New()
	sessionTokenStr := sessionToken.String()
	l.Infof("Generated session token: %s", sessionTokenStr)
	return c.JSON(http.StatusOK, LoginResponse{Token: sessionTokenStr, Data: RegisterResponse{LogLevel: "Info", Message: "User found, successfully logged in"}})
}
