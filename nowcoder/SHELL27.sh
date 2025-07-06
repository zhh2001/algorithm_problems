#!/bin/bash

awk '
    BEGIN{sum = 0}
    {
        if ($9 == 404 && $15 ~ /baidu/)
            sum++
    }
    END{
        print sum
    }
' nowcoder.txt
exit 0
