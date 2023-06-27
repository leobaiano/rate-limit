# Rate Limit

This repository is a rate limit implementation study, the idea is to create API and Clients in a few different languages to understand the difference between them and how they behave.

## About Rate Limit

Rate Limit is a good security practice for the API's, basically you define how many requests, originating from the same IP address, you want to accept in a time interval before blocking. The idea is to identify hostile traffic sources based on how often they are trying to access a resource.

## How to use

### Up the server

1. Access the `server` directory
2. Run `go mod tidy` command to install dependencies
3. Run the `go build` command to build the application
4. In the root of the project, execute the command `./hello_world.exe` to start the server
5. Access the API in the browser via the URL `http://localhost:8080/hello`

If everything went well you should get a `Hello, World!` return message.