# Golang REST API application with Redis and PostgreSQL, Jenkins pipeline for building it.

## Development

1. The Web Service reached under the path /humai/echoservice.

2. The Application has to realize the following logic:

	a. An HTTP request is sent against the endpoint /humai/echoservice.

	b. The request-payload contains a JSON object. The JSON object is extracted
and stored in a database (Redis as cache and PostgreSQL).

	c. A response is sent back with the original request-payload with in addition a
filed containing the server’s unix timestamp at the time of receiving the request.

## CI Pipeline

1. Jenkins pipeline executes the following steps:

	a. Checkout the repository developed in the previous section (Development).

	b. Build Docker image(s) for the application developed in the previous section.

	c. Push it to a Docker Hub

## General description

Docker, docker-compose, curl should be preinstalled.

Used Docker version 19.03.1 and docker-compose version 1.24.0.

Golang REST API application, stores data in PostgreSQL and use Redis as cache - on POST requests application checks the Redis before store the data in PostgreSQL. 

Application returns JSON in format: 

```
{"status" : boolean, "message" : "text", "data": "text", "timestamp": "timestamp"}
```

Fields `"data"` and `"timestamp"` are optional. And `"timestamp"` returns on POST request only.

#### Directories:

`app` - Golang application.

`data` - Redis and PostgreSQL data.

`services` - files for Jenkins server running (description is below). Pipeline job is in `services/jenkins/config/jobs.groovy` file.

## How to deploy

The Docker image can be rebuilt by (multi-stage building)

for development:

```
docker build --target builder -t alrf/go-redis-pg:latest .
```

for production:

```
docker build --target app -t alrf/go-redis-pg:latest .
```

Deploy application:

```
docker-compose up -d 
```
 
It will run go, redis, pg, jenkins containers and available as (shows last 10 records in DB, if they exist): 

http://127.0.0.1:8080/humai/echoservice

Stop application:

```
docker-compose down
```

Jenkins with Build_job (from CI Pipeline section) is available as:

http://127.0.0.1:8090/job/Build_job

Pipeline job itself is in `services/jenkins/config/jobs.groovy` file.

### Testing

Request Body Structure for Post Request:

```
{"department": "Dep1", "section": 1, "equipment": "Sect1", "description": "Descr1"}
```

`Makefile` - contains curl commands to send.

`make test-post` - send POST requests to populate data.

Example:

```
$ make test-post
for i in {1..20}; do \
	curl -X POST -H 'Content-Type: application/json' -d "{"\"department"\": "\"Dep$i"\", "\"section"\": $i, "\"equipment"\": "\"Sect$i"\", "\"description"\": "\"Descr$i"\"}" "http://127.0.0.1:8080/humai/echoservice"; \
done
{"data":{"department":"Dep1","section":1,"equipment":"Sect1","description":"Descr1"},"message":"Inventory has been created","status":true,"timestamp":"2019-09-22T09:23:09.044023317Z"}
{"data":{"department":"Dep2","section":2,"equipment":"Sect2","description":"Descr2"},"message":"Inventory has been created","status":true,"timestamp":"2019-09-22T09:23:09.104060885Z"}
{"data":{"department":"Dep3","section":3,"equipment":"Sect3","description":"Descr3"},"message":"Inventory has been created","status":true,"timestamp":"2019-09-22T09:23:09.154428129Z"}
{"data":{"department":"Dep4","section":4,"equipment":"Sect4","description":"Descr4"},"message":"Inventory has been created","status":true,"timestamp":"2019-09-22T09:23:09.20499422Z"}
{"data":{"department":"Dep5","section":5,"equipment":"Sect5","description":"Descr5"},"message":"Inventory has been created","status":true,"timestamp":"2019-09-22T09:23:09.28245457Z"}
```

`make test-get` - send GET requests to retrieve data.

Example:
```
$ make test-get
curl "http://127.0.0.1:8080/humai/echoservice";
{"data":[{"department":"Dep20","section":20,"equipment":"Sect20","description":"Descr20"},{"department":"Dep19","section":19,"equipment":"Sect19","description":"Descr19"},{"department":"Dep18","section":18,"equipment":"Sect18","description":"Descr18"},{"department":"Dep17","section":17,"equipment":"Sect17","description":"Descr17"},{"department":"Dep16","section":16,"equipment":"Sect16","description":"Descr16"},{"department":"Dep15","section":15,"equipment":"Sect15","description":"Descr15"},{"department":"Dep14","section":14,"equipment":"Sect14","description":"Descr14"},{"department":"Dep13","section":13,"equipment":"Sect13","description":"Descr13"},{"department":"Dep12","section":12,"equipment":"Sect12","description":"Descr12"},{"department":"Dep11","section":11,"equipment":"Sect11","description":"Descr11"}],"message":"success","status":true}
```
