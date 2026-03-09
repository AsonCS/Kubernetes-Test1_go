FROM golang:1.23

WORKDIR /home/go/app

COPY . .

ENV PORT=80

RUN go mod init github.com/AsonCS/test1_go || :
RUN go build -o server ./src/

EXPOSE 80

CMD ["./server"]
