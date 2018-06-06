FROM golang:1.10-alpine

# Prepare app source directory
ENV APP_PATH /go/src/github.com/nyaruka/courier
RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH
COPY . .

# Create Spool directory
RUN mkdir -p /var/spool/courier

# Install courier application
RUN go get -d -v ./...
RUN go install -v ./cmd/...

EXPOSE 80
ENTRYPOINT ["courier"]