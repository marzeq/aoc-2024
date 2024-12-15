#!/bin/bash
lastday=$(ls | grep -P "day\d+" | sed 's/^day//' | sort -n | tail -n 1)
newday=$(( lastday + 1 ))

if [ "${newday}" = "" ]; then
  exit 1
fi

dayname="day${newday}"

cp -r daytemplate $dayname && \
sed -i "s/dayx/${dayname}/" $dayname/main.go && \

sed -i "/\/\/ <dayimport>/i \"github.com/marzeq/aoc-2024/${dayname}\"" main.go && \
sed -i "/\/\/ <daycase>/i case ${newday}: printRes(${dayname}.Run(part, lines), tstart)" main.go && \
gofumpt -w main.go && \
mkdir -p inputs && vim "${dayname}/main.go" "inputs/${newday}.txt"
