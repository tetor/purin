//
package cmd

import (
	"testing"
)

func TestGetCmd(t *testing.T) {
	if !GetCmd.IsAvailableCommand() {
		t.Errorf("Root command can't run")
	}
}
