FROM alpine:edge AS build

RUN apk add --no-cache --update go gcc g++

WORKDIR /go/src/github.com/naufalfmm/cryptocurrency-price-api

COPY go.mod go.sum ./

RUN GO111MODULE=on go mod download

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o cryptocurrency-price-api


FROM alpine:edge

WORKDIR /usr/src
COPY --from=build /go/src/github.com/naufalfmm/cryptocurrency-price-api/cryptocurrency-price-api cryptocurrency-price-api

CMD ["./cryptocurrency-price-api"]