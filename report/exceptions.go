package report

import (
    "encoding/json"
    "fmt"
    "os"
)

// Exception represents a finding that should be marked as exception
type Exception struct {
    RuleID string `json:"rule_id"`
    Secret string `json:"secret"`
    File   string `json:"file"`
	Match  string `json:"match"`
}

// LoadExceptions loads exceptions from a JSON file
func LoadExceptions(exceptionsPath string) ([]Exception, error) {
    if exceptionsPath == "" {
        return []Exception{}, nil
    }

    file, err := os.Open(exceptionsPath)
    if err != nil {
        return nil, fmt.Errorf("could not open exceptions file: %w", err)
    }
    defer file.Close()

    var exceptions []Exception
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&exceptions)
    if err != nil {
        return nil, fmt.Errorf("could not parse exceptions JSON: %w", err)
    }

    return exceptions, nil
}

// IsException checks if a finding matches any exception
func IsException(finding Finding, exceptions []Exception) bool {
    for _, exception := range exceptions {
        // Exact match on all three fields: rule_id, secret, file
        if exception.RuleID == finding.RuleID &&
            exception.Secret == finding.Secret &&
            exception.File == finding.File &&
			exception.Match == finding.Match {
            return true
        }
    }
    return false
}

// MarkExceptions updates severity to "Exception" for matching findings
func MarkExceptions(findings []Finding, exceptions []Exception) []Finding {
    exceptionCount := 0
    for i := range findings {
        if IsException(findings[i], exceptions) {
            findings[i].Severity = "Exception"
            exceptionCount++
        }
    }
    
    if exceptionCount > 0 {
        fmt.Printf("Marked %d findings as exceptions\n", exceptionCount)
    }
    
    return findings
}




//Exception FIle FORMAT
// [
//   {
//     "rule_id": "aws-access-key",
//     "secret": "AKIAIOSFODNN7EXAMPLE",
//     "file": "src/config.py"
//   },
//   {
//     "rule_id": "generic-api-key", 
//     "secret": "apikey_12345",
//     "file": "api/config.js"
//   },
//   {
//     "rule_id": "slack-token",
//     "secret": "xoxb-1234567890",
//     "file": "config.json"
//   }
// ]