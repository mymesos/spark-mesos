FROM mymesos/java8-mesos:0.25.0
MAINTAINER Dave Choi <goodoi09@gmail.com>

COPY . /usr/src/spark-mesos
WORKDIR /usr/src/spark-mesos

ENV SPAKR_NAME spark-1.5.2-bin-hadoop2.4

RUN wget http://d3kbcqa49mib13.cloudfront.net/$SPAKR_NAME.tgz && \
tar -xvzf $SPAKR_NAME.tgz && rm -f $SPAKR_NAME.tgz
RUN mv start_server $SPAKR_NAME/

WORKDIR $SPAKR_NAME

CMD ./start_server