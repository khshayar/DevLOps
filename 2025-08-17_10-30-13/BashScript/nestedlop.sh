#!/bin/bash

a=0
while [ $a -le 10 ]
do
	b=$a
	while [ $b -gt 0 ]
	do
		echo -n "$b  "
		b=$(( b-1 ))
	done
	echo 
	a=$(( a+1))
	
done 
 
