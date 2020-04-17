### Exponential Backoff
In a distributed system, many operations have a chance of failure. For example, networks might be down, buffers might be full, or services might be restarting.

However, we can make the system more robust by introducing ”exponential backoff”, in which we retry failing operations with increasing ”backoff” periods (wait 1 second, then wait 2 seconds, then wait 4 seconds, etc.) until we either succeed or give up.

In addition, we can add ”jitter” to randomize the backoff periods and prevent ”thundering herd” effects (e.g. wait a random period between 0.5 and 1.5 seconds, instead of exactly 1 second).

### Usage:
```
$ go run main.go

Execute with BackOff Strategy:
Trying to perform the operation ...
Failed to say hello to the world!
wait for: 1 seconds before retrying
Trying to perform the operation ...
Failed to say hello to the world!
wait for: 3 seconds before retrying
Trying to perform the operation ...
Failed to say hello to the world!
wait for: 7 seconds before retrying
Trying to perform the operation ...
Failed to say hello to the world!
wait for: 15 seconds before retrying
Trying to perform the operation ...
Failed to say hello to the world!
Failed: Max no. of attempts: 5 . Total duration: 31 seconds
exit status 1
```
