#!/usr/bin/env bash
ssh -T root@captcha.mojotv.cn <<'ENDSSH'
    cd $GOPATH/src/github.com/mojocn/base64Captcha
    git reset --hard
    git pull
ENDSSH

#http://www.cnblogs.com/ilfmonday/p/ShellRemote.html
#ssh root@MachineB 'bash -s' < local_script.sh
#https://stackoverflow.com/questions/305035/how-to-use-ssh-to-run-a-shell-script-on-a-remote-machine
#ssh root@MachineB 'bash -s' < local_script.sh