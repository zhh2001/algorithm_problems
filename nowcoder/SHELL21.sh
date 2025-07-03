#!/bin/bash

while read line; do
    s=""
    for ((i = ${#line} - 1; i >= 0; i--)); do
        s="${line:i:1}${s}"
        if [[ $(((${#line} - i) % 3)) -eq 0 && ${i} -ne 0 ]]; then
            s=",${s}"
        fi

    done

    echo $s
done
exit 0
