ARG GO_VERSION

FROM golang:${GO_VERSION}

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest 
RUN go install github.com/volatiletech/sqlboiler/v4@latest 
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest