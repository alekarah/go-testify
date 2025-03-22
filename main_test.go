package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestCorrectRequest проверяет что
// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
func TestCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "status should return 200 OK")

	assert.NotEmpty(t, responseRecorder.Body.String(), "response body must not be empty")
}

// TestWrongCity проверяет что
// Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestWrongCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=berlin", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code, "status should return 400 Bad Request")

	assert.Equal(t, "wrong city value", responseRecorder.Body.String(), "message should return 'wrong city value'")
}

// TestCountMoreThanTotal проверяет что
// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
func TestCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	list := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, list, totalCount, "all available cafes should return")
}
