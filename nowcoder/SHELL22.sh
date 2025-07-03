#!/bin/bash

declare -A dict

while read line; do
    key=$(echo "${line}" | awk -F':' '{print $1}')
    value=$(echo "${line}" | awk -F':' '{print $2}')
    dict[${key}]="${dict[${key}]} ${value}"
done <nowcoder.txt

for key in $(echo "${!dict[@]}" | tr ' ' '\n' | sort -n); do
    echo "[${key}]"
    for value in ${dict[${key}]}; do
        echo "${value}"
    done
done
exit 0
