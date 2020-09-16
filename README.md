# Simple Mail Service Application

This application implements http protocol to send/receive messages for its users

## Usage:
In order to use app you need to run it with docker-compose command "docker-compose up", but previously please PULL USER-SERVICE project and insert its path to "docker-compose.yml" file.

After running current project with USER-SERVICE app:
* POST request to PATH:PORT_FOR_USER_SERVICE/signup will add user to USER-SERVICE and will give you unique JWT token. USER-SERVICE must be run on PORT_FOR_USER_SERVICE (add Name&Password in requests body) //user

* with POST request to PATH:PORT/messages you can write a message. Please fill the next fields in requests body: {"id": "","in": "","content": "INPUT","author": {"idAU": "","firstname": "INPUT","lastname": "INPUT"},"name": "INPUT",},
where "Content" - your massage, "firstname" - your name, "name" - name of person which will receive the message. Please provide JWT token from previous request to the header with "Authorization" key of current request. //user

* GET request to PATH:PORT/messages gives all messages ever sent //admin

* GET request to PATH:PORT/message will show your received messages. (Provide your name in requests body) //user

* with PUT request to PATH:PORT/messages/{id} you are able to Update message. //admin

* with DELETE request to PATH:PORT/messages/{id} you are able to Delete the message. //admin

* with GET request to PATH:PORT/messages/{id} you are able to read message by its id. //admin

                                                                                                                                         