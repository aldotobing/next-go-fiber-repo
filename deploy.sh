#!/bin/bash

ENTER_DIR="Entering project dir cd /home/ec2-user/nextbasis-service-golang/"
echo $ENTER_DIR
KILL_PROC="kill $(lsof -t -i:5000)"
echo $KILL_PROC
sudo kill -9 $(sudo lsof -t -i:5000)
echo "Proc Killed!"
eval `ssh-agent -s`
ssh-add ~/.ssh/deployssh
cd /home/ec2-user/nextbasis-service-golang/
pwd
echo "Pulling from git"
git pull -v origin rebuild
echo "Enter server directory to start main.go"
cd /home/ec2-user/nextbasis-service-golang/server/
pwd
nohup bash -c "go run main.go 2>&1 &"
echo $?
exit 0
