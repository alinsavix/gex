#!/bin/sh
for i in $(seq -f "%03g" 0 116)
do
    ../gex -v -o maze$i.png maze$i > maze$i.txt
done
