#!/bin/bash

sum=0
while read -r line;
do
    read -r var1 var2 < <(echo $line | grep -E '[0-9]{1,3},[0-9]{1,3}' -o | awk -F, '{ print $1, $2 }') 
    sum=$((var1*var2 + sum))
done < <(grep -E "mul\([0-9]{1,3},[0-9]{1,3}\)" input.txt -o)

echo $sum