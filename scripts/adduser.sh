#!/bin/bash
user=$1
pass=$2
home=$3
useradd -b $home -m -p $pass -s "/bin/bash" -U $user
