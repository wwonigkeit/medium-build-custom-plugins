FROM golang:1.15-buster as build

WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v 
RUN CGO_ENABLED=0 go build -o /identify-lang -ldflags="-s -w" main.go


FROM scratch
COPY --from=build /identify-lang /

CMD ["/identify-lang"]
