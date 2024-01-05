param(
    [Parameter(Mandatory=$true)]
    [string]$image_name,

    [Parameter(Mandatory=$true)]
    [int]$num_copies
)

$image_file = "files\$image_name.jpeg"
$temp_image = "files\temp_$image_name.jpeg"

if (-not (Test-Path $image_file)) {
    Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): The image file '$image_file' was not found."
    exit 1
}

& convert $image_file -rotate 90 -resize 'x792' $temp_image

$printer = Get-Printer | Where-Object { $_.Default } | Select-Object -ExpandProperty Name

Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Starting printing of $num_copies copies of file '$temp_image' (resized) on printer '$printer'"

for ($i = 0; $i -lt $num_copies; $i++) {
    Start-Process -FilePath 'mspaint' -ArgumentList $temp_image -Wait
}

Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Printing of $num_copies copies of file '$temp_image' (resized) on printer '$printer' completed"
"Resized image '$temp_image' sent for printing ($num_copies copies) to $printer" | Out-File -Append -FilePath print_log.txt
"Logs saved in print_log.txt" | Out-File -Append -FilePath print_log.txt

Remove-Item $temp_image 
