FROM golang:1.23.3-alpine3.20 AS build

WORKDIR /app

COPY . .
RUN addgroup -S nonroot && adduser -S nonroot -G nonroot

RUN GO111ENABLED=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    go -C "cmd/" build -ldflags "-s -w" -o /app/runs/program

FROM scratch

COPY --from=build /etc/passwd /etc/passwd

USER nonroot

COPY --from=build /app/config ./config
COPY --from=build /app/runs/program .

EXPOSE 3030

CMD ["./program"]

# docker build -f ./docker/Dockerfile -t api  