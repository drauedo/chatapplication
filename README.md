# GOLANG chat application #

Chat application

## USAGE ##

Build image:

```docker build -t chat```

Deploying image by default:

```docker run -dp 8080:8080 chat```

Deploying image specifying port:

```docker run -dp 9090:9090 -e PORT=":9090" chat```

Once the  server is deployed, you can access it using a telnet:

```telnet localhost [PORT]```

The available commands are:

- /nick : Define the user name.

- /rooms: List all the available rooms

- /join : Join a room. If it doesn't exist,it creates the room.

- /msg  : Send a message. You need to be in the room to send a message.

- /quit : Close the connection.

- /help : List the available commands.
