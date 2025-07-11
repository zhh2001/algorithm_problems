#!/bin/bash

cat nowcoder.txt | awk -v OFS=":" '/3306/{print $5,$6}' | awk -F ":" '{
    a++
    if(arr[$3]==""){
        arr[$3]=1;
    }else{arr[$3]+=1;}
    if(arr1[$1]==""){
        arr1[$1]=1
    }
}END{
    print "TOTAL_IP " length(arr1)
    for(i in arr){
        print i" "arr[i]
    }
    print "TOTAL_LINK " a
}'
exit 0
