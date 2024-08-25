package earthbucks

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Domain represents a domain name.
type Domain struct {
	domainStr string
}

// NewDomain creates a new Domain instance.
func NewDomain(domainStr string) *Domain {
	return &Domain{domainStr: domainStr}
}

// FromString creates a Domain from a string.
func FromStringDomain(domainStr string) *Domain {
	return NewDomain(domainStr)
}

// ToString returns the domain as a string.
func (d *Domain) ToString() string {
	return d.domainStr
}

// IsValid checks if the domain is valid based on our simple validation rules.
func (d *Domain) IsValid() bool {
	return isValidDomain(d.domainStr)
}

// isValidDomain validates the domain string with the specified rules.
func isValidDomain(domain string) bool {
	domainStr := strings.TrimSpace(domain)
	if len(domainStr) < 4 {
		return false
	}
	if strings.HasPrefix(domainStr, ".") || strings.HasSuffix(domainStr, ".") {
		return false
	}
	if !strings.Contains(domainStr, ".") {
		return false
	}
	if strings.Contains(domainStr, "..") {
		return false
	}
	domainParts := strings.Split(domainStr, ".")
	if len(domainParts) < 2 {
		return false
	}
	if len(domainParts) > 10 {
		return false
	}
	for _, part := range domainParts {
		if len(part) > 63 {
			return false
		}
		if matched, _ := regexp.MatchString("^[a-z0-9]+$", part); !matched {
			return false
		}
	}
	return true
}

// DomainToBaseURL converts the domain to a base URL based on certain rules.
func DomainToBaseURL(domain string) string {
	if strings.Contains(domain, "localhost") {
		parts := strings.Split(domain, ".")
		if len(parts) > 1 {
			portStr := parts[0]
			port, err := strconv.Atoi(portStr)
			if err == nil && strings.HasSuffix(domain, "localhost") {
				return fmt.Sprintf("http://%d.localhost:%d", port, port)
			}
		}
	}
	return fmt.Sprintf("https://%s", domain)
}