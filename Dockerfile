FROM golang

ADD ./apidarksky.exe /usr/local/bin/

ENTRYPOINT /usr/local/bin/apidarksky.exe

EXPOSE 8090
