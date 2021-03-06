# The getweather application

## Programming

This is a simple containerized application.

Given variables listed below, it fetches and prints to stdout data from
[openweathermap.org API](https://openweathermap.org/api), specifically from
[current weather data](https://openweathermap.org/current) API, using specified
format.

* `CITY_NAME` (e.g. `Honolulu`)
* `OPENWEATHER_API_KEY` (e.g. `1234567890abcdef`)

There are three versions of application included:

* BusyBox Shell script
* Go application
* Standalone Go application

```code
getweather/
├── v2/
│   ├── app.go
│   ├── appStructs.go
│   └── Dockerfile
├── app.go
├── app.sh*
├── Dockerfile
├── Dockerfile.shell
└── README.md
```

Shell script application essentially calls `curl` and parses JSON with `jq`
command. This application appropriately handles absence of required variables,
errors running `curl` while fetching API URL, and `jq` errors while parsing JSON
data. Shell script supports Bash and BusyBox shell syntax.

Application written in Go utilizes
[github.com/vascocosta/owm](https://github.com/vascocosta/owm) library. It
creates new client interface, calls `WeatherByName`, and prints out necessary
values using predefined format. This application appropriately handles absence
of `CITY_NAME`, but fails to appropriately handle problems with incorrect or
missing `OPENWEATHER_API_KEY`; this is due to the HTTP code *401 Unauthorized*
not being handled properly by the library. This can be improved, of course.

Standalone Go application in `v2/` subdirectory has been added to this repository
later. It's self-contained: does not use 3rd party libraries, and resulting
container is based on `scratch` image. This application appropriately handles
absence of both `CITY_NAME` and `OPENWEATHER_API_KEY`.

All versions of applications use no command line arguments.

## Build container

By default, Go application container is being built. Use `Dockerfile.shell` to
build Shell script -based container. The `--rm` flag removes intermediate
containers on successful build.

Although examples here use Docker, same `Dockerfile`'s and containers are
verified to be working well with Podman (for RHEL8, latest OpenShift, and
latest Fedora versions). Just use `podman` instead of running `docker` in
commands below.

Run as root:

```bash
docker build --rm -t getweather:1.0 .
```

```bash
docker build --rm -t getweather:2.0 v2/
```

```bash
docker build --rm -t getweather:1.0 -f Dockerfile.shell .
```

Two containers are based on `alpine:3.10`, standalone container is based on
`scratch` image (specific version could be overridden by providing `FROM=xxx`
and `FROMBLD=xxx` variables during build with `--build-args` commandline key).

Resulting container sizes are just about 7.99MB for shell-based, 7.51MB for
standalone Go, and about 12.9MB for Go application container.

Both Go application containers utilize multi-stage build to reduce image size
from builder image size of approximately 390MB.

## Run container

Run as root (add `-d` parameter after `run` to avoid output to stdout, message
would still be delivered to syslog; use different tag if needed):

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
docker run --rm \
   -e CITY_NAME="${CITY_NAME}" \
   -e OPENWEATHER_API_KEY="${OPENWEATHER_API_KEY}" \
   getweather:1.0
```

## Run as standalone application

Run as user (assuming Go is installed):

Go app:

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
go get github.com/vascocosta/owm
go build app.go
./app
```

Standalone Go app:

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
go build -o app
./app
```

Shell script:

```bash
declare -x OPENWEATHER_API_KEY="xxxxxxxxxxxx"
declare -x CITY_NAME="Honolulu"
bash ./app.sh
```

## Output example

```code
source=openweathermap, city="Honolulu", description="light rain", temp=23.72, humidity=78
```

Output could also be seen in syslog messages file:
```bash
vagrant@n0:/vagrant/scanner$ sudo grep openweathermap /var/log/syslog
Nov 22 15:30:21 localhost 073e2de20629[6894]: source=openweathermap, city="Honolulu", description="light rain", temp=25.94, humidity=61
```
