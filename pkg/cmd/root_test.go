package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRedishCommand(t *testing.T) {
	cmd := NewRedishCommand()

	if err := cmd.Execute(); err != nil {
		assert.NoError(t, err)
	}
}
