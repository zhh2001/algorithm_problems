#!/bin/bash

filename="nowcoder.txt"
result=$(cat <$filename | grep '23/Apr/2020:2[0-3]:' | cut -d'-' -f1 | sort | uniq -c | wc -l)
echo "${result}"
exit 0
