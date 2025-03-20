#!/bin/sh

# Check if the address-book process is running
if pgrep address-book > /dev/null; then
    # Check if the data directory exists and is writable
    if [ -d "/app/data" ] && [ -w "/app/data" ]; then
        exit 0
    fi
fi

exit 1 