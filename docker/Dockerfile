FROM golang:1.16.6-buster
COPY . .
RUN go mod vendor
RUN go build .
CMD [ "./proj-web" ]