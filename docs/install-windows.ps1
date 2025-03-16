# install.ps1 - Installation script for Gitoko on Windows

$ErrorActionPreference = "Stop"

Write-Host "Installing Gitoko - Git Cherry-Picking Tool" -ForegroundColor Blue

# Set download URL based on latest release
$repo = "usman1100/gitoko"
$apiUrl = "https://api.github.com/repos/$repo/releases/latest"
$binaryName = "gitoko-windows-amd64.exe"

# Get the latest release info
try {
    $releaseInfo = Invoke-RestMethod -Uri $apiUrl -Method Get
}
catch {
    Write-Host "Error: Failed to fetch release information. $_" -ForegroundColor Red
    exit 1
}

# Extract download URL
$downloadUrl = $releaseInfo.assets | Where-Object { $_.name -eq $binaryName } | Select-Object -ExpandProperty browser_download_url

if (-not $downloadUrl) {
    Write-Host "Error: Could not find download URL for $binaryName" -ForegroundColor Red
    exit 1
}

# Create installation directory if it doesn't exist
$installDir = "$env:USERPROFILE\bin"
if (-not (Test-Path $installDir)) {
    New-Item -ItemType Directory -Path $installDir | Out-Null
    Write-Host "Created directory: $installDir" -ForegroundColor Green
}

# Download the binary
Write-Host "Downloading Gitoko..." -ForegroundColor Blue
try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile "$installDir\gitoko.exe"
}
catch {
    Write-Host "Error: Failed to download Gitoko. $_" -ForegroundColor Red
    exit 1
}

# Add to PATH if not already there
$currentPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($currentPath -notlike "*$installDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$currentPath;$installDir", "User")
    Write-Host "Added $installDir to your PATH" -ForegroundColor Green
    Write-Host "You may need to restart your terminal for changes to take effect" -ForegroundColor Yellow
}
else {
    Write-Host "$installDir is already in your PATH" -ForegroundColor Green
}

Write-Host "Gitoko has been successfully installed!" -ForegroundColor Green
Write-Host "You can now use it by running 'gitoko' in any git repository." -ForegroundColor Blue