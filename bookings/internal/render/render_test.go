package render

import (
	"net/http"
	"testing"

	"github.com/liber/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		t.Error(err)
	}
	resutl := AddDefaultData(td, r)

	if resutl == nil {
		t.Error("failed")
	}
}
