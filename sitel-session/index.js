import express from "express";
import morgan from "morgan";
import redis from "redis";
import dotenv from "dotenv";

// Init dotenv
dotenv.config();

// init express
const app = express();

// use morgan middleware for request and response logging
app.use(morgan("dev"));

const port = process.env.PORT || 3003;
const { REDIS_HOST, REDIS_PORT } = process.env;

const redisClient = redis
  .createClient({
    socket: {
      port: REDIS_PORT,
      host: REDIS_HOST,
    },
  })
  .on("error", (err) => {
    console.log("Redis client error: ", err);
  });

redisClient.connect();

app.use(express.json());

/**
 * POST /set-session
 * Sets sessionId in redis
 * The sitel-iam service calls this route on successful register and login
 * to store the serverside session for the user
 **/
app.post("/set-session", (req, res) => {
  console.log("/set-session route called");
  // Get sessionID
  var sessionId = req.body.sessionId;

  // set sessionId in redis
  console.log("Setting sessionId in redis");
  (async (key, res) => {
    const setResponse = await redisClient.set(key, "authorised", {
      EX: 60 * 60 * 2,
    });
    if (setResponse != "OK") {
      console.log(`Error setting sesionId in redis : ${setResponse}`);
      return res.status(500).json({ message: "Failed to create session" });
    }

    console.log("Successfully set sessionId in redis, returning 200");
    return res.status(200).json({ message: "Session stored in Redis" });
  })(sessionId, res);
});

/**
 * POST /get-session
 * Attepts to retrieve the posted sesionId from redis
 * If sessionId exists, retruns OK
 * If sessionId doesn't exist, returns 401
 */
app.post("/get-session", (req, res) => {
  console.log("/get-sesion called");
  var sessionId = req.body.sessionId;

  console.log("Getting sessionId from redis");
  (async (key, res) => {
    const authorised = await redisClient.get(key);
    if (authorised == null) {
      console.log("sessionId not in redis, returning 401");
      return res.status(401).json({ message: "unauthorsied" });
    }

    console.log("sessionId found in redis, returning 200");
    return res.status(200).json({ message: "authorised" });
  })(sessionId, res);
});

app.listen(port, () => {
  console.log(`Server listening on port: ${port}`);
});
