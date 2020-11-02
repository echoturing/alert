FROM ubuntu:latest
RUN  mkdir -p /data/apps/alert
WORKDIR /data/apps/alert
COPY ./.build/alert-be-linux /data/apps/alert/be/alert-be-linux
ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /data/apps/alert
RUN chmod +x wait-for-it.sh
CMD ["./be/alert-be-linux"]