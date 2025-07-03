#!/bin/bash

awk -F'/' '{
    a[$3]++
}END{
    for(i in a){
        printf "%s %s\n",a[i],i
    }
}' | sort -rn
exit 0
