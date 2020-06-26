FROM golang:1.14-alpine

ENV APP_NAME mortred

RUN mkdir /go/src/${APP_NAME}

COPY ./ /go/src/${APP_NAME}

WORKDIR /go/src/${APP_NAME}

# Install curl
RUN apk --no-cache add \
	bash \
	curl \
	git \
	gcc \
	g++ \
	openssh-client

RUN go mod vendor \
    && go install

EXPOSE 8000

CMD ["mortred"]
