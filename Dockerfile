FROM golang:1.17-buster
## Create app folder. It stored source files.
RUN mkdir /app
## Copy everything in the root directory into our /app directory that we created previous step.
ADD . /app
## Specify the working directory
WORKDIR /app
## Build to compile the binary executable of our Go program
## Build to compile the binary executable of our Go program
RUN go mod tidy
RUN go get ./...
RUN go build -o main .
#EXPOSE 8080
CMD ["/app/main"]
