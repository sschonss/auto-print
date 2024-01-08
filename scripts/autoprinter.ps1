param(
    [Parameter(Mandatory=$true)]
    [string]$image_name,

    [Parameter(Mandatory=$true)]
    [int]$num_copies,

    [Parameter(Mandatory=$true)]
    [string]$printer_name
)

$image_file = ".\files\$image_name.jpeg"

if (-not (Test-Path $image_file)) {
    Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): The image file '$image_file' was not found."
    exit 1
}

Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Starting printing of $num_copies copies of file '$image_file' on printer '$printer_name'"

for ($i = 0; $i -lt $num_copies; $i++) {
    Start-Process -FilePath "C:\Windows\System32\mspaint.exe" -ArgumentList "/pt `"$image_file`" `"$printer_name`""
}

Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Printing of $num_copies copies of file '$image_file' on printer '$printer_name' completed"
"Image '$image_file' sent for printing ($num_copies copies) to $printer_name" | Out-File -Append -FilePath print_log.txt
"Logs saved in print_log.txt" | Out-File -Append -FilePath print_log.txt
