#!/bin/bash

if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <image_name> <num_copies> <printer_name>"
    exit 1
fi

image_file="files/$1.jpeg"
num_copies="$2"
printer_name="$3"
temp_image="./files/temp_$1.jpeg"

if [ ! -f "$image_file" ]; then
    echo "$(date '+%Y-%m-%d %H:%M:%S'): The image file '$image_file' was not found."
    exit 1
fi

convert "$image_file" -rotate 90 -resize 'x792' "$temp_image"

{
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Starting printing of $num_copies copies of file '$temp_image' (resized) on printer '$printer_name'"
    lp -o orientation-requested=3 -n "$num_copies" -o position=top-bottom -d "$printer_name" "$temp_image"
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Printing of $num_copies copies of file '$temp_image' (resized) on printer '$printer_name' completed"
} >> print_log.txt

rm "$temp_image"

echo "Resized image '$temp_image' sent for printing ($num_copies copies) to $printer_name"
echo "Logs saved in print_log.txt"
