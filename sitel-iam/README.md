# todo-iam
IAM service for the todo example application

Handles registration and login

## Config
Uses the godotenv package so environment variables can be set in a `.env` file for development.

### Required Env Vars
| Environment Variable | Description                           | Accepted Values               | 
|----------------------|---------------------------------------|-------------------------------|
| MONGO_DB_URI         | MongoDB connection string             | Any MongoDB connection string |
| LOGGING_LEVEL        | Logging level for the application     | DEBUG, INFO, WARN, ERROR, OFF |
