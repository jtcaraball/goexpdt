# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.21.6 AS build-stage

RUN apt-get update && apt-get install unzip

WORKDIR /goexpdt

# Get kissat repo
ADD https://github.com/arminbiere/kissat/archive/refs/tags/rel-3.1.1.zip rel-3.1.1.zip
RUN unzip rel-3.1.1.zip
RUN mv kissat-rel-3.1.1 kissat
RUN rm rel-3.1.1.zip

# Compile solver binary
WORKDIR /goexpdt/kissat
RUN ./configure
RUN make

# Move binary to root and remove left overs
WORKDIR /goexpdt
RUN mv kissat/build/kissat /kissat
RUN rm -rf kissat

# Copy project
COPY . .

FROM build-stage AS run-tests-stage
# Run the tests in the container
RUN go test ./...
