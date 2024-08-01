#!/usr/bin/bash
# ทิชานนท์ รตนแสนวัน
# 660510702
# HW05_2
# 204203 Sec 001

local_maxima() {
    arr=($1)
    count=0
    for (( i=1; i<${#arr[@]}-1; i++ )); do
        if [ ${arr[i]} -gt ${arr[i-1]} ] && [ ${arr[i]} -gt ${arr[i+1]} ]; then
            count=$((count+1))
        fi
    done
    echo $count
}

# Main script
main() {
    result=$(local_maxima "$1")
    echo $result
}


# Execute main function if this script is run directly
if [ "$0" = "${BASH_SOURCE[0]}" ]; then
    while read -r line; do
        # echo $line
        line="${line%$'\n'}"
        [[ -n $line ]] && main "$line"
    done < /dev/stdin
fi