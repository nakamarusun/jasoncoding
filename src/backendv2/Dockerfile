FROM golang:1.18-alpine
WORKDIR /app

# Install make
RUN apk add make

# Go packages for installing deps
COPY go.sum ./
COPY go.mod ./
COPY Makefile ./

# Install dependencies and build
RUN make deps
COPY . .
RUN make build


# Use another image and run the app
FROM alpine:latest
RUN apk add imagemagick

WORKDIR /app

# Copy from the first build
COPY --from=0 /app/jcbackendv2 ./

# Copy fonts
COPY ./cool/fonts ./fonts
ENV FONT_PATH=/app/fonts

ENV ENVIRONMENT=production
ENV GIN_MODE=release

CMD ./jcbackendv2
