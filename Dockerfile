#Using latest golang base image  
FROM golang:latest

#Maintainer
LABEL maintainer="Sai Kiran Ambati <saikiranambati942@gmail.com>"

#setting workdir
RUN mkdir -p /usr/local/go/src/loanprocess-rest-service
WORKDIR /usr/local/go/src/loanprocess-rest-service

#Copying files
ADD . /usr/local/go/src/loanprocess-rest-service

#Build app
RUN go build -o loanprocessor cmd/loanprocessor/main.go

#Ports
EXPOSE 8080

# run the binary
CMD ["./loanprocessor"]

