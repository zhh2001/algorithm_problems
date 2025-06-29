#!/bin/bash

function main() {
    declare -A dict
    local dict
    while read line; do
        for i in ${line[*]}; do
            flag=0
            for key in ${!dict[*]}; do
                if [ $i = $key ]; then
                    let dict[$key]++,flag=1
                    break
                fi
            done
            if ((flag == 0)); then
                dict[$i]=1
            fi
        done
    done <$1

    local arr_key=(${!dict[*]})
    local arr_val=(${dict[*]})
    len=${#arr_key[*]}
    for ((i = 0; i < len - 1; i++)); do
        for ((j = i + 1; j < len; j++)); do
            if ((${arr_val[$i]} > ${arr_val[$j]})); then
                temp=${arr_val[$i]}
                temp2=${arr_key[$i]}
                arr_val[$i]=${arr_val[$j]}
                arr_key[$i]=${arr_key[$j]}
                arr_val[$j]=$temp
                arr_key[$j]=$temp2
            fi
        done
    done
    for ((i = 0; i < len; i++)); do
        echo "${arr_key[$i]} ${arr_val[$i]}"
    done
    exit 0
}

main nowcoder.txt
