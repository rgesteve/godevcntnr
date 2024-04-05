podman build -t nginx_text .
# Cannot expose port 80 on the Windows side, so using 7000 instead
podman run -d -p 7000:80 localhost/nginx_test