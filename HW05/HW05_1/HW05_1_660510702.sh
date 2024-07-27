#!/bin/bash

# ทิชานนท์ รตนแสนวัน
# 660510702
# HW05_1
# 204203 Sec 001

transform_name() {
    arr=($1)
    first_name="${arr[0]^}"
    last_name="${arr[1]^}"
    fl="${#first_name}"
    ll="${#last_name}"

    fll=$(( 5 - fl ))
    lll=$(( 9 - ll ))

    if [ $fll -lt 0 ]; then
        fll=0
    fi

    if [ $lll -lt 0 ]; then
        lll=0
    fi

    if [ $fl -gt 5 ]; then
        first_name=${first_name:0:(( 5 + lll ))}
    fi

    if [ $ll -gt 9 ]; then
        last_name=${last_name:0:(( 9 + fll ))}
    fi

    echo "$last_name.$first_name"
}

# Main script
main() {
    # echo $1
    result=$(transform_name "$1")
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