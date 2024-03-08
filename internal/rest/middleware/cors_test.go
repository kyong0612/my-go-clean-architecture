package middleware_test

import (
	"net/http"
	test "net/http/httptest"
	"testing"

	"github.com/kyong0612/my-go-clean-architecture/internal/rest/middleware"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCORS(t *testing.T) {
	// TODO:
	e := echo.New()
	req := test.NewRequest(echo.GET, "/", nil)
	res := test.NewRecorder()
	c := e.NewContext(req, res)

	h := middleware.CORS(echo.HandlerFunc(func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	}))

	err := h(c)
	require.NoError(t, err)
	assert.Equal(t, "*", res.Header().Get("Access-Control-Allow-Origin"))
}
