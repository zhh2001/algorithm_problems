#!/bin/bash

while read line; do
    if [[ ! $line =~ [B,b] ]]; then
        echo $line
    fi

done <nowcoder.txt
exit 0
