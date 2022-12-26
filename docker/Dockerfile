ARG GO_VERSION
FROM golang:${GO_VERSION}

ARG API_CODE_WORKDIR
ARG API_CODE_LOCAL_PATH
WORKDIR ${API_CODE_WORKDIR}

RUN mkdir -p ${API_CODE_WORKDIR}
ADD ./server/go.mod ${API_CODE_WORKDIR}
ADD ./server/.air.toml ${API_CODE_WORKDIR}

RUN go mod tidy && \
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	go install github.com/volatiletech/sqlboiler/v4@latest && \
	go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest && \
	go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml"]