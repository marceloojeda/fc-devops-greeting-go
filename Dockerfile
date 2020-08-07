FROM scratch

COPY ./src /go/bin/app

ENTRYPOINT ["/app"]