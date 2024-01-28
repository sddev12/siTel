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
    console.log("Redis client error", err);
  });

redisClient.connect();

app.use(express.json());

app.post("/set-session", (req, res) => {
  var sessionId = req.body.sessionId;
  redisClient.set(sessionId, "authorised", {
    EX: 60,
  });

  res.status(200).json({ message: "Session stored in Redis" });
});

app.listen(port, () => {
  console.log(`Server listening on port: ${port}`);
});
