FROM golang:1.22.1

# set working directory
WORKDIR /go/src/app

# copy the source code 
COPY . .


#expose the port

EXPOSE 8080

#Build the go app

RUN go build -o main cmd/main.go

#Run the executable 

CMD [ "./main" ]
