FROM golang

ADD ./apigooglegeo.exe /usr/local/bin/

ENTRYPOINT /usr/local/bin/apigooglegeo.exe

EXPOSE 8090
