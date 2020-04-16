# TimeZone server
Returns current time for the given time_zone. By default it returns the UTC time.

### Usage:
```
# Server
$ go run time_zone.go
...

# Client
$ curl "localhost:8080"
{"CurrentTime":"08:37:44","TimeZone":"UTC","Message":"ok"}

$ curl "localhost:8080/?time_zone=America/New_York"
{"CurrentTime":"04:37:54","TimeZone":"America/New_York","Message":"ok"}
```
