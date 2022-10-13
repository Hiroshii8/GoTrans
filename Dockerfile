FROM golang:1.19 as build

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/gotrans

FROM gcr.io/distroless/static-debian11

EXPOSE 8080
EXPOSE 8081

COPY --from=build /go/bin/gotrans /
CMD ["/gotrans"]
