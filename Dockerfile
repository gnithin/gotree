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
ENV GOPATH="/usr/local/go_workspace"
ENV PATH $PATH:$GOPATH:/usr/local/go/bin
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR $GOPATH/src
COPY ./ ./
