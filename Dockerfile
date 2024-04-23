FROM golang:bullseye

ARG VERSION="0.0.0"

COPY main.go /
COPY go.mod /
COPY go.sum /

RUN echo "building stun server" \
 	&& cd / && go build main.go

COPY entry.sh /


ENTRYPOINT ["/entry.sh"]
RUN ["chmod", "+x", "/entry.sh"]
