# microservice-common
A common module for go microservices.

TODO: Create NewMongoClient and NewRedisClient, can you make struct initializer private?

How to run tests: *Must have local redis server*
 1) create .env in /test
 2) set mongo_uri variable
 3) go to root folder and exec go test -v ./test/
