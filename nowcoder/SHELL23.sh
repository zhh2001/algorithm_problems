#!/bin/bash

awk '/^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}/{
    if($0~/23\/Apr\/2020/){
        a[$1]++
    }
}END{
    for(i in a){
        printf "%s %s\n",a[i],i
    }
}' | sort -rn
exit 0
