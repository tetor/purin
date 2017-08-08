package cmd

import (
	"testing"
)

func TestRootCmd(t *testing.T) {
	if !RootCmd.IsAvailableCommand() {
		t.Errorf("Root command can't run")
	}
}
