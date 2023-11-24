FROM golang

WORKDIR /apps

COPY . .

RUN go build -o main main.go

CMD [ "./main" ]