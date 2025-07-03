#!/bin/bash

while read line; do
    m=0

    for ((i = 0; i < ${#line}; i++)); do
        if [[ ${line:i:1} =~ [0-9] ]]; then
            m=$(($m + 1))
        fi
    done
    if [[ $m -eq 1 ]]; then
        echo $line
    fi
done <nowcoder.txt
exit 0
