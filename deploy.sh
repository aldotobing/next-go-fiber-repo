#!/bin/bash

sudo -i
ENTER_DIR="cd /home/ec2-user/nextbasis-service-golang/"
echo $ENTER_DIR
KILL_PROC="kill $(lsof -t -i:5000)"
echo $KILL_PROC
kill -9 $(lsof -i:5000 -t)
echo "Proc Killed!"
cd /home/ec2-user/nextbasis-service-golang/
echo "Pulling from git"
sudo git pull -v origin rebuild
echo "Enter server directory to start main.go"
cd /home/ec2-user/nextbasis-service-golang/server
sudo nohup go run main.go &
echo "SERVER'S UP!!!!!!"
