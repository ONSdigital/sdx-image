FROM golang:1.21-alpine
COPY . /app
WORKDIR /app
RUN go mod download
RUN go build -o /sdx-image
EXPOSE 5000
CMD [ "/sdx-image" ]
