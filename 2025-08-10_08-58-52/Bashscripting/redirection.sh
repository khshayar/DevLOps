#!/bin/bash

#this example for shoing stdout and std erorr
#apt update > stdout.txt
apt update &>stdout.txt
#this is optional if you want deleting stdout and std err use this command 
########################### apt update & > /dev/null

