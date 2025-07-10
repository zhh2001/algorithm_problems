#!/bin/bash

grep 'ESTABLISHED' nowcoder.txt |
	awk '{ print $5 }' |
	awk -F ":" '
    {
        if ($2 == "3306")
            arr[$1]++
    }
    END{
        for (i in arr)
            print arr[i], i
    }
' |
    sort -k1nr
exit 0
