import { Server } from "socket.io";
import express from "express";
import http from "http";
import cors from "cors";
import dotenv from "dotenv";

dotenv.config()

const app = express();
const server = http.createServer(app);
app.use(cors());

const io = new Server(server, {
  cors: {
    origin: process.env.FRONTEND_URL,
    methods: ["GET", "POST"],
  },
});

let playerCount = 0;
io.on("connection", (socket) => {
  if (playerCount <= 0) playerCount = 0;

  socket.on("client_connected", () => {
    playerCount++;
    socket.emit("player_number", playerCount);

    if (playerCount === 2) {
      io.emit("start_timer", 0);
    }
  });

  socket.on("disconnect", () => {
    console.log("Player disconnect...");
    playerCount--;
    if (playerCount <= 0) playerCount = 0;
  });

  socket.on("send_message", ({ name, object }) => {
    socket.broadcast.emit("receive_message", { name, object });
  });
});

server.listen(5174, () => {
  console.log("Server is running");
});
