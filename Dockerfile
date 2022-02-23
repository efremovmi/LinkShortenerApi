FROM golang:1.16-alpine

WORKDIR LinkShortenerApi
RUN mkdir -p LinkShortenerApi/pkg/errors
RUN mkdir -p LinkShortenerApi/pkg/internal
RUN mkdir -p LinkShortenerApi/pkg/store_with_db
RUN mkdir -p LinkShortenerApi/pkg/internal/store_without_db

COPY . ./

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ARG FLAG_STORE_WITH_DATABASE=true
ENV FLAG_STORE_WITH_DATABASE="${FLAG_STORE_WITH_DATABASE}"


RUN go build -o /link_shortener_api

EXPOSE 8080

CMD [ "/link_shortener_api" ]
