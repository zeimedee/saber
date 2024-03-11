# PROJECT SABER
 this is an event based system using kafka as our message broker

 this project depends on docker to run kafka

 to get started first install docker 

 start kafka by running in the terminal
 ```
 docker- compose up -d
 ```
this downloads the kafka image and runs the container in the background

start the REST server using the makefile
```
make start
```

 A Post request is made to to send a value

```
http://localhost:4000/saber/send
```
request body
```
{
    "value":4
}
```

connect to a websocket to read the response

```
localhost:4000/saber/ws
```

