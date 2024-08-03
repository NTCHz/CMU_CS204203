#!/usr/bin/bash
# ทิชานนท์ รตนแสนวัน
# 660510702
# HW05_3
# 204203 Sec 001

declare -A Ree
count=0 
while read -r line; do
    if [ $count -eq 0 ]; then
        count=$((count+1))
        continue
    fi
    line="${line%$'\n'}"
    line="${line//,/}"
    arr=($line)
    for (( i=0; i<${#arr[@]}; i++ )); do
        Ree["${arr[$i]}"]+="(v$count $(($i+1))) "
    done
    count=$((count+1))
done 

for key in "${!Ree[@]}"; do
    echo "$key: ${Ree[$key]} "
done | sort
