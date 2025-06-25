#!/bin/bash

for i in $(cat nowcoder.txt);
    do
        if [ ${#i} -lt 8 ]; then
            echo ${i}
        fi
    done
exit 0