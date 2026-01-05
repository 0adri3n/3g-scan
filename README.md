# 3g-scan (GoGoGadgeto-Scan)

3g-scan (short for **GoGoGadgeto-Scan**) is a lightweight **network scanner written in Go**.  
It scans one or multiple IP ranges, checks host availability using ICMP, and retrieves **MAC addresses** on the local network with OS-specific implementations.

The tool is designed to be **simple, fast, and cross-platform**, with native support for **Windows, Linux, and macOS**.

---

## ✨ Features

- Scan one or multiple **CIDR IP ranges**
- **Ping** hosts to check availability
- Retrieve **MAC addresses** for reachable hosts
- OS-specific implementations:
  - **Windows**: ARP table lookup
  - **Linux / macOS**: Native ARP requests
- Optional **debug logging**
- Lightweight and dependency-minimal
- Written in pure **Golang**

---

## 🧠 How It Works

1. Parse IP ranges from CLI arguments
2. Expand CIDR ranges into individual IPs
3. Ping each IP to check if the host is alive
4. If reachable:
   - Retrieve the MAC address using the appropriate OS method
5. Display scan progress and results

---

## 📦 Project Structure

```text
3g-scan/
├── main.go
├── ggg_network/
│   ├── cidr.go        # CIDR expansion
│   ├── ping.go        # ICMP ping
│   ├── mac_linux.go  # Linux/macOS ARP resolution
│   ├── mac_windows.go# Windows ARP table parsing
│   └── ...
└── README.md
````

---

## 🚀 Installation

### Requirements

* Go **1.20+**
* Administrator / root privileges (required for ICMP & ARP)

### Clone the repository

```bash
git clone https://github.com/0adri3n/3g-scan.git
cd 3g-scan
```

### Build

```bash
go build -o 3g-scan
```

---

## ▶️ Usage

```bash
3g-scan -ranges <CIDR[,CIDR,...]> -iface <interface> [options]
```

### Arguments

| Flag      | Description                              |
| --------- | ---------------------------------------- |
| `-ranges` | CIDR IP ranges to scan (comma separated) |
| `-iface`  | Network interface to use (Linux/macOS)   |
| `-debug`  | Enable debug logs (`true` / `false`)     |

---

### Examples

#### Linux / macOS

```bash
sudo ./3g-scan \
  -ranges 192.168.1.0/24,10.0.0.0/24 \
  -iface eth0 \
  -debug true
```

#### Windows (Administrator)

```powershell
.\3g-scan.exe -ranges 192.168.1.0/24 -debug true
```

> ⚠️ On Windows, MAC addresses are resolved using the system ARP table.
> Hosts must be reachable (pinged) to appear in the ARP cache.

---

## 🖥️ Platform Support

| OS      | Ping | MAC Address |
| ------- | ---- | ----------- |
| Windows | ✅    | ARP table   |
| Linux   | ✅    | Native ARP  |
| macOS   | ✅    | Native ARP  |

---

## 🔐 Permissions

This tool **must be run with elevated privileges**:

* Linux / macOS: `sudo`
* Windows: Run terminal as **Administrator**

Without privileges:

* ICMP ping may fail
* ARP resolution may not work

---

## 🧪 Debug Mode

Enable debug logging with:

```bash
-debug true
```

Logs are printed to `stderr` and include:

* Scan progress
* Host availability
* Internal scan details

---

## ⚠️ Limitations

* MAC addresses can only be retrieved on the **local network**
* ARP does not work across VLANs or routed networks
* Windows does not support raw ARP packets natively
* Firewalls may block ICMP

---

## 🛣️ Roadmap

* [X] Vendor lookup (OUI)
* [ ] JSON / CSV output
* [X] Parallel scanning (routines)
* [X] Hostname resolution
* [X] Service checking
* [ ] Export scan results

---

## 📜 License

MIT License

---

## 👤 Author

**0adri3n**
GitHub: [https://github.com/0adri3n](https://github.com/0adri3n)

---

## ⭐ Contributing

Pull requests are welcome!
If you find a bug or have an idea, feel free to open an issue.