#!/usr/bin/awk

# This program parses "nmap -oG ..." output and
# prints results to files per each host.
# It could have been embedded into scanner.sh, but left as a separate
# file for convenience (syntax highlighting in editors).

# Examine each line starting with "Host: "
/^Host: / {

	# Touch new file per host, just in case there are no open ports at all
	host=$2
	system("touch "host"-new.log")

	# Skip lines we are not interested in, i.e. not containing list of ports
	if ($4 != "Ports:") {
		next
	}

	ports=""
	# In nmap greppable output, ports start to be listed starting from 5th field
	# add all of them into ports variable; they would be delimited with ","
	for (f = 5; f <= NF; f++) {
		# Discard all fields starting from the one with word "Ignored"
		if ($f ~ /^Ignored/) {
			break
		}
		ports=ports $f
	}
	split(ports,portsArray,",")

	# Print one line per port to the file
	for (port in portsArray) {
		printf "Host: %s Ports: %s\n", host, portsArray[port] > host"-new.log"
	}

}
