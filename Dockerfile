FROM golang:1.20-alpine AS build

RUN apk --no-cache add bash make gcc gettext musl-dev

WORKDIR /app

# dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#build

COPY . ./
RUN go build -o ./bin/app cmd/main.go


FROM alpine

COPY --from=build /app/bin/app /
COPY configs/config.yml /config.yml

CMD ["/app"]