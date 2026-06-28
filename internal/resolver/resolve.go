package resolver

import (
	"fmt"
	"net"
)

// Resolve accepts a hostname, IP address, or CIDR range and returns the
// expanded list of target IP strings.
func Resolve(target string) ([]string, error) {
	// CIDR block
	if _, network, err := net.ParseCIDR(target); err == nil {
		hosts := expandCIDR(network)
		if len(hosts) == 0 {
			return nil, fmt.Errorf("CIDR %q contains no usable hosts", target)
		}
		return hosts, nil
	}

	// Plain IP
	if ip := net.ParseIP(target); ip != nil {
		return []string{ip.String()}, nil
	}

	// Hostname
	addrs, err := net.LookupHost(target)
	if err != nil {
		return nil, fmt.Errorf("could not resolve %q: %w", target, err)
	}
	return addrs, nil
}

func expandCIDR(network *net.IPNet) []string {
	var ips []string
	for ip := cloneIP(network.IP); network.Contains(ip); incrementIP(ip) {
		ips = append(ips, ip.String())
	}
	// Drop network address and broadcast address
	if len(ips) > 2 {
		ips = ips[1 : len(ips)-1]
	}
	return ips
}

func cloneIP(ip net.IP) net.IP {
	clone := make(net.IP, len(ip))
	copy(clone, ip)
	return clone
}

func incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] != 0 {
			break
		}
	}
}
