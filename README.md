# Cards Game Test

## Required Tools
This program is created with Go 1.20.1 version, and is not tested in previous versions, so no backwards compatibility is guaranteed.

To install the latest version of Go for your system please visit:
https://go.dev/doc/install

To check if you have "Go" correctly installed, open a terminal window and type:
```
go version
# Should show something like
# go version go1.20.1 linux/amd64
```

## Download the example

Execute in your terminal:
```
git clone https://github.com/ferbal02/testcardsgame.git
```

## Setup program

- cd into the newly downloaded repo folder
```
cd ~/testcardsgame
```
- Download Go missing dependencies
```
go get
```

## Run the service

The program is a microservice with an embedded web server (Gin framework) and a manual testing UI (Swagger Open API).

By default the service will be listening on "localhost:8080" (This port should not be in use to run this service)

In the repo folder type the following command to start the service:
```
git run main.go
```
Important urls:
- localhost:8080 -> Will show a simple presentation page
- localhost:8080/swagger/v1 -> Swagger UI. Allows to use and test the API from a web interface.
- localhost:8080/api/v1 -> Root api url.

## Run Unit Tests

In order to run all test cases and see the status for each of them, type:
```
go test -v test/cardsgame/tests/controllers
```
No testing errors should be displayed

## Swagger UI
To use the swagger UI, run the service and open this URL with your browser "http://localhost:8080/swagger/v1"

**The interface only works from localhost, as no CORSS issues have been considered**

The UI will show you a description of the 3 operations available with a friendly interface to interact with it.
- /deck -> Create new deck  and returns the new deck as reponse. (POST request)
- /deck/{uuid} -> Returns the requested Deck if exists, otherwise returns error. (GET request)
- /deck/{uuid}/cards -> Returns as many cards from a deck as requested. If deck not found or too many cards requested returns error.

## Improvements
Due to the expected excercise time, there are some improvements that I would add to the program in normal conditions:
- Unit Test or End2End test for the endpoints by mocking the web server requests to validate parameter conversions, possible user authentication...
- Unit Test conversions between API objects (DTO) and business logic objects
- As I have used a "Memory" database (MemoryDeckRepository class) , I have not added concrete tests for the persistence layer, but
in case of using a real SQL database, this could also be tested.


