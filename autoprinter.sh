#!/bin/bash

# Usage: ./autoprinter.sh <image_name> <num_copies>

if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <image_name> <num_copies>"
    exit 1
fi

image_file="files/$1.jpeg"
num_copies="$2"

if [ ! -f "$image_file" ]; then
    echo "$(date '+%Y-%m-%d %H:%M:%S'): The image file '$image_file' was not found."
    exit 1
fi

impressora=$(lpstat -d | awk -F ": " '{print $2}')

{
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Starting printing of $num_copies copies of file '$image_file' on printer '$impressora'"
    lp -n "$num_copies" -o position=top-bottom -d "$impressora" "$image_file"
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Printing of $num_copies copies of file '$image_file' on printer '$impressora' completed"
} >> print_log.txt

echo "Image '$image_file' sent for printing ($num_copies copies) to $impressora"
echo "Logs saved in print_log.txt"
