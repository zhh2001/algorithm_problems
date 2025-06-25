#!/bin/bash

for i in {0..500}
    do
        if [[ $(( $i % 7)) -eq 0 ]]
        then
            echo $i
        fi
    done
exit 0
