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

// Handler for the POST /login route
func LoginHandler(app *app.App, c echo.Context) error {

	l := app.Echo.Logger

	l.Info("GET /login called")
	// Get username from request body
	reqBody := new(RegisterRequest)
	if err := c.Bind(reqBody); err != nil {
		l.Errorf("failed to parse request body on register request: %v", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to parse request body on register request"})
	}

	username := reqBody.Username
	l.Infof("Received username: %s", username)

	if username == "" {
		l.Error("No username supplied")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "no username supplied"})
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
		return c.JSON(http.StatusNotFound, ErrorResponse{Error: "user account not found. the user doesn't exist"})
	}

	// Error communicating with the DB
	if err != nil && err != mongo.ErrNoDocuments {
		l.Errorf("Unable to retrieve user record from db: %s", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "an error occurred during user validation"})
	}

	// User found
	l.Infof("retreived user from DB - Login successful - User: %s", userResult.Username)
	l.Info("Generating session token")
	sessionId := uuid.New().String()
	l.Infof("Generated session token: %s", sessionId)
	return c.JSON(http.StatusOK, SuccessResponse{SessionId: sessionId, Message: "User found, successfully logged in"})
}
