# SSH_ASKPASS=./askpass ssh -F /dev/null -o PreferredAuthentications=password pi@raspberrypi.local mkdir .ssh
# SSH_ASKPASS=./askpass scp -F /dev/null -o PreferredAuthentications=password $HOME/.ssh/id_rsa.pub pi@raspberrypi.local:.ssh/authorized_keys
# password='raspberry'

ssh_with_password() {
  expect -c "
  spawn $1
  expect \"pi@raspberrypi.local's password: \"
  sleep 1
  send \"$2\n\"
  interact
  "
}

ssh_with_password "ssh -F /dev/null -o PreferredAuthentications=password pi@raspberrypi.local mkdir -p .ssh" raspberry
ssh_with_password "scp -F /dev/null -o PreferredAuthentications=password $HOME/.ssh/id_rsa.pub pi@raspberrypi.local:.ssh/authorized_keys" raspberry
