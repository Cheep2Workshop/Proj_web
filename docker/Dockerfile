FROM golang:1.16.6-buster
FROM golang:1.16
COPY . .
RUN go mod vendor
RUN go build .
CMD [ "./proj-web" ]