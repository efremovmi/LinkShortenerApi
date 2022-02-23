package store_with_db

import (
	"database/sql"
	"fmt"
	"github.com/genridarkbkru/LinkShortenerApi/pkg/errors"
	hash "github.com/genridarkbkru/LinkShortenerApi/pkg/internal"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type RepositoryWithDB struct {
	psqlconn  string
	db        *sql.DB
	tableName string
}

func (r *RepositoryWithDB) NewDB(psqlconn, tableName string) {
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(errors.IncorrectParamsConnectBD.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(errors.BDnotWorking.Error())
	}
	r.tableName = tableName
	r.db = db
	r.psqlconn = psqlconn
	log.Println("Database connection was successful!")
}

func (r *RepositoryWithDB) Create(url string) (string, error, int) {
	shortUrl := hash.GetShortUrl(url)
	_, err, _ := r.FindByShortUrl(shortUrl)
	if err == nil {
		return shortUrl, nil, http.StatusOK
	}

	r.db, _ = sql.Open("postgres", r.psqlconn)
	defer r.db.Close()

	var id int

	query := fmt.Sprintf("INSERT INTO %s(url, short_url) ", r.tableName)

	err = r.db.QueryRow(query+"Values ($1, $2) RETURNING id",
		url,
		shortUrl,
	).Scan(&id)

	return shortUrl, err, http.StatusCreated
}

func (r *RepositoryWithDB) FindByShortUrl(short_url string) (string, error, int) {

	var id int
	var url string

	r.db, _ = sql.Open("postgres", r.psqlconn)
	defer r.db.Close()

	query := fmt.Sprintf("SELECT id, url  FROM %s", r.tableName)

	if err := r.db.QueryRow(query+" WHERE short_url = $1",
		short_url,
	).Scan(&id, &url); err != nil {

		if err == sql.ErrNoRows {
			return "", errors.RecordNotFound, http.StatusNotFound
		}

	}
	return url, nil, http.StatusFound
}
