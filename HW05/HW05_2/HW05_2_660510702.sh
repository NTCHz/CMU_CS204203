#!/usr/bin/bash
# ทิชานนท์ รตนแสนวัน
# 660510702
# HW05_2
# 204203 Sec 001

local_maxima() {
   echo "$1"
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