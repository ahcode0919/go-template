#!/bin/bash

coverage_file=$1
threshold=$2

if [[ "$#" -lt 2 ]]; then
    echo "Error... Usage: coverage_check.sh {path_to_coverage_file} {threshold}"
    exit 1
fi

if [ ! -f $coverage_file ]; then
    echo "Error... Coverage file not found!"
    exit 1
fi

if ! [[ $threshold =~ ^[0-9]+([.][0-9]+)?$ ]]; then
    echo "Error... Threshold must be a valid decimal number!"
    exit 1
fi

coverage_report=$(go tool cover -func=$coverage_file)
total_line=$(tail -1 <<< "$coverage_report")
percent_covered=$(awk '{print $3}' <<< "$total_line")
number=$(tr -d '%' <<< "$percent_covered")

echo "Total coverage: ${percent_covered}"
echo "Threshold: ${threshold}%"

if (( $(echo "$number < $threshold" | bc -l) )); then
    echo -e "\033[31mCoverage below threshold\033[0m"
    exit 1
else
    echo -e "\033[32mCoverage passed!\033[0m"
    exit 0
fi