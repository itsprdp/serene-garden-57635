# Simple HTTP TimeZone Server
Returns current time for the given time_zone. By default it returns the UTC time.

### Live Instance:
https://serene-garden-57635.herokuapp.com/

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

$ curl https://serene-garden-57635.herokuapp.com/\?time_zone\="Asia/Kolkata"
{"CurrentTime":"20:25:16","TimeZone":"Asia/Kolkata","Message":"ok"}
```
