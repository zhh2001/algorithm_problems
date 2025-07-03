#!/bin/bash

grep '192.168.1.22' nowcoder.txt |
	awk '
    {
        arr[$7]++
    }
    END{
        for (i in arr)
            print arr[i], i
    }
' |
    sort -k1,1r
exit 0
