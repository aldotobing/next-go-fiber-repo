#!/bin/bash

ENTER_DIR="Entering project dir"
echo $ENTER_DIR
cd /home/ec2-user/nextbasis-service-golang && echo OK || echo Failed
echo "Killing running GO process on port 5000 ..."
sleep 2
sudo kill -9 $(sudo lsof -t -i:5000) && echo "Proc Killed!" || echo Failed
sleep 1
echo "Activating SSH Agent ..."
eval `ssh-agent -s` && echo OK || echo Failed
ssh-add ~/.ssh/deployssh && echo OK || echo Failed
sleep 2
echo "entering project dir ..."
cd /home/ec2-user/nextbasis-service-golang
pwd && echo OK || echo Failed
echo "Pulling from git repository ..." 
sleep 1
git pull -v origin rebuild && echo OK || echo Failed
echo "Entering server directory to start main.go" && echo OK || echo Failed
sleep 2
cd /home/ec2-user/nextbasis-service-golang/server/
pwd && echo OK || echo Failed
echo "Starting Service ..."
sleep 2
nohup bash -c "go run main.go 2>&1 &"
echo $?
exit 0
