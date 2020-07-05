## go_mongodb



The project consists of MongoDB and Golang code that connects to the database. The project  packaged in Docker and consists of three services: mongo, mongo-express (GIU ) and api. Golang code connects to the database, creates and reads a fixed structure.  

Start: `bash build.sh`

Test:  `curl http://localhost:8080/v1/url.insert`

Mongo-express `http://localhost:8081`

Stop: `bash purge.sh`

***

## go_mongodb


Проект состоит из Mongo DB и кода на Go , подключающийся к базе. Проект упакован в Docker и  состоит из трех сервисов : mongo, mongo-express ( GIU ), api. Gо код соединяется с базой, создает и читает некую фиксированную структуры.

Start: `bash build.sh`

Test:  `curl http://localhost:8080/v1/url.insert`

Mongo-express `http://localhost:8081`

Stop: `bash purge.sh`
