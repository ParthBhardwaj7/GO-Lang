# Go Server Monitoring Stack

A simple Go web server hooked up with Prometheus and Grafana to monitor 
real-time traffic, errors, and response times. Built this to learn how 
observability actually works in production systems.

## What's inside

- Go HTTP server with `/hello` and `/form` endpoints
- Chaos engine that randomly throws 500 errors and slow responses
- Prometheus scraping metrics every 5 seconds
- Grafana dashboards showing live request rates, errors, response times
- Everything runs in Docker Compose — one command and you're up

## Stack

- Go 1.23
- Prometheus
- Grafana
- Docker + Docker Compose

## Running locally

Make sure Docker Desktop is running, then:

```bash
git clone https://github.com/ParthBhardwaj7/go-monitoring
cd go-monitoring
docker-compose up --build
```

Then open:
- App → http://localhost:8080/hello
- Metrics → http://localhost:8080/metrics
- Prometheus → http://localhost:9090
- Grafana → http://localhost:3000

Grafana login is admin / admin

## Grafana Setup

1. Add Prometheus as data source → URL: http://prometheus:9090
2. Create a new dashboard
3. Add panels with these queries:

Request rate:
rate(http_requests_total[1m])

Error rate:
rate(http_errors_total[1m])

Response time:
rate(http_request_duration_seconds_sum[1m])

## Why I built this

Wanted to understand how monitoring actually works beyond just reading 
about it. The chaos engine makes it interesting — the server randomly 
fails and slows down so you can actually see alerts trigger and graphs 
spike in real time.

## Project structure
go-monitoring/
├── main.go
├── go.mod
├── Dockerfile
├── docker-compose.yml
├── prometheus.yml
└── static/
└── index.html