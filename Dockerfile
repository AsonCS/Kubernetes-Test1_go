FROM golang:1.22

WORKDIR /home/go/app

COPY . .

ENV PORT=80

RUN go build -o server ./src/

EXPOSE 80

CMD ["./server"]
