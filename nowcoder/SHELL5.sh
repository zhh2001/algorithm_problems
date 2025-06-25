#!/bin/bash

num=0
while read line
    do
        ((num++))
        [ "$line" == "" ] && echo $num
    done < nowcoder.txt
exit 0