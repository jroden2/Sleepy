package services

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewExampleService(t *testing.T) {
	got := NewExampleService()
	want := &exampleService{}

	assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(got))
}