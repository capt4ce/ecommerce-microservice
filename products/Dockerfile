FROM golang:1.9.2-alpine3.6 AS build

ENV ROOT_SRC github.com/capt4ce/ecommerce-microservice
ENV SRV products
ENV SRV_SRC $ROOT_SRC/$SRV
RUN echo $ROOT_SRC
RUN echo $SRV_SRC

# Install dependencies
RUN apk add --no-cache git mercurial
# RUN go get gopkg.in/mgo.v2
# RUN go get github.com/gorilla/mux
RUN go get github.com/labstack/echo
RUN go get github.com/labstack/echo/middleware
RUN go get github.com/sirupsen/logrus

# copy the local package files to the container workspace
ADD . /go/src/$SRV_SRC

# Setting up working directory
WORKDIR /go/src/$SRV_SRC

# RUN go get -d ./ ...

# Build the users command inside the container.
RUN go install $SRV_SRC

# Run the users microservice when the container starts.
ENTRYPOINT ["/go/bin/$SRV"]

# Service listens on port 8080.
EXPOSE 8080
