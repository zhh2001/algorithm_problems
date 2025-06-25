#!/bin/bash

read -a arr
sum=0
while [ ${#arr[@]} -ge 1 ]
    do
        sum=`expr $sum + 1`
        read -a arr
    done
    echo $sum
exit 0
