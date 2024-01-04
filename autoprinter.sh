#!/bin/bash

if [ $# -eq 0 ]; then
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Please provide the filename to print." >> print_log.txt
    exit 1
fi

image_file="$1.jpeg"

if [ ! -f "$image_file" ]; then
    echo "$(date '+%Y-%m-%d %H:%M:%S'): The image file '$image_file' was not found." >> print_log.txt
    exit 1
fi

printer=$(lpstat -d | awk -F ": " '{print $2}')

{
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Starting printing of file '$image_file' on printer '$printer'"
    lp -d "$printer" "$image_file"
    echo "$(date '+%Y-%m-%d %H:%M:%S'): Printing of file '$image_file' on printer '$printer' completed"
} >> print_log.txt

echo "Image '$image_file' sent for printing to $printer"
echo "Logs saved in print_log.txt"

#.\autoprinter.sh 2468
