package services

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewExampleService(t *testing.T) {
	got := NewExampleService()
	want := &exampleService{}

	assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(got))
}

func Test_exampleService_ExposedFunction(t *testing.T) {
	w := httptest.NewRecorder()
	NewExampleService().ExposedFunction(w)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), "Hello, World - Vanilla\n")
}