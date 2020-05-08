# icat-project-work

ICAT project work is a course in University of Vaasa
where students can do a software project and document the work.


In this project an arduino is used to measure blood pressure and other health related measurements, and these measurements are displayed in a web application. The data from arduino is sent to a web server, and the web client displays data from server.

The architecture of this software project uses microservices. The API microservices are built using golang, the web client using react and database is a NoSQL MongoDB database. All services are running in Docker containers.

## Requirements
- Docker
- docker-compose

## Development
`docker-compose up --build`
