#!/bin/ash
# ^^^^^^^ - this is BusyBox ash (A shell)

# Given variables listed below, fetches and prints data from
# openweathermap.org API using specified format
#   CITY_NAME
#   OPENWEATHER_API_KEY

# shellcheck shell=dash
# SC2169 note: BusyBox ash supports [[

set -e

# shellcheck disable=SC2169
if [[ -z "${CITY_NAME}" ]]; then
	echo "ERROR: CITY_NAME is not defined" >&2; exit 1
fi
# shellcheck disable=SC2169
if [[ -z "${OPENWEATHER_API_KEY}" ]]; then
	echo "ERROR: OPENWEATHER_API_KEY is not defined" >&2; exit 1
fi

PARAMS="?q=${CITY_NAME}&APPID=${OPENWEATHER_API_KEY}&units=metric"
URL="https://api.openweathermap.org/data/2.5/weather${PARAMS}"

json=$(curl --fail --silent --show-error "${URL}")

temp="$(jq -n --argjson json "${json}" '$json.main.temp')"
humidity="$(jq -n --argjson json "${json}" '$json.main.humidity')"
description="$(jq -n --argjson json "${json}" '$json.weather[].description')"

fmt="source=openweathermap, city=\"%s\", description=%s, temp=%.2f, humidity=%d\n"
# We do use variable to set formatting pattern here below
# shellcheck disable=SC2059
printf "${fmt}" "${CITY_NAME}" "${description}" "${temp}" "${humidity}"
