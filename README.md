# 3g-scan (GoGoGadgeto-Scan)

**3g-scan** (short for **GoGoGadgeto-Scan**) is a fast, lightweight, and efficient **network discovery and scanning tool written in Go**.

It is designed to perform **native network scans only**, relying exclusively on **built-in system capabilities** (ARP, ICMP, TCP, DNS) instead of heavy third-party scanning libraries.  
This approach makes 3g-scan **extremely fast, predictable, and suitable for both enterprise and personal environments**.

---

## 🚀 Key Philosophy

> **No external scanning modules. No wrappers. No overhead.**

3g-scan uses:
- Native **ICMP ping**
- Native **ARP table / neighbor discovery**
- Native **TCP port scanning**
- Native **DNS reverse lookup**

This guarantees:
- ⚡ Maximum performance
- 🧠 Full control over behavior
- 🔐 Minimal attack surface
- 📦 Very small binary size

---

## ✨ Features

- Scan one or multiple **CIDR IP ranges**
- Detect host **availability (Up / Down)**
- Resolve **hostnames (reverse DNS)**
- Retrieve **MAC addresses**
- Identify **hardware vendors**
- Lightweight **TCP port scanning**
- Optional **concurrent scanning (goroutines)**
- Export results to **CSV** and **YAML**
- Configurable via CLI flags
- Cross-platform: **Windows, Linux, macOS**

---

## 🧠 How 3g-scan Works

1. Expand CIDR ranges into individual IP addresses
2. Ping each host using native ICMP
3. If the host is reachable:
   - Resolve hostname(s)
   - Retrieve MAC address
   - Identify vendor from MAC
   - Scan common TCP ports
4. Aggregate results
5. Export to CSV and/or YAML if requested

---

## 🏢 Enterprise & Personal Usage

3g-scan is **equally suitable for**:

### 🏢 Enterprise Networks
- Asset discovery
- Shadow IT detection
- Network audits
- Inventory generation
- Segmentation validation

### 🏠 Personal / Lab Usage
- Home network discovery
- Device identification
- IoT visibility
- Learning networking concepts

Its **non-intrusive and native approach** makes it safe to run in production environments.

---

## 📦 Project Structure

```text
3g-scan/
├── main.go
├── ggg_network/
│   ├── cidrLister.go
│   ├── pinger.go
│   ├── hostnamer.go
│   ├── maccer.go
│   ├── portScanner.go
│   └── resources/
|       └── ieee-oui.txt
└── README.md
````

---

## 🔧 Installation

### Requirements

* Go **1.20+**
* Administrator / root privileges (required for ICMP & ARP)

### Clone

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
3g-scan -ranges <CIDR[,CIDR,...]> [options]
```

### CLI Options

| Flag       | Description                              |
| ---------- | ---------------------------------------- |
| `-ranges`  | IP ranges to scan (comma separated CIDR) |
| `-p_scan`  | Enable or disable port scanning          |
| `-routine` | Enable concurrent scanning (goroutines)  |
| `-debug`   | Enable debug logs                        |
| `-csv`     | Output results to a CSV file             |
| `-yaml`    | Output results to a YAML file            |

---

### Examples

#### Basic scan

```bash
sudo ./3g-scan -ranges 192.168.1.0/24
```

#### Scan with port scanning and exports

```bash
sudo ./3g-scan \
  -ranges 192.168.1.0/24,10.0.0.0/24 \
  -p_scan true \
  -routine true \
  -csv results.csv \
  -yaml results.yaml
```

#### Windows (Administrator)

```powershell
.\3g-scan.exe -ranges 192.168.1.0/24
```

---

## 🖥️ Platform Support

| OS      | Ping | MAC        | Hostname | Ports |
| ------- | ---- | ---------- | -------- | ----- |
| Windows | ✅    | ARP cache  | ✅        | ✅     |
| Linux   | ✅    | Native ARP | ✅        | ✅     |
| macOS   | ✅    | Native ARP | ✅        | ✅     |

---

## 🔐 Permissions

3g-scan **must be run with elevated privileges**:

* Linux / macOS: `sudo`
* Windows: Run as **Administrator**

This is required for:

* ICMP echo requests
* ARP table access

---

## ⚠️ Limitations

* MAC addresses are only available on the **local network**
* ARP does not work across routed networks or VLANs
* Firewalls may block ICMP or TCP probes
* OS detection is **heuristic-based**, not guaranteed

---

## 📊 Benchmark

> *This section will be completed with real-world performance results.*

```text
TBD
```

---

## 🛣️ Roadmap

* [ ] TUI interface
* [ ] JSON export
* [ ] OS fingerprint improvements
* [ ] Custom port lists
* [ ] Vendor database optimization
* [ ] Scan profiling / statistics

---

## 📜 License

MIT License

---

## 👤 Author

**0adri3n**
GitHub: [https://github.com/0adri3n](https://github.com/0adri3n)

---

## ⭐ Contributing

Contributions are welcome!
Feel free to open issues or submit pull requests.