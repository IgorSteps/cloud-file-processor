package routes_test

import (
	"app/internal/adapters/rest/routes"
	mocks_routes "app/mocks/routes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test handler to return from the factory.
var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
})

func TestNewRouter(t *testing.T) {
	// Assemble
	mockHandlerFactory := newHappyHandlerFactory(t)
	middlewares := []func(http.Handler) http.Handler{}

	// Act
	r := routes.NewRouter(mockHandlerFactory, middlewares)

	// Assert
	assert.NotNil(t, r)
	mockHandlerFactory.AssertExpectations(t)
}

func newHappyHandlerFactory(t *testing.T) *mocks_routes.HandlerFactory {
	mockHandlerFactory := mocks_routes.NewHandlerFactory(t)
	mockHandlerFactory.EXPECT().UploadFile().Return(testHandler).Once()
	return mockHandlerFactory
}
