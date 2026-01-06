#!/usr/bin/env bash

set -euo pipefail

echo "Fetching latest release metadata..."
release_json=$(curl -fsSL "https://api.github.com/repos/0adri3n/3g-scan/releases/latest")

# Extract the browser_download_url for the asset named exactly "3g-scan"
url=$(printf '%s' "$release_json" | perl -0777 -ne 'print $1 if /"name"\s*:\s*"3g-scan".*?"browser_download_url"\s*:\s*"([^\"]+)"/s')

if [ -z "$url" ]; then
	echo "No asset named '3g-scan' found in latest release." >&2
	exit 1
fi

echo "Downloading $url -> 3g-scan"
curl -L "$url" -o 3g-scan
chmod +x 3g-scan
sudo mv 3g-scan /usr/local/bin/3g-scan

echo -e "\n🚀 Welcome to 3g-scan (GoGoGadgeto-Scan)\n\nUsage:\n  sudo 3g-scan -ranges <CIDR[,CIDR,...]> [options]\n\nExamples:\n  sudo 3g-scan -ranges 192.168.1.0/24\n  sudo 3g-scan -ranges 192.168.1.0/24 -csv scan.csv -yaml scan.yaml\n\nOptions:\n  -p_scan    Enable port scanning (default: true)\n  -routine   Enable concurrent scanning (default: true)\n  -debug     Enable debug logs\n\nDocs:\n  https://github.com/0adri3n/3g-scan\n"