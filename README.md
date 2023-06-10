## go_mongodb

![Docker Compose Actions Workflow](https://github.com/dxavx/go_mongodb/workflows/Docker%20Compose%20Actions%20Workflow/badge.svg?branch=master)

The project consists of MongoDB and Golang code that connects to the database. The project  packaged in Docker and consists of three services: mongo, mongo-express (GIU ) and api. Golang code connects to the database, creates and reads a fixed structure.  

Start: `bash build.sh`

Test:  `curl http://localhost:8080/v1/url.insert`

Mongo-express `http://localhost:8081`

Stop: `bash purge.sh`

***
