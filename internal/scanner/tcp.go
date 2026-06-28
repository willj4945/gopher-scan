package scanner

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

type Result struct {
	Port    int
	Service string
}

// ScanTCP performs a concurrent TCP connect scan against host on the given
// ports. threads controls the goroutine pool size; timeoutMs is per-connection.
func ScanTCP(host string, ports []int, timeoutMs, threads int) []Result {
	timeout := time.Duration(timeoutMs) * time.Millisecond
	sem := make(chan struct{}, threads)
	resCh := make(chan Result, len(ports))

	var wg sync.WaitGroup
	for _, port := range ports {
		wg.Add(1)
		sem <- struct{}{}
		go func(p int) {
			defer wg.Done()
			defer func() { <-sem }()

			addr := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", addr, timeout)
			if err == nil {
				conn.Close()
				resCh <- Result{Port: p, Service: serviceName(p)}
			}
		}(port)
	}

	wg.Wait()
	close(resCh)

	var results []Result
	for r := range resCh {
		results = append(results, r)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Port < results[j].Port
	})
	return results
}
