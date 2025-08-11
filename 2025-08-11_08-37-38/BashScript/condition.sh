#!/bin/bash
name=" "
age=0
echo "hi What is your name ?"

read -p "Enter your name: " name

read -p "hi $name, how old are you " age

if [ $age -gt 18 ]; then
    echo "You are an big."
elif [ $age -eq 18 ]; then
    echo "You just became an adult."
else
    echo "You are a minor."
fi
