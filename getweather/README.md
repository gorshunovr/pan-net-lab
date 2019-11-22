# The getweather application

## Programming

This is a simple containerized application.

Given variables listed below, it fetches and prints to stdout data from
[openweathermap.org API](https://openweathermap.org/api), specifically from
[current weather data](https://openweathermap.org/current) API, using specified
format.

* `CITY_NAME` (e.g. `Honolulu`)
* `OPENWEATHER_API_KEY` (e.g. `1234567890abcdef`)

There are two versions of application included:

* BusyBox Shell script
* Go application

Shell script application essentially calls `curl` and parses JSON with `jq`
command. This applicaiton appropriately handles absence of required variables,
errors running `curl` while fetching API URL, and `jq` errors while parsing JSON
data. Shell script supports Bash and BusyBox shell syntax.

Application written in Go utilises
[github.com/vascocosta/owm](https://github.com/vascocosta/owm) library. It
creates new client interface, calls `WeatherByName`, and prints out necessary
values using predefined format. This applications appropriately handles absence
of `CITY_NAME`, but fails to appropriately handle problems with incorrect or
missing `OPENWEATHER_API_KEY`; this is due to the HTTP code 401 Unauthorized not
being handled properly by the library. This can be improved, of course.

Both versions of applications use no command line arguments.

```code
getweather/
├── app.go
├── app.sh
├── Dockerfile
├── Dockerfile.shell
└── README.md
```

## Build container

By default, Go application container is being built. Use `Dockerfile.shell` to
build Shell script -based container.

Run as root:

```bash
docker build -t getweather:1.0 .
```

```bash
docker build -t getweather:1.0 -f Dockerfile.shell .
```

Both containers are based on `alpine:3.10` (specific version could be overridden
by providing `FROM=xxx` variable during build).

Resulting container sizes are just about 8.5MB for shell-based, and about 13.5MB
for Go application container.

Go application container utilizes multi-stage build to reduce image size from
builder image size of approximately 390MB.

## Run container

Run as root:

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
docker run \
   -e CITY_NAME="${CITY_NAME}" \
   -e OPENWEATHER_API_KEY="${OPENWEATHER_API_KEY}" \
   getweather:1.0
```

## Run as standalone application

Run as user:

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
go get github.com/vascocosta/owm
go build app.go
./app
```

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
bash ./app.sh
```

## Output example

```bash
source=openweathermap, city="Honolulu", description="light rain", temp=23.72, humidity=78
```
