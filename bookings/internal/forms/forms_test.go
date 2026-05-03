package forms

import (
	"net/http/httptest"
	"testing"
)

func Testform_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
}
