FROM golang:1.10.0

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/minhajuddinkhan/ttd
WORKDIR /go/src/github.com/minhajuddinkhan/ttd/cli/ttd
# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get
RUN go install github.com/minhajuddinkhan/minhajuddinkhan/ttd/cli/ttd


#This entry point should be used in prod environment
# Run the outyet command by default when the container starts.
#ENTRYPOINT /go/bin/comments

# Document that the service listens on port 8080.
EXPOSE 3000

CMD [ "go", "run", "/go/src/github.com/minhajuddinkhan/ttd/cli/ttd/main.go -a localhost:3000" ]