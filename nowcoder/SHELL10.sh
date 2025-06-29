#!/bin/bash

filename=nowcoder.txt

if [[ ! -f $filename ]]; then
	echo "File $filename does not exist!"
	exit 1
fi

col2=$(awk '{print $2}' $filename | sort)

counts=$(echo "$col2" | uniq -cd | sort -k1n -k2)

echo "$counts" | awk '{print $1,$2}'
exit 0
