## Prometheus metrics example tutorial

### Settings

metrics endpoint settings
```yml
prometheus:
  api:
    port: 8092
    path: /metrics
```

metrics writter settings
```yml
metric:
    enabled: true
    interval: 60s
    writers: [prom]
```


### Run instructions

run `go run main.go`

The application exposes the prometheus metrics on `http://localhost:8092/metrics`
Additionally there is a basic api server running on port `:8088` with the following endpoints

### `GET http://localhost:8088/current-value`
return the current value of the counter
### `GET http://localhost:8088/increase`
increases the counter by one and returns the current value
### `GET http://localhost:8088/decrease`
decreases the counter by one and returns the current value

all the above endpoints also have a custom metric `api_request` which increments
every time an endpoint is called (the handler name is used as a label)


example output of custom metrics when requesting the `/metrics` endpoint
```
# HELP gosoline:dev:metrics:prometheus_api_request 
# TYPE gosoline:dev:metrics:prometheus_api_request counter
gosoline:dev:metrics:prometheus_api_request{handler="cur"} 10
gosoline:dev:metrics:prometheus_api_request{handler="decr"} 5
gosoline:dev:metrics:prometheus_api_request{handler="incr"} 10
# HELP gosoline:dev:metrics:prometheus_important_counter 
# TYPE gosoline:dev:metrics:prometheus_important_counter gauge
gosoline:dev:metrics:prometheus_important_counter 5
```
