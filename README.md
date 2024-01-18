# Mailing Service

This is an example customer mailing service. 
It exposes simple REST-like API for adding, deleting and sending emails. 

# How to run

docker-compose up 

# Examples

#### Add new email message:
`
curl -X POST localhost:8080/api/messages -H 'Content-Type: application/json' -d '{
    "email": "bozenka.kowalska@example.com",
    "title": "urgent email",
    "content": "simple text",
    "mailing_id": 1,
    "insert_time": "2021-02-24T01:42:38Z"
}'
`

#### Deleting all messages older than 5 minutes by mailing id 
`curl -X DELETE localhost:8080/api/messages/{id}`
