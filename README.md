# interview1-assessment

## Project structure
This project follow the guidance of Golang project layout: https://github.com/golang-standards/project-layout

Business logic will be inside internal folder and necessary folders to store business definition will be there.

## Assessment

### Context

We have a website on the Internet and we would like to get some very simple indication of how visitors navigate the pages.

For that purpose, we managed to configure our website to send an event every time a visitor navigates to a page. Our website is capable of generating unique identifiers for visitors as a string of characters.

The system generating that event is able to talk to a REST HTTP interface and represents each individual event as a JSON document containing two attributes: the unique identifier of the visitor and the URL of the visited page.

Our product team is starting a new sprint. We are picking the following user story:
**As a digital marketeer, I need to know how many distinct visitors navigated to a page, knowing its URL.**
 
### Task
Build a GoLang web service capable of:
* Ingesting user navigation JSON events via a REST HTTP endpoint. Each event is to be ingested via a separate HTTP request (i.e. no batch and no streaming ingestion)
* Serving the number of distinct visitors for any given page via another REST HTTP endpoint. The page URL we are interested in should be a query parameter of the HTTP request. The number of distinct visitors for that URL is returned as a number in plain text.

**REST Interface**

* Register a new tracking event
```curl --location --request POST 'http://{hosot}/tracking/events' \
--header 'Content-Type: application/json' \
--data-raw '{
    "uid": "user-unique-identifier",
    "url": "http://localhost:8080/tracking/events?bar=2&foo=1"
}' -v
```

* Get number of visits from a given url
```
curl http://{host}/tracking/metrics\?url\=http%3A%2F%2Flocalhost%3A8080%2Ftracking%2Fevents%3Ffoo%3D1%26bar%3D2 -v
```
* Notice previous url value needs to be encoded previously:
* * url to get data from: http://localhost:8080/tracking/events?bar=2&foo=1
* * url encoded to be compatible with http protocol: http%3A%2F%2Flocalhost%3A8080%2Ftracking%2Fevents%3Ffoo%3D1%26bar%3D2


### Constraints
* There is no need for persistence to a database. Everything can be kept in memory.
* The web service must be capable of handling concurrent requests on both endpoints.


## Make it work
```
go install
go run cmd/api/main.go
```
