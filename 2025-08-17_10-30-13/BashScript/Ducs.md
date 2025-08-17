Bash script 

special variable 

$0 >>>>>>>>>>>>>>>> file name of script 
$#
$*
$@
$? ................. command excuted status 
$$

operators: 
+-*/ = !=

relations operators:

-eq >>>>>>>>>>>>>>>>>>equal =
-ne >>>>>>>>>>>>>>>>>>not equal !=
-gt >>>>>>>>>>>>>>>>>> greater than >
-le >>>>>>>>>>>>>>>>>> less than <
-ge >>>>>>>>>>>>>>>>>>> greater equal >=
-le >>>>>>>>>>>>>>>>>>> less equal <=
! >>>>>>>>>>>>>>>>>>>>>> not 
-o >>>>>>>>>>>>>>>>>>>>> or 
-a >>>>>>>>>>>>>>>>>>>>> and
=
!=
-z >>>>>>>>>>>>>>>>>>>>> check if the len off string is empty then is true
-n>>>>>>>>>>>>>>>>>>>>>> same above white none 
str >>>>>>>>>>>>>>>>>>>> if is an empty then give back false

file test operators

-b file  >>>>>> block special file 
-c file  >>>>>> character special file 
-d file >>>>>>>> directory 
-f file >>>>>>>> directory or special file 
-g file >>>>>>>> group id
-k file >>>>>>>>
-p file >>>>>>>> pipe 
-t 
-u
-r
-w 
-x
-s
-e 


shell lops 
for and while 
for [conditions ]
do
    body

done

while [condition]
do  
    body
done

for i in (your box )
do 
    body 
done 

condition in bash

if [conditions] 
then 
    body 
fi 

if [condition]; then
    body 
else
    body
fi

if [condition]; then
    body
elif[condition2] ;then
    body
else
    body
fi

function in bash

function name (){
    body
}