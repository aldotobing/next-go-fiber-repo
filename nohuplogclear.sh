#!/bin/bash

# Set the file path and command to execute
FILE_PATH="/home/ec2-user/nextbasis-service-golang/server/nohup.out"
COMMAND="cat /dev/null > /home/ec2-user/nextbasis-service-golang/server/nohup.out"

# Check the file size in a loop
while true; do
    # Get the file size in bytes
    FILE_SIZE=$(stat -c%s "$FILE_PATH")

    # Check if the file size is greater than or equal to 1 gigabyte (1073741824 bytes)
    if [ $FILE_SIZE -ge 1073741824 ]; then
        # Execute the command
        $COMMAND

        # Exit the loop
        break
    fi

    # Sleep for 1 second before checking again
    sleep 1
done 