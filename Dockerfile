FROM ubuntu:latest

RUN apt-get update

# Installing pre-requisites
RUN apt-get install -y apt-utils
RUN apt-get install -y curl

# Installing golang
RUN cd /tmp/
RUN curl -fsSL https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz -o golang.tar.gz
RUN tar -C /usr/local -xzf golang.tar.gz 

RUN mkdir -p "/usr/local/go_workspace"
ENV GOPATH "/usr/local/go_workspace"

RUN apt-get install -y git

# TODO: Make this more compact
# Add the path to golang workspace and the bin where all the installed stuff go
ENV PATH $PATH:$GOPATH:/usr/local/go/bin:$GOPATH/bin
ENV GOBIN "/usr/local/go/bin"
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH/src/gotree
COPY ./ ./

# Add glide
RUN curl https://glide.sh/get | sh
RUN glide install

# Just run the tests. That'll be cleaner
# Runs all the tests
RUN ./run_tests.sh
