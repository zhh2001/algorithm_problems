#!/bin/bash

numy=$(wc -l <nowcoder.txt)
numx=$(awk 'NR==1{print NF}' nowcoder.txt)

for ((i = 1; i <= numx; i++)); do
	awk -v col=$i '{print $col}' nowcoder.txt | tr '\n' ' '
	printf "\n"
done
exit 0
