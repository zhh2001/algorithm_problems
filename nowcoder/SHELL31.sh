#!/bin/bash

grep -v 'State' nowcoder.txt |
	awk '
    {
        if ($6 != "" && $6 !~ /[0-9]/)
            print $5
    }
' |
	awk -F ":" '
    {
        arr[$1]++
    }
    END {
        for (i in arr)
            print i, arr[i]
    }
' |
    sort -k2nr -k1r
exit 0
