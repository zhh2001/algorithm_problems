#!/bin/bash

awk '{for (i=6;i<=NF;i++) printf $i" ";print ""}' nowcoder.txt |
    awk -F":|," '{switch(NR){
        case 1:
            print "serverVersion:"$2;
            break
        case 3:
            print "serverName:"$2;
            break
        case 4:
            print "osName:"$2;
            print "osVersion:"$4;
            break
        default:
            break;
    }}'
exit 0
