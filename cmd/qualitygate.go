package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zricethezav/gitleaks/v8/report"
)

// QualityGate defines a threshold for a specific severity level
type QualityGate struct {
	Severity string // "critical", "high", "medium", "low", "exception"
	MaxCount int    // Maximum allowed findings
}

// QualityGateResult contains the result of quality gate checks
type QualityGateResult struct {
	Passed   bool
	Failures []string
	Summary  map[string]int // severity: count
}

// ParseQualityGates converts command line specs to QualityGate objects
func ParseQualityGates(gateSpecs []string) ([]QualityGate, error) {
	var gates []QualityGate
	for _, spec := range gateSpecs {
		parts := strings.Split(spec, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid quality gate format: %s (expected severity=count)", spec)
		}

		maxCount, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid count in quality gate %s: %v", spec, err)
		}

		// PRESERVE THE EXACT CASE from command line
		severity := parts[0] // Don't convert to lowercase
		gates = append(gates, QualityGate{
			Severity: severity,
			MaxCount: maxCount,
		})
	}
	return gates, nil
}

// CheckQualityGates validates findings against quality gates
func CheckQualityGates(findings []report.Finding, gates []QualityGate) QualityGateResult {
	result := QualityGateResult{
		Passed:   true,
		Summary:  make(map[string]int),
		Failures: []string{},
	}

	// Count findings by severity (exact case)
	for _, finding := range findings {
		result.Summary[finding.Severity]++
	}

	// Check each gate (EXACT case matching)
	for _, gate := range gates {
		count := result.Summary[gate.Severity] // Exact case match
		if count > gate.MaxCount {
			result.Passed = false
			result.Failures = append(result.Failures,
				fmt.Sprintf("%s findings: %d (max allowed: %d)",
					gate.Severity, count, gate.MaxCount))
		}
	}

	return result
}

// PrintQualityGateSummary prints the results of quality gate checks
func PrintQualityGateSummary(result QualityGateResult) {
	if result.Passed {
		fmt.Println("✅ All quality gates passed")
	} else {
		fmt.Println("❌ Quality gates failed:")
		for _, failure := range result.Failures {
			fmt.Printf("   - %s\n", failure)
		}
	}

	// Print summary table
	fmt.Println("\n📊 Findings Summary:")
	fmt.Println("   Severity   | Count | Status")
	fmt.Println("   -----------|-------|--------")
	
	// Display in consistent order
	severityOrder := []string{"Critical", "High", "Medium", "Low", "Exception"}
	for _, severity := range severityOrder {
		count, exists := result.Summary[severity]
		if exists {
			status := "✅"
			for _, failure := range result.Failures {
				if strings.Contains(failure, severity) {
					status = "❌"
					break
				}
			}
			fmt.Printf("   %-10s | %5d | %s\n", severity, count, status)
		}
	}
	fmt.Println()
}