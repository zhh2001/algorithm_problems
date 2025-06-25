#!/bin/bash

cat nowcoder.txt | awk -F '\n' '(NR==5){print}'
exit 0
