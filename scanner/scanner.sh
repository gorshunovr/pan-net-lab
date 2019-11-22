#!/usr/bin/env bash

# This program provides a way to run repetitive network scans displaying
# differences between subsequent scans.
# * Target of the scan must be provided as CLI argument
# * Target can be single IP address as well as network range
#
# Since program uses nmap internally, you can provide different nmap options as
# you need by setting NMAP_OPTS variable.
# Parsing of the scan log files is done by awk script parser.awk.

set -e

# Run scan with TCP CONNECT scan, it's a fast scan option,
# this does not require root priveleges.
# By default, nmap would scan 1,000 popular ports.
: "${NMAP_OPTS:="-sT"}"

if [[ -z "${1}" ]]; then
    echo "ERROR: No target defined" >&2; exit 1
fi
target="${1}"

mkdir -p logs/; cd logs/

# Run nmap
nmap -oG last_scan_results.log ${NMAP_OPTS} "${target}" > /dev/null

# Parse full_scan_results.log file; create per-host files
awk -f ../parser.awk last_scan_results.log

for logfile in *-new.log; do

    [[ -e "$logfile" ]] || break

    # Extract IP from filename
    ip="${logfile%%-new.log}"
    # Old logs file name
    oldlogfile="${ip}".log

    # Compare current and previous scan results and print results
    if cmp --silent "${logfile}" "${oldlogfile}"; then
        echo "*Target - ${ip}: No new records found in the last scan.*"
    else
        echo "*Target - ${ip}: Full scan results:*"
        cat "${logfile}"
    fi

    # Replace previous scan results with current scan results
    mv "${logfile}" "${oldlogfile}"

done

cd ../
