# exgohttpheaders
Example service to show http headers as json format



DOCKER BUILD
=========================

ALPINE BUILD

docker build -t mkawserm/exgohttpheaders:latest .


SLIM BUILD

docker build -f Dockerfile:slim  -t mkawserm/exgohttpheaders:slim .



RUN CONTAINER
========================

ALPINE BUILD

docker run -p 8080:8080 mkawserm/exgohttpheaders:latest


SLIM BUILD

docker run -p 8080:8080 mkawserm/exgohttpheaders:slim
