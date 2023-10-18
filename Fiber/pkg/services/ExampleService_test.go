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

func Test_exampleService_ExposedFunction(t *testing.T) {
	gotStatus, bodyContent := NewExampleService().ExposedFunction()

	assert.Equal(t, gotStatus, 200)
	assert.Equal(t, bodyContent, struct{ Status string }{Status: "Hello, World. - Fiber"})
}
