# STAGE1---------------Build the dedicated q3a server in this stage
FROM golang:1.18-alpine as builder


WORKDIR /app

# Build the go server
COPY . .
RUN go mod vendor 
RUN go build -mod=vendor




#  STAGE2-------------Build final Image-------------
FROM alpine:3.17.0 AS imager
LABEL "Maintainer" "Tuhin Sengupta <tuhin3737@gnail.com>"


# create a non root user
RUN adduser server -D
WORKDIR /app
RUN mkdir prefix-wordlists

# copy from builder stage
COPY --chown=server --from=builder /app/truecaller-prefix /app


# switch user
USER server
EXPOSE 5000/tcp
# Run go server
CMD ["/app/truecaller-prefix"]