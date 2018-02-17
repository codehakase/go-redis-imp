# go-redis-imp
A simple implementation of Redis key value storage in Go

## WORK IN PROGRESS
At this point, this is just a program I spinned off in an evening. It basically implements the basics operations on Redis (GET, SET, DELETE).

## Why?
If nothing else, this code helps me understand Redis better.

## Run this
```shell
$ git clone https://github.com/codehakase/go-redis-imp $GOPATH/src/github.com/go-redis-imp

$ cd $GOPATH/src/go-redis-imp

$ go build

$ ./go-redis-imp
.... some debugging crap ...

```

On another terminal connect via telnet
```shell
$ telnet localhost 1234
```

Run basic redis commands:
```$
SET alias codehakase

GET alias

...-> codehakase

DEL alias

```
