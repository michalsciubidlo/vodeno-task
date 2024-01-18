# Mailing Service

This is an example customer mailing service. 
It exposes simple REST-like API for adding, deleting and sending emails. 
It uses postgreSQL as database.

## How to run

### Running migrations:
Using: https://github.com/golang-migrate/migrate


`migrate -database postgres://postgres:example@localhost:5432/api?sslmode=disable -path db/migrations up`

### Launch the application & database:
 `make launch`

## Examples

#### 1. Add new email message:
`
curl -X POST localhost:8080/api/messages -H 'Content-Type: application/json' -d '{
    "email": "bozenka.kowalska@example.com",
    "title": "urgent email",
    "content": "simple text",
    "mailing_id": 1,
    "insert_time": "2021-02-24T01:42:38Z"
}'
`

#### 2. Deleting all messages older than 5 minutes by mailing id 
`
curl -X DELETE localhost:8080/api/messages/{id}
`

#### 3. Sending messages (simulated)
`
curl -X POST localhost:8080/api/messages/send -H 'Content-Type: application/json' -d '{"mailing_id":120}'
`
