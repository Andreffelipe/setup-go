FROM golang:1.7 as BUILDER
WORKDIR /usr/app
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /usr/app
COPY --from=BUILDER /usr/app ./
CMD ["./app"] 