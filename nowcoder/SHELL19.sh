#!/bin/bash

read num
for ((i = 1; i <= $num; i++)); do
    for ((j = $num; j > $i; j--)); do
        echo -n " "
    done
    for ((j = 0; j < $i; j++)); do
        echo -n "* "
    done
    printf '\n'
done
exit 0
