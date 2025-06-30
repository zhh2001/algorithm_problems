#!/bin/bash

linecount=0
sum=0
count=0

while read line; do
    for ((i = 0; i < ${#line}; i++)); do
        if [[ ${line:$i:1} =~ [1-5] ]]; then
            count=$(($count + 1))
        fi
    done
    linecount=$(($linecount + 1))
    echo "line$linecount number:$count"
    sum=$(($sum + $count))
    count=0

done <nowcoder.txt
echo "sum is $sum"
exit 0
