package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arthurc/bokfor/internal"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	e := echo.New()
	e.Renderer = &internal.Templates
    req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, index(c))

	assert.Equal(t, rec.Code, http.StatusOK)
	assert.Contains(t, rec.Body.String(), "<title>Bokfor</title>")
}
