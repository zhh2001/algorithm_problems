#!/bin/bash

awk '
    {
        arr[$1]++
    }
    END{
        for (i in arr){
            if (arr[i] > 3)
                print arr[i], i
        }
    }
' nowcoder.txt | sort -k1,1r
exit 0
