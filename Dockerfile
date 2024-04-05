FROM docker.io/library/ubuntu:latest

RUN apt update && apt install -y curl nginx
CMD ["nginx", "-g", "daemon off;"]
