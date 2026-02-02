package utils

import (
	"fmt"
	"strings"

	"github.com/ttacon/libphonenumber"
)

func ToNationalFormat(phone, countryCode string) (string, error) {
	// Normalize: remove spaces, dashes, parentheses, etc.
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")

	// Parse the phone number with country hint (country
	num, err := libphonenumber.Parse(phone, countryCode)
	if err != nil {
		return "", fmt.Errorf("failed to parse phone number: %v", err)
	}

	// Optional: verify it's actually from the expected country
	region := libphonenumber.GetRegionCodeForNumber(num)
	if strings.ToUpper(region) != strings.ToUpper(countryCode) {
		return "", fmt.Errorf("phone number does not belong to %s (detected: %s)", countryCode, region)
	}

	// Format as NATIONAL → this removes the country code and gives local format
	nationalFormat := libphonenumber.Format(num, libphonenumber.NATIONAL)

	// For Vietnam: it returns "081 831 1855" → we usually want "0818311855" (no spaces)
	// Remove all spaces for clean output (common in VN)
	nationalFormat = strings.ReplaceAll(nationalFormat, " ", "")

	return nationalFormat, nil
}
