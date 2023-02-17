FROM ubuntu:22.04
RUN mkdir -p /scripts
COPY main.sh /scripts
WORKDIR /scripts
RUN apt update
RUN apt install curl -y
RUN apt install jq -y
RUN chmod +x main.sh
EXPOSE 9000
CMD ["./main.sh"]
