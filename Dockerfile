FROM golang:alpine

RUN apk update && apk add --no-cache gcc && apk add --no-cache libc-dev && apk add --no-cache make && apk add --no-cache gzip

# Set the current working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. they will be cached of the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the WORKDIR inisde the container
COPY . .

# Exporse port 3000 or 8000 to the outisde world
EXPOSE 8080

# Command to run the executable
CMD ["make", "gen"]
CMD ["make", "serve"]
