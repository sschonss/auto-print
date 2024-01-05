param(
    [Parameter(Mandatory=$true)]
    [string]$image_name,

    [Parameter(Mandatory=$true)]
    [int]$num_copies
)

$image_file = "files\$image_name.jpeg"

if (-not (Test-Path $image_file)) {
    Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): The image file '$image_file' was not found."
    exit 1
}

$printer = Get-WmiObject -Query "SELECT * FROM Win32_Printer WHERE Default=$true"
$printerName = $printer.Name

Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Starting printing of $num_copies copies of file '$image_file' on printer '$printerName'"

for ($i = 0; $i -lt $num_copies; $i++) {
    Start-Process -FilePath "C:\Windows\System32\rundll32.exe" -ArgumentList "C:\Windows\System32\shimgvw.dll,ImageView_PrintTo /pt `"$printerName`" `"$image_file`""
}

Write-Host "$(Get-Date -Format 'yyyy-MM-dd HH:mm:ss'): Printing of $num_copies copies of file '$image_file' on printer '$printerName' completed"
"Image '$image_file' sent for printing ($num_copies copies) to $printerName" | Out-File -Append -FilePath print_log.txt
"Logs saved in print_log.txt" | Out-File -Append -FilePath print_log.txt
