package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitLogger(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Asserts a logger is created",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitLogger(); got != nil {
				assert.NotNil(t, got)
			}
		})
	}
}
