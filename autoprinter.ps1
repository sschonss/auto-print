param(
    [Parameter(Mandatory=$true)]
    [string]$fileName
)

$extension = (Get-Item $fileName).Extension
if ($extension -ne '.jpeg' -and $extension -ne '.jpg') {
    $fileName += '.jpeg'
}

if (-not (Test-Path $fileName)) {
    "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): The image file '$fileName' was not found." | Out-File -Append -FilePath .\print_log.txt
    exit 1
}

$printer = Get-Printer | Where-Object { $_.Default } | Select-Object -ExpandProperty Name

"$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Starting printing of file '$fileName' on printer '$printer'" | Out-File -Append -FilePath .\print_log.txt
Start-Process -FilePath 'mspaint' -ArgumentList $fileName -Wait
"$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Printing of file '$fileName' on printer '$printer' completed" | Out-File -Append -FilePath .\print_log.txt

Write-Host "Image '$fileName' sent for printing to $printer"
Write-Host "Logs saved in print_log.txt"

# .\autoprinter.ps1 "2468"
# .\autoprinter.ps1 "2468.jpeg"