#!/bin/bash

multiply=1
sum=0
while read -r line;
do
    # echo $line
    if [ $line == "do()" ] ; then
        multiply=1
    elif [ $line == "don't()" ] ; then
        multiply=0
    else 
        if [ $multiply == "1" ] ; then
            read -r var1 var2 < <(echo $line | grep -E '[0-9]{1,3},[0-9]{1,3}' -o | awk -F, '{ print $1, $2 }') 
            sum=$((var1*var2 + sum))
        fi
    fi
done < <(grep -E "(do|don't)\(\)|(mul\([0-9]{1,3},[0-9]{1,3}\))" input.txt -o)

echo $sum