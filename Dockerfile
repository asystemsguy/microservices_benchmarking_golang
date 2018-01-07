FROM golang:latest 
RUN apt-get update
RUN apt-get install -y sysbench
RUN mkdir /app 
COPY  src/ /app/ 
WORKDIR /app 
RUN go build -o server . 
CMD ["/app/server"]
