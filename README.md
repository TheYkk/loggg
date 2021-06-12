# Simple log dumper

This projects aims to generate sample log lines with json, text, with color.

You can use this project to test your log scraper like fluentbit or promtail

## Build

You can build and run with docker

```
docker build -t loggg .
docker run --rm loggg
```

or you can build with go

```
go mod download
go build main.go -o loggg
./loggg
```

Pre builded docker image
```
docker run --rm theykk/loggg
```
## Rate

Current rate is 130 log lines per second