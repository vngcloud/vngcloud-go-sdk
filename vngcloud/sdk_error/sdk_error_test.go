package sdk_error

import (
	lregexp "regexp"
	"testing"
)

func TestRegex(t *testing.T) {
	regexErrorSubnetNotFound = lregexp.MustCompile(patternNotAllowDelete)
	patternServerCreating1 := "cannot delete server with status creating"                // "Server is creating"
	patternServerCreatingBilling1 := "cannot delete server with status creating-billing" // "The number of servers exceeds the quota"
	patternServerDeleting1 := "cannot delete server with status deleting"
	patternFake := "cannot delete server with status fake"

	if regexErrorSubnetNotFound.FindString(patternServerCreating1) == "" {
		t.Errorf("patternServerCreating should match patternNotAllowDelete")
	}

	if regexErrorSubnetNotFound.FindString(patternServerCreatingBilling1) == "" {
		t.Errorf("patternServerCreatingBilling should match patternNotAllowDelete")
	}

	if regexErrorSubnetNotFound.FindString(patternServerDeleting1) == "" {
		t.Errorf("patternServerDeleting should match patternNotAllowDelete")
	}

	if regexErrorSubnetNotFound.FindString(patternFake) != "" {
		t.Errorf("patternFake should not match patternNotAllowDelete")
	}
}
