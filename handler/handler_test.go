package handler

import (
	"bootcamp/config"
	"bootcamp/mock"
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Wallets(t *testing.T) {
	t.Run("get with all wallet", func(t *testing.T) {
		service := mock.NewMockIWalletService(gomock.NewController(t))
		serviceResult := map[string]int{"test1": 5, "test2": 5}
		service.EXPECT().
			GetWallets().
			Return(serviceResult).
			Times(1)

		handler := NewHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.Wallets(w, r)

		myResult := "{\"test1\":5,\"test2\":5}"

		assert.Equal(t, w.Body.String(), myResult)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))

	})
	t.Run("get with username wallet", func(t *testing.T) {
		service := mock.NewMockIWalletService(gomock.NewController(t))
		serviceResult := 5
		service.EXPECT().
			GetWalletsByID("test1").
			Return(serviceResult, http.StatusOK).
			Times(1)

		handler := NewHandler(service)

		r := httptest.NewRequest(http.MethodGet, "/test1", nil)
		w := httptest.NewRecorder()
		handler.Wallets(w, r)

		myResult := "5"

		assert.Equal(t, w.Body.String(), myResult)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))

	})

	t.Run("put with create wallet", func(t *testing.T) {
		service := mock.NewMockIWalletService(gomock.NewController(t))
		serviceResult := map[string]int{"test1": config.C.InitialBalanceAmount}

		service.EXPECT().
			WalletCreate("test1").
			Return(serviceResult, http.StatusOK).
			Times(1)

		handler := NewHandler(service)

		r := httptest.NewRequest(http.MethodPut, "/test1", nil)
		w := httptest.NewRecorder()
		handler.Wallets(w, r)

		assert.Equal(t, w.Body.String(), "")
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)

	})

	t.Run("post with update wallet", func(t *testing.T) {
		service := mock.NewMockIWalletService(gomock.NewController(t))

		service.EXPECT().
			WalletTransaction("test1", 50).
			Return(http.StatusInternalServerError).
			Times(1)

		handler := NewHandler(service)
		body := map[string]int{"balance": 50}
		lastBody, _ := json.Marshal(body)

		r := httptest.NewRequest(http.MethodPost, "/test1", bytes.NewReader(lastBody))
		w := httptest.NewRecorder()
		handler.Wallets(w, r)

		assert.Equal(t, w.Body.String(), "")
		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)

	})
	t.Run("other ", func(t *testing.T) {

		handler := NewHandler(nil)

		r := httptest.NewRequest(http.MethodHead, "/test1", nil)
		w := httptest.NewRecorder()
		handler.Wallets(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusNotImplemented)

	})
}
