package cmd

import (
	"fmt"
	"os"
	"time"

	"go-scan/internal/resolver"
	"go-scan/internal/scanner"

	"github.com/spf13/cobra"
)

var (
	portRange string
	timeoutMs int
	threads   int
)

var scanCmd = &cobra.Command{
	Use:   "scan [target]",
	Short: "Scan a target for open TCP ports",
	Long: `Perform a concurrent TCP connect scan against a host, IP, or CIDR range.

Examples:
  goscan scan 192.168.1.1
  goscan scan example.com -p 22,80,443
  goscan scan 10.0.0.0/24 -p 1-65535 -n 1000`,
	Args: cobra.ExactArgs(1),
	Run:  runScan,
}

func init() {
	scanCmd.Flags().StringVarP(&portRange, "ports", "p", "1-1024", `Port range to scan (e.g. "80", "1-1024", "22,80,443", "all")`)
	scanCmd.Flags().IntVarP(&timeoutMs, "timeout", "t", 500, "Connection timeout in milliseconds")
	scanCmd.Flags().IntVarP(&threads, "threads", "n", 500, "Number of concurrent goroutines")
	rootCmd.AddCommand(scanCmd)
}

func runScan(cmd *cobra.Command, args []string) {
	target := args[0]

	hosts, err := resolver.Resolve(target)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	ports, err := scanner.ParsePorts(portRange)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid port spec: %v\n", err)
		os.Exit(1)
	}

	for _, host := range hosts {
		scanHost(host, ports)
	}
}

func scanHost(host string, ports []int) {
	const line = "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	fmt.Println()
	fmt.Println(line)
	fmt.Printf("  Target  : %s\n", host)
	fmt.Printf("  Ports   : %d\n", len(ports))
	fmt.Printf("  Started : %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(line)

	start := time.Now()
	results := scanner.ScanTCP(host, ports, timeoutMs, threads)
	elapsed := time.Since(start)

	if len(results) == 0 {
		fmt.Println("  No open ports found.")
	} else {
		fmt.Printf("\n  %-8s %-20s %s\n", "PORT", "SERVICE", "STATE")
		fmt.Println("  ─────────────────────────────────────────")
		for _, r := range results {
			fmt.Printf("  %-8d %-20s %s\n", r.Port, r.Service, "open")
		}
		fmt.Println()
	}

	fmt.Printf("  Scanned %d ports in %.2fs  |  %d open\n", len(ports), elapsed.Seconds(), len(results))
	fmt.Println(line)
	fmt.Println()
}
