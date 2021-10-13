FROM golang:1.16-alpine AS build
WORKDIR /src
COPY . .
RUN go build -o /out/hello .

FROM alpine:3.13
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=build /out/hello .
CMD ["./hello"]