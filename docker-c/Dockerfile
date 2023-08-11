FROM ubuntu:latest

RUN apt -y update
RUN apt -y install gcc

COPY . /

RUN gcc main.c -o main
CMD ["./main"]
