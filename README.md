# Simple Mail Service Application

This application implements http protocol to send/receive messages for its users

## Usage:
Docker-compose corrently unavailable, please run 2 server files which are - "mainserver/server.go and and grpcserver/grpcserver.go" (In order to use app you need to run it with docker-compose command "docker-compose up", but previously please PULL USER-SERVICE project and insert its path to "docker-compose.yml" file.)

After running project app you can use admin requests or user request:
* POST request to PATH:PORT_FOR_USER_SERVICE/signup will add user to USER-SERVICE and will give you unique JWT token. USER-SERVICE must be run on PORT_FOR_USER_SERVICE (add Name&Password in requests body) //user

* with POST request to PATH:PORT/messages you can write a message. Please fill the next fields in requests body ("name" - person you write to): 
```
{
    "content": "Hi Steve, how are you?",
    "firstname": "Jake",
    "lastname": "Gyllenhaal",
    "name": "Steve"
}
```
Request is for users use. Auth with JWT token is temporarily unavailable.

* GET request to PATH:PORT/inbox will show your received messages. (Provide your name in requests body) 
```
{
    "firstname": "Jake",
}
```
Request is for users use. Auth with JWT token is temporarily unavailable.

* with PUT request to PATH:PORT/messages/{id} you are able to Update message.
```
{
    "content": "Hi Steve, how are you and how is your mom?",
    "firstname": "Jake",
    "lastname": "Gyllenhaal",
    "name": "Steve"
}
```
Request is for admin use.

* with DELETE request to PATH:PORT/messages/{id} you are able to Delete the message. 

Provide id of the message.

Request is for admin use.

* with GET request to PATH:PORT/messages/{id} you are able to read message by its id.

Provide id of the message.

Request is for admin use.

* with DELETE request to PATH:PORT/delete you are able to delete your income messages.

Please write your name in "firstname" field and the name of person who wrote you in "name" field.

```
{
    "firstname": "Steve",
    "name": "Jake"
}
```

Request is for users use.