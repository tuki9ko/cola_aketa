ARG GO_VERSION=${GO_VERSION}
ARG API_CODE_WORKDIR=${API_CODE_WORKDIR}

FROM golang:${GO_VERSION}

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest