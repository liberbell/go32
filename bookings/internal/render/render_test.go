package render

import (
	"net/http"
	"testing"

	"github.com/liber/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	resutl := AddDefaultData(&td, r)

	if resutl == nil {
		t.Error("failed")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ := session.Load(ctx, r.Header.Get("X-session"))

	r = r.WithContext(ctx)

	return r, nil
}
