#!/bin/bash

while read line
    do
        if [ $line != "" ];then
            echo $line
        fi
    done <  nowcoder.txt
exit 0