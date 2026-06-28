# go-scan

A fast, concurrent network and port scanner written in Go. Built for asset discovery, port enumeration, and mapping what's on your network.

## Features

- **Concurrent TCP scan** — goroutine worker pool scans thousands of ports in seconds
- **Flexible port ranges** — single ports (`80`), ranges (`1-1024`), comma-separated (`22,80,443`), or all (`all`)
- **CIDR support** — scan entire subnets (`192.168.1.0/24`) with a single command
- **Hostname resolution** — accepts hostnames, IPv4/IPv6 addresses, and CIDR blocks
- **Service identification** — maps open ports to well-known service names (ssh, http, mysql, rdp, etc.)
- **Tunable performance** — configurable thread count and connection timeout

## Coming Soon

- [ ] **Host discovery / ping sweep** — find live hosts on a subnet before port scanning
- [ ] **UDP scanning** — enumerate UDP services (DNS, SNMP, DHCP, TFTP, etc.)
- [ ] **Banner grabbing** — pull raw service banners to identify software and versions
- [ ] **OS fingerprinting** — infer operating system from TCP/IP stack characteristics
- [ ] **SYN scan** — stealthy half-open scan without completing the TCP handshake (requires raw socket / root)
- [ ] **SSL/TLS inspection** — extract certificate details, cipher suites, and expiry from TLS services
- [ ] **HTTP probing** — detect web servers, grab page titles, response headers, and redirect chains
- [ ] **ARP scan** — layer-2 host discovery for local network segments
- [ ] **Output formats** — export results as JSON, CSV, or XML
- [ ] **WHOIS integration** — lookup registration and ASN info for external targets
- [ ] **Traceroute** — map the network path to a target host
- [ ] **Saved scan profiles** — store and replay common scan configurations
- [ ] **CVE correlation** — cross-reference detected service versions with known vulnerabilities

## Installation

```bash
git clone https://github.com/yourusername/go-scan.git
cd go-scan
go build -o goscan .
```

## Usage

```text
goscan scan [target] [flags]
```

### Examples

```bash
# Scan common ports (1–1024) on a host
goscan scan 192.168.1.1

# Scan specific ports on a hostname
goscan scan example.com -p 22,80,443,8080

# Scan a port range
goscan scan 10.0.0.1 -p 1-65535

# Scan all 65535 ports
goscan scan 10.0.0.1 -p all

# Scan every host in a subnet
goscan scan 192.168.1.0/24

# Tune for speed (more threads, shorter timeout)
goscan scan 192.168.1.1 -p 1-1024 -n 1000 -t 200
```

### Flags

| Flag | Short | Default | Description |
| ---- | ----- | ------- | ----------- |
| `--ports` | `-p` | `1-1024` | Port specification (`80`, `1-1024`, `22,80,443`, `all`) |
| `--timeout` | `-t` | `500` | Per-connection timeout in milliseconds |
| `--threads` | `-n` | `500` | Concurrent goroutine count |

## Requirements

- Go 1.22+

## License

MIT
