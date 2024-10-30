package main

import (
	"fmt"
	"strings"
)

// EmailInfo extracts the domain and TLD from the given email address.
func EmailInfo(email string) string {
	// Split the email into parts using "@" and "."
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "Invalid email format"
	}
	
	domainParts := strings.Split(parts[1], ".")
	if len(domainParts) < 2 {
		return "Invalid email format"
	}

	// The domain is the first part after "@" and before the first "."
	domain := domainParts[0]
	// The TLD is everything after the first "."
	tld := strings.Join(domainParts[1:], ".")

	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// Main function for testing
func main() {
	fmt.Println(EmailInfo("admin@yahoo.com"))        // Expected: Domain: yahoo dan TLD: com
	fmt.Println(EmailInfo("ptmencaricintasejati@gmail.co.id")) // Expected: Domain: gmail dan TLD: co.id
}
