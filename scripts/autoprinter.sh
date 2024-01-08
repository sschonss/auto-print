#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <image_name> <num_copies>"
    exit 1
fi

image_file="files/$1.jpeg"
num_copies="$2"
temp_image="./files/temp_$1.jpeg"

if [ ! -f "$image_file" ]; then
    echo "$(date '+%Y-%m-%d %H:%M:%S'): The image file '$image_file' was not found."
    exit 1
fi

convert "$image_file" -rotate 90 -resize 'x792' "$temp_image"

impressora=$(lpstat -d | awk -F ": " '{print $2}')

{
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Starting printing of $num_copies copies of file '$temp_image' (resized) on printer '$impressora'"
    lp -o orientation-requested=3 -n "$num_copies" -o position=top-bottom -d "$impressora" "$temp_image"
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Printing of $num_copies copies of file '$temp_image' (resized) on printer '$impressora' completed"
} >> print_log.txt

rm "$temp_image"

echo "Resized image '$temp_image' sent for printing ($num_copies copies) to $impressora"
echo "Logs saved in print_log.txt"
