# REST API for Student CRUD operations

## Technologies used:
1. `Go`
2. `Cassandra`
3. `Docker`

## Setting up DB
Run the following commands
1. `docker pull cassandra`
2. `Start the docker container and enter into CLI`
3. `CREATE KEYSPACE connectwise WITH replication = {'class':'SimpleStrategy', 'replication_factor' : 3};`
4. `USE KEYSPACE connectwise;`
5. `CREATE TABLE student_db(id uuid, age int, class int, deleted boolean, name text, subject text, primary key (id));`
6. `Navigate into cloned folder`
7. `go run main.go`