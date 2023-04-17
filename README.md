# ohlc
This is ohlc server side and client side
1. Client side is in client-service folder
2. Server side is in service folder

- Redis Folder & func
    This function is for set config and inisiate the redis client

- Kafka Foler & func
    This function is for set config and inisiate the kafka client

How to run:
- Server Side (service):
1. type cd service on the terminal
2. type go run .
3. Test service health: `GET http://localhost:3000/health`

- Client side (client-service):
1. type cd client-service on the terminal
2. type go run .
3. Test service health: `GET http://localhost:3001/health`

    

