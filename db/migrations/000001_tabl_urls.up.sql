CREATE TABLE IF NOT EXISTS tabl_urls(
                                    Id SERIAL PRIMARY KEY,
                                    Url VARCHAR(500) NOT NULL,
                                    Short_url VARCHAR(10) NOT NULL
);