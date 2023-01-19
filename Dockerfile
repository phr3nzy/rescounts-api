FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o rescounts-api

FROM alpine

WORKDIR /app

RUN apk --no-cache add tini

ENV APP_ENV production

ENV UID=10001

RUN addgroup -S rescounts-api-service

RUN adduser -D \    
	--disabled-password \    
	--gecos "" \    
	--home "/nonexistent" \    
	--shell "/sbin/nologin" \    
	--no-create-home \    
	--uid "${UID}" \    
	rescounts-api-user \ 
	-G rescounts-api-service

USER rescounts-api-user

COPY --chown=rescounts-api-user:rescounts-api-service --from=builder /app/rescounts-api /app/rescounts-api

ENTRYPOINT [ "/sbin/tini", "--" ]

CMD ["/app/rescounts-api"]
