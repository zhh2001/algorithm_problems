#!/bin/bash

filename="nowcoder.txt"

while read line; do
	echo $line | awk '$0!~/this/ {print $line}'
done <$filename
exit 0
