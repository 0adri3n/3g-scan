curl -fsSL https://api.github.com/repos/0adri3n/3g-scan/releases/latest \
| grep browser_download_url \
| grep linux \
| cut -d '"' -f 4 \
| wget -qi - \
&& chmod +x 3g-scan \
&& sudo mv 3g-scan /usr/local/bin/3g-scan \
&& echo -e "\n🚀 Welcome to 3g-scan (GoGoGadgeto-Scan)\n\nUsage:\n  sudo 3g-scan -ranges <CIDR[,CIDR,...]> [options]\n\nExamples:\n  sudo 3g-scan -ranges 192.168.1.0/24\n  sudo 3g-scan -ranges 192.168.1.0/24 -csv scan.csv -yaml scan.yaml\n\nOptions:\n  -p_scan    Enable port scanning (default: true)\n  -routine   Enable concurrent scanning (default: true)\n  -debug     Enable debug logs\n\nDocs:\n  https://github.com/0adri3n/3g-scan\n"