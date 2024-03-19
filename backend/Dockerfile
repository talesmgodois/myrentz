# syntax=docker/dockerfile:1

FROM golang:1.21.1

# Set destination for COPY
WORKDIR /application

# Download Go notifications

COPY go.mod go.sum ./
COPY /config/base.yml base.yml
RUN go mod download

RUN apt update
RUN apt install vim -y

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/main.go


# Run
CMD ["./server"]