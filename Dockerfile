FROM golang
MAINTAINER Luke Atherton "lukeatherton@hivebase.io"

# Install dependencies
run apt-get update
run apt-get install -y curl git

WORKDIR /go/src/github.com/lukeatherton/healthchecker
ADD . /go/src/github.com/lukeatherton/healthchecker

# Add files
ADD ./bin/boot.sh             /boot.sh

EXPOSE 8001

# Start the container
RUN chmod +x /boot.sh
CMD /boot.sh

#Install App
RUN go get
RUN go build
