package pkg

type Repository interface {
	NewDB(psqlconn, tableName string)
	Create(url string) (string, error, int)
	FindByShortUrl(short_url string) (string, error, int)
}
