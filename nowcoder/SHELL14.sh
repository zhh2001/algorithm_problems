#!/bin/bash

len=''
i=0
sum=0
while read num; do
    if [ -z "${len}" ]; then
        len=$num
        continue
    fi
    ((sum += num))

    ((i++))
    if [ ${i} -eq $len ]; then
        break
    fi
done
printf "%.3f" $(echo "scale=3; ${sum} / ${len}" | bc)
exit 0
