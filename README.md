# api-response
Standard API response of Microservices

## Map for Errors code
1. Create a const for each of the error code with the name of the errors. For example,
```go
const (
	ErrNameExists    int = 2002
	ErrInvalidParams int = 2003
	ErrDoesNotExist  int = 2004
)
```
2. Create a map with type map[int]response.ErrorStruct wherein mapping the error code to respective error message and http status code.
```go
var ErrorsMap = map[int]response.ErrorStruct{
	ErrNameExists:    {ErrorMsg: "entity with same name already exists", RespCode: http.StatusNotAcceptable},
	ErrInvalidParams: {ErrorMsg: "invalid incoming request", RespCode: http.StatusBadRequest},
	ErrDoesNotExist:  {ErrorMsg: "entity does not exist", RespCode: http.StatusNotFound},
}
```

## Return the response
In order to return the response in standard way, pass the gin context pointer along with the error Map, status(boo), data to be sent( nil in case where no data is sent in response) and error Code:
```go
response.ApiResponse(c, ErrorsMap, status, data, errCode)
```
This call will return the required response in json to the client.

