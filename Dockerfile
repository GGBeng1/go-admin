FROM alpine:latest 

WORKDIR / 

ADD main / 

EXPOSE 3000 

ENTRYPOINT ["./main"] 