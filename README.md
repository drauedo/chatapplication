# GOLANG chat application

Chat application

#USAGE

docker build -t chat
docker run -dp 8080:8080 chat

Once the  server is deployed, you can access it using a telnet:

telnet localhost 8888

The available commands are:

/nick : Define the user name.
/join : Join a room. If it doesn't exist,it creates the room.
/msg  : Send a message. You need to be in the room to send a message.
/quit : Close the connection.
/help : List the available commands.