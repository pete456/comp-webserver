#!/bin/bash
user=$1
pass=$2
useradd -b /webserver/web/usersites -m -p $pass -s "/bin/bash" -U $user
