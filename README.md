# Path Finder

This is a microservice API that can help you understand and track a person's flight path given a list of flights. The API accepts a request in the form of a JSON array that includes a list of flights. Each flight is defined by a source and destination airport code. The flights may not be listed in order and may contain circular flights.

## API Endpoint

### Request

To request the flight path for a person, you can send a POST request to the following endpoint:

```
POST /calculate
```

The request body should be a JSON array of flights, where each flight is represented as a JSON array with two strings, the source and destination airport codes:

```json
[
  ["IND", "EWR"],
  ["SFO", "ATL"],
  ["GSO", "IND"],
  ["ATL", "GSO"]
]
```

### Response

If the request is successful, the API will respond with a JSON array of strings representing the flight path starting and ending airports:

```json
["SFO", "EWR"]
```

If the request body is invalid or the server encounters an error, the API will respond with an appropriate HTTP status code and an error message in the response body.

## Development

To start the server, run the following command:

```
go run main.go
```

The server will listen on port 8080 by default.
