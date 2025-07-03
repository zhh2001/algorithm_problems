#!/bin/bash

for i in $(cat nowcoder.txt); do
    if [[ $i =~ [[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3}\.[[:digit:]]{1,3} ]]; then
        a=$(echo $i | awk -F "." '{print $4}')
        b=$(echo $i | awk -F "." '{print $3}')
        c=$(echo $i | awk -F "." '{print $2}')
        d=$(echo $i | awk -F "." '{print $1}')
        if [[ $a > 255 || $b > 255 || $c > 255 || $d > 255 ]]; then
            echo no
        else
            echo yes
        fi
    else
        echo error
    fi
done
exit 0
