package store_without_db

import (
	"github.com/genridarkbkru/LinkShortenerApi/pkg/errors"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/internal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStore_Create(t *testing.T) {
	store := RepositoryWithHashTables{}
	store.NewDB("", "")

	testCases := []struct {
		name             string
		payload          string
		expectedCode     int
		expectedShortUrl string
		expectedError    error
	}{
		{
			name:             "record created",
			payload:          "example",
			expectedCode:     http.StatusCreated,
			expectedShortUrl: internal.GetShortUrl("example"),
			expectedError:    nil,
		},
		{
			name:             "the record is already in the table",
			payload:          "example",
			expectedCode:     http.StatusOK,
			expectedShortUrl: internal.GetShortUrl("example"),
			expectedError:    nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actualShortUrl, actualErr, actualCodeErr := store.Create(tc.payload)
			asserts(t, tc.expectedError, actualErr,
				tc.expectedCode, actualCodeErr,
				tc.expectedShortUrl, actualShortUrl)

		})
	}

}

func TestStore_FindByShortUrl(t *testing.T) {
	store := RepositoryWithHashTables{}
	store.NewDB("", "")
	store.Create("example")

	testCases := []struct {
		name          string
		payload       string
		expectedCode  int
		expectedUrl   string
		expectedError error
	}{
		{
			name:          "record in table",
			payload:       internal.GetShortUrl("example"),
			expectedCode:  http.StatusFound,
			expectedUrl:   "example",
			expectedError: nil,
		},
		{
			name:          "record not in table",
			payload:       internal.GetShortUrl("example1"),
			expectedCode:  http.StatusNotFound,
			expectedUrl:   "",
			expectedError: errors.RecordNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actualUrl, actualErr, actualCodeErr := store.FindByShortUrl(tc.payload)
			asserts(t, tc.expectedError, actualErr,
				tc.expectedCode, actualCodeErr,
				tc.expectedUrl, actualUrl)

		})
	}
}

func asserts(t *testing.T, expectedErr, actualErr error,
	expectedCode, actualCode int,
	expectedUrl, actualUrl string) {

	assert.Equal(t, expectedErr, actualErr)
	assert.Equal(t, expectedCode, actualCode)
	assert.Equal(t, expectedUrl, actualUrl)
}
