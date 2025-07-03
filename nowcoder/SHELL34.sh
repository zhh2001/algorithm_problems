#!/bin/bash

grep -v 'USER' nowcoder.txt |
	awk '
    BEGIN{
        vsz_sum = 0
        rss_sum = 0
    }
    {
        vsz_sum += $5
        rss_sum += $6
    }
    END{
        print "MEM TOTAL"
        printf("VSZ_SUM:%0.1fM", vsz_sum/1024)
        printf ","
        printf("RSS_SUM:%0.3fM", rss_sum/1024)
    }
'
exit 0
