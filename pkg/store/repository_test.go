package store_with_db

import (
	"database/sql"
	"fmt"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/errors"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/internal"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestStore_Create(t *testing.T) {
	store := RepositoryWithDB{}
	store.psqlconn, store.tableName = internal.GetSqlconnAndTableName()
	store.NewDB(store.psqlconn, store.tableName)

	var err error
	defer func() {
		store.db, err = sql.Open("postgres", store.psqlconn)
		assert.NoError(t, err)
		store.db.Exec(fmt.Sprintf("truncate %s;", store.tableName))
		store.db.Close()
	}()

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
		{
			name:             "record created",
			payload:          "example1",
			expectedCode:     http.StatusCreated,
			expectedShortUrl: internal.GetShortUrl("example1"),
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
	store := RepositoryWithDB{}
	store.psqlconn, store.tableName = internal.GetSqlconnAndTableName()
	store.NewDB(store.psqlconn, store.tableName)

	var err error
	defer func() {
		store.db, err = sql.Open("postgres", store.psqlconn)
		assert.NoError(t, err)
		store.db.Exec(fmt.Sprintf("truncate %s;", store.tableName))
		store.db.Close()
	}()

	testCases := []struct {
		name            string
		payload         string
		expectedCode    int
		expectedUrl     string
		expectedError   error
		makeCreateInEnd string
	}{
		{
			name:            "record not in table",
			payload:         internal.GetShortUrl("example"),
			expectedCode:    http.StatusNotFound,
			expectedUrl:     "",
			expectedError:   errors.RecordNotFound,
			makeCreateInEnd: "example",
		},
		{
			name:            "record in table",
			payload:         internal.GetShortUrl("example"),
			expectedCode:    http.StatusFound,
			expectedUrl:     "example",
			expectedError:   nil,
			makeCreateInEnd: "",
		},
		{
			name:            "record not in table",
			payload:         internal.GetShortUrl("example1"),
			expectedCode:    http.StatusNotFound,
			expectedUrl:     "",
			expectedError:   errors.RecordNotFound,
			makeCreateInEnd: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			actualUrl, actualErr, actualCodeErr := store.FindByShortUrl(tc.payload)
			asserts(t, tc.expectedError, actualErr,
				tc.expectedCode, actualCodeErr,
				tc.expectedUrl, actualUrl)
			if len(tc.makeCreateInEnd) > 0 {
				store.Create(tc.makeCreateInEnd)
			}
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
