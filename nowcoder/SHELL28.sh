#!/bin/bash

awk '
    {
        tmp = substr($4, 14, 5)
        arr[tmp]++
    }
    END{
        for (i in arr)
            print arr[i], i
    }
' nowcoder.txt | sort -k1,1r -k2.2r
exit 0
