package store

import (
	"database/sql"
	"fmt"
	"github.com/genridarkbkru/LinkShortenerApi/pkg"
	hash "github.com/genridarkbkru/LinkShortenerApi/pkg/internal"
	_ "github.com/lib/pq"
	"net/http"
)

type Store struct {
	psqlconn string
	db       *sql.DB
}

func (r *Store) NewDB(psqlconn string) error {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	r.db = db
	r.psqlconn = psqlconn
	fmt.Println("Connected!")
	return nil
}

func (r *Store) Create(url string) (string, error, int) {
	shortUrl := hash.GetShortUrl(url)
	_, err, _ := r.FindByShortUrl(shortUrl)
	if err == nil {
		return shortUrl, nil, http.StatusOK
	}

	r.db, _ = sql.Open("postgres", r.psqlconn)
	defer r.db.Close()

	var id int
	err = r.db.QueryRow(
		"INSERT INTO tabl_urls (url, short_url) "+
			"Values ($1, $2) RETURNING id",
		url,
		shortUrl,
	).Scan(&id)

	return shortUrl, err, http.StatusCreated
}

func (r *Store) FindByShortUrl(short_url string) (string, error, int) {

	var id int
	var url string

	r.db, _ = sql.Open("postgres", r.psqlconn)
	defer r.db.Close()

	if err := r.db.QueryRow(
		"SELECT id, url  FROM"+
			" tabl_urls WHERE short_url = $1", short_url,
	).Scan(&id, &url); err != nil {

		if err == sql.ErrNoRows {
			return "", pkg.ErrRecordNotFound, http.StatusNotFound
		}

	}
	return url, nil, http.StatusFound
}
