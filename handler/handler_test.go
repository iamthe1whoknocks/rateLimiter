package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/iamthe1whoknocks/rateLimiter/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	testIP           = "127.0.0.1"
	testMask         = "24"
	testRequestLimit = 3
	testTime         = 1 * time.Minute
)

//Тест rate limiter
func TestLimitMiddleware(t *testing.T) {
	e := echo.New()

	h := New(testMask, testRequestLimit, testTime)

	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	}

	testCases := []struct {
		id   string
		code int
	}{
		{testIP, http.StatusOK},
		{testIP, http.StatusOK},
		{testIP, http.StatusOK},
		{testIP, http.StatusTooManyRequests},
		{testIP, http.StatusTooManyRequests},
		{testIP, http.StatusTooManyRequests},
		{testIP, http.StatusTooManyRequests},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Add(echo.HeaderXRealIP, tc.id)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = h.LimitMiddleware(handler)(c)

		assert.Equal(t, tc.code, rec.Code)
	}

	//проверка времени ожидания после ограничения
	time.Sleep(testTime)

	testCasesAfter := []struct {
		id   string
		code int
	}{
		{testIP, http.StatusOK},
		{testIP, http.StatusOK},
		{testIP, http.StatusOK},
	}

	for _, tc := range testCasesAfter {
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		req.Header.Add(echo.HeaderXRealIP, tc.id)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = h.LimitMiddleware(handler)(c)

		assert.Equal(t, tc.code, rec.Code)
	}

}

//Проверка сброса лимита
func TestDrop(t *testing.T) {
	e := echo.New()

	h := New(testMask, testRequestLimit, testTime)

	req := httptest.NewRequest(http.MethodGet, "/drop", nil)
	req.Header.Add(echo.HeaderXRealIP, testIP)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	_ = h.Drop(c)

	assert.Contains(t, rec.Body.String(), "Dropped")

	emptyLimiter := utils.CreateLimiter(testTime, testRequestLimit)

	testSubnet, err := utils.GetSubnetFromIP(testIP, testMask)
	if err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, emptyLimiter, h.Subnets[testSubnet])

}
