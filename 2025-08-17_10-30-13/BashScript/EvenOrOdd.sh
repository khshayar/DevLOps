#!/bin/bash
numbers="1 2 3 4 5 6 7 9 8 10"

for num in $numbers 
do 
	q=$(( num % 2))
	if [ $q -eq 0 ]
	then
		echo "this is even $num "
		continue
	fi
	echo "this is odd $num"
done 
