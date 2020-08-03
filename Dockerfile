FROM golang:alpine

RUN apk update
RUN apk add --no-cache gcc
RUN apk add --no-cache libc-dev
RUN apk add --no-cache make
RUN apk add --no-cache gzip

# Set the current working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. they will be cached of the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the WORKDIR inisde the container
COPY . .

# Command to run the executable
RUN make gen
CMD make PORT=$PORT serve
