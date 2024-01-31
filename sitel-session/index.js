import express from "express";
import redis from "redis";
import dotenv from "dotenv";

dotenv.config();

const app = express();
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

app.post("/set-session", (req, res) => {
  var sessionId = req.body.sessionId;
  (async (key, res) => {
    const setResponse = await redisClient.set(key, "authorised", {
      EX: 120,
    });
    if (setResponse != "OK") {
      return res.status(500).json({ message: "Failed to create session" });
    }

    return res.status(200).json({ message: "Session stored in Redis" });
  })(sessionId, res);
});

app.post("/get-session", (req, res) => {
  var sessionId = req.body.sessionId;
  (async (key, res) => {
    const authorised = await redisClient.get(key);
    if (authorised == null) {
      return res.status(401).json({ message: "unauthorsied" });
    }

    return res.status(200).json({ message: "authorised" });
  })(sessionId, res);
});

app.listen(port, () => {
  console.log(`Server listening on port: ${port}`);
});
