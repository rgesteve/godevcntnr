FROM docker.io/library/ubuntu:latest

RUN apt update && apt install -y curl nginx
COPY nginx.conf /etc/nginx/sites-enabled/test.conf
#WORKDIR /var/www/html
WORKDIR /var/www/test
ADD index.html ./
CMD ["nginx", "-g", "daemon off;"]
