FROM golang:1.15 as builder
WORKDIR /go/src/github.com/vpoletaev11/avitoParser

# Copy the local project files to the container's workspace.
ADD . /go/src/github.com/vpoletaev11/avitoParser

# Build the project inside the container.
RUN GOOS=linux go build .


# Execute the binary
FROM ubuntu
EXPOSE 8080

RUN apt-get update && apt-get install ca-certificates -y
COPY --from=builder /go/src/github.com/vpoletaev11/avitoParser   /
ENTRYPOINT ["/avitoParser"]
