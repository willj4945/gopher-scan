package scanner

import (
	"fmt"
	"strconv"
	"strings"
)

// ParsePorts converts a port specification string into a slice of port numbers.
// Accepted formats: "80", "22,80,443", "1-1024", "all" / "-"
func ParsePorts(spec string) ([]int, error) {
	spec = strings.TrimSpace(spec)
	if spec == "all" || spec == "-" {
		ports := make([]int, 65535)
		for i := range ports {
			ports[i] = i + 1
		}
		return ports, nil
	}

	seen := make(map[int]bool)
	var ports []int

	for _, part := range strings.Split(spec, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if strings.Contains(part, "-") {
			bounds := strings.SplitN(part, "-", 2)
			lo, err1 := strconv.Atoi(strings.TrimSpace(bounds[0]))
			hi, err2 := strconv.Atoi(strings.TrimSpace(bounds[1]))
			if err1 != nil || err2 != nil || lo < 1 || hi > 65535 || lo > hi {
				return nil, fmt.Errorf("invalid range %q (must be 1–65535)", part)
			}
			for p := lo; p <= hi; p++ {
				if !seen[p] {
					ports = append(ports, p)
					seen[p] = true
				}
			}
		} else {
			p, err := strconv.Atoi(part)
			if err != nil || p < 1 || p > 65535 {
				return nil, fmt.Errorf("invalid port %q (must be 1–65535)", part)
			}
			if !seen[p] {
				ports = append(ports, p)
				seen[p] = true
			}
		}
	}

	if len(ports) == 0 {
		return nil, fmt.Errorf("no valid ports in %q", spec)
	}
	return ports, nil
}
