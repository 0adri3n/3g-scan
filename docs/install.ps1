$repo = "0adri3n/3g-scan"
$release = Invoke-RestMethod "https://api.github.com/repos/$repo/releases/latest"
$asset = $release.assets | Where-Object { $_.name -match "3g-scan.exe" } | Select-Object -First 1

$dest = "$env:ProgramFiles\3g-scan"
New-Item -ItemType Directory -Force -Path $dest | Out-Null

Invoke-WebRequest $asset.browser_download_url -OutFile "$dest\3g-scan.exe"

$envPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
if ($envPath -notlike "*$dest*") {
    [Environment]::SetEnvironmentVariable(
        "Path",
        "$envPath;$dest",
        "Machine"
    )
}

Write-Host ""
Write-Host "🚀 Welcome to 3g-scan (GoGoGadgeto-Scan)"
Write-Host ""
Write-Host "Usage:"
Write-Host "  3g-scan -ranges <CIDR[,CIDR,...]> [options]"
Write-Host ""
Write-Host "Examples:"
Write-Host "  3g-scan -ranges 192.168.1.0/24"
Write-Host "  3g-scan -ranges 192.168.1.0/24 -csv scan.csv -yaml scan.yaml"
Write-Host ""
Write-Host "Options:"
Write-Host "  -p_scan    Enable port scanning (default: true)"
Write-Host "  -routine   Enable concurrent scanning (default: true)"
Write-Host "  -debug     Enable debug logs"
Write-Host ""
Write-Host "Docs:"
Write-Host "  https://github.com/0adri3n/3g-scan"
Write-Host ""
