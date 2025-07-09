#!/bin/bash

awk '
    {
        if ($1 == "tcp")
            arr[$6]++
    }
    END {
        for (i in arr)
            print i, arr[i]
    }
' nowcoder.txt |
    sort -k2nr
exit 0
