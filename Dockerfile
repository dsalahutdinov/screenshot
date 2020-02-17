FROM golang:1-buster
LABEL maintainer="Dmitry Salahutdinov <dsalahutdinov@gmail.com>"


RUN apt-get update -qq

WORKDIR /app
COPY . .

RUN go build -v -o screenshot .

FROM debian:buster-slim
LABEL maintainer="Dmitry Salahutdinov <dsalahutdinov@gmail.com>"

RUN apt-get update -qq

RUN apt-get install -y vim wget gnupg2

RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list'
RUN apt-get update && apt-get install -y google-chrome-stable
    
COPY ./fonts/* /usr/share/fonts/truetype/


COPY --from=0 /app/screenshot /usr/local/bin/
CMD ["screenshot"]

EXPOSE 8080
