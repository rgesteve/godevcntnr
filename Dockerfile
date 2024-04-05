FROM docker.io/library/ubuntu:latest

RUN apt update && apt install -y curl nginx
WORKDIR /var/www/html
ADD index.html ./
CMD ["nginx", "-g", "daemon off;"]
