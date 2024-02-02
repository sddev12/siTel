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

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
Handler for the POST /register route
*/
func RegisterHandler(app *app.App, c echo.Context) error {

	// Alias logger
	l := app.Echo.Logger

	l.Info("POST /register called")

	// Get username from request bosy
	reqBody := new(RegisterRequest)
	if err := c.Bind(reqBody); err != nil {
		l.Errorf("failed to parse request body on register request: %v", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "failed to parse request body on register request"})
	}

	username := reqBody.Username
	l.Infof("Received username: %s", username)

	if username == "" {
		l.Error("No username supplied")
		return c.JSON(http.StatusBadRequest, ErrorResponse{Error: "No username supplied"})
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
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "An error occurred when creating user account"})
		}

		jsonUser, err := json.Marshal(insertResult)
		if err != nil {
			l.Errorf("ERROR: Could not marshall insertResult to JSON: %v", err)
			return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to serialise db RegisterResponse"})
		}
		l.Infof("User created successfully: %s", string(jsonUser))

		// Create session ID and set sesion active with POST request to session service
		sessionId := uuid.New().String()
		sessionServiceUrl := fmt.Sprintf("http://%s:3003/set-session", os.Getenv("SITEL_SESSION_HOST"))
		err = utils.SetSession(sessionServiceUrl, sessionId)
		if err != nil {
			l.Errorf("set session call failed: %v", err)
		}

		return c.JSON(http.StatusOK, SuccessResponse{SessionId: sessionId, Message: "Successfully created account"})

	}

	// If some other error
	if err != nil && err != mongo.ErrNoDocuments {
		l.Errorf("Unable to retrieve user record from db: %s", err)
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "an error occurred during user validation"})
	}

	// If user record is found in MongoDB, user already exists error
	l.Infof("User already exists for: Submitted Username: %s - Found Username: %s", username, userResult.Username)
	return c.JSON(http.StatusConflict, ErrorResponse{Error: "user already exists"})
}
