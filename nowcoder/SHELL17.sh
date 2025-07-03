#!/bin/bash

files=nowcoder.txt
while read line; do
    echo $line | awk -F: '{for(i=NF;i>=1;i--)printf "%s%s",$i,(i==1 ? ORS : OFS=":")}'
done <$files
exit 0
