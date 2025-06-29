#!/bin/bash

sum=0
line=1
for i in $(awk '{print $4}' nowcoder.txt)
    do
        if [ $line -gt 1 ]; then
            sum=$(echo $i+$sum|bc)
        fi
        let line++
    done
echo $sum
exit 0