#!/bin/bash

# Navigate to the project directory
cd /home/ubuntu/site-portfolio

# Load environment variables (if required)
export $(cat .env | xargs)

# Ensure the log directory exists, and create the log file if necessary
LOG_FILE="/var/log/myservice.log"
if [ ! -f "$LOG_FILE" ]; then
    sudo touch "$LOG_FILE"
    sudo chmod 666 "$LOG_FILE"  # Grant read/write permissions to all users (adjust as needed)
fi

# Start the Go service in the background
nohup /usr/bin/go run cmd/main.go > "$LOG_FILE" 2>&1 &

# Output the PID of the Go process for reference
echo "Portfolio service started with PID: $!"

#ps aux | grep "go run cmd/main.go"

#ps aux | grep go
#ubuntu     33188  0.0  2.0 1609192 20480 pts/0   Sl   22:06   0:00 /tmp/go-build2550502408/b001/exe/main
#ubuntu     34023  0.0  0.2   7076  2048 pts/0    S+   22:40   0:00 grep --color=auto go
