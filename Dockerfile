FROM golang:1.20
WORKDIR /workdir
COPY . /workdir 
RUN go build -o app /workdir
EXPOSE 8080
ENTRYPOINT [ "./app" ]