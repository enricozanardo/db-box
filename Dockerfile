# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
WORKDIR /app
# Set an env var that matches github repo name,
ENV SRC_DIR=/go/src/github.com/onezerobinary/web-box
# Add the source code:
ADD . $SRC_DIR
# Build it:
RUN go get -v -u github.com/Masterminds/glide
RUN export PATH=$SRC_DIR/bin:$PATH
RUN cd $SRC_DIR; glide install

RUN cd $SRC_DIR; go build -o web-box; cp web-box /app/
ENTRYPOINT ["./web-box"]