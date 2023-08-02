#!/bin/bash
set -e

main() {
  ENTER_DIR="Entering project dir"
  echo $ENTER_DIR

  cd /home/ec2-user/nextbasis-service-golang && echo OK || { echo "Failed: $ENTER_DIR"; exit 1; }

  echo "Killing running GO process on port 5000 ..."
  sleep 2
    if [ -n "$(sudo lsof -t -i:5000)" ]; then
      sudo kill -9 $(sudo lsof -t -i:5000) && echo "Proc Killed!" || { echo "Failed: Killing running GO process"; exit 1; }
  else
    echo "No process to kill on port 5000. Continuing..."
  fi
  sleep 1
  echo "Activating SSH Agent ..."
  eval `ssh-agent -s` && echo OK || { echo "Failed: Activating SSH Agent"; exit 1; }

  ssh-add ~/.ssh/deployssh && echo OK || { echo "Failed: Adding SSH key"; exit 1; }

  sleep 2
  echo "entering project dir ..."
  cd /home/ec2-user/nextbasis-service-golang && echo OK || { echo "Failed: Entering project dir"; exit 1; }
  pwd

  echo "Pulling from git repository ..." -
  sleep 1
  git pull -v origin rebuild && echo OK || { echo "Failed: Pulling from git repository"; exit 1; }

  echo "Entering server directory to start main.go" 
  sleep 2
  cd /home/ec2-user/nextbasis-service-golang/server && echo OK || { echo "Failed: Entering server directory"; exit 1; }
  pwd 

  echo "Clearing nohup.out log file ..."
  cat /dev/null > nohup.out && echo Cleared! || { echo "Failed: Clearing nohup.out log file"; exit 1; }

  echo "Starting GO Service ..."
  sleep 2
  nohup bash -c "go run main.go 2>&1 &"

  echo $?
  exit 0
}

handle_error() {
  local exit_code="$?"
  local command="$BASH_COMMAND"
  echo "Error: Command '$command' exited with status $exit_code"
  echo "Stack Trace:"
  while caller $((n++)); do :; done
  exit "$exit_code"
}

trap 'handle_error' ERR

main "$@"
