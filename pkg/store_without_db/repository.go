package store_without_db

import (
	"github.com/genridarkbkru/LinkShortenerApi/pkg/errors"
	hash "github.com/genridarkbkru/LinkShortenerApi/pkg/internal"
	"net/http"
)

type RepositoryWithHashTables struct {
	tableName   string
	ShortToLong map[string]string
	LongToShort map[string]string
}

func (r *RepositoryWithHashTables) NewDB(psqlconn, tableName string) {
	r.tableName = ""
	r.ShortToLong = make(map[string]string)
	r.LongToShort = make(map[string]string)
}

func (r *RepositoryWithHashTables) Create(url string) (string, error, int) {

	if _, ok := r.LongToShort[url]; ok {
		return r.LongToShort[url], nil, http.StatusOK
	}

	shortUrl := hash.GetShortUrl(url)

	r.ShortToLong[shortUrl] = url
	r.LongToShort[url] = shortUrl

	return shortUrl, nil, http.StatusCreated
}

func (r *RepositoryWithHashTables) FindByShortUrl(short_url string) (string, error, int) {

	if url, ok := r.ShortToLong[short_url]; ok {
		return url, nil, http.StatusFound
	}
	return "", errors.RecordNotFound, http.StatusNotFound
}
