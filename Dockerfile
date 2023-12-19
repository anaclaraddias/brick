FROM golang:1.21.3

RUN go install github.com/cosmtrek/air@latest
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

ENV PATH="$PATH:/go/bin"

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

CMD [ "air" ]
