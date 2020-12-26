FROM alpine:latest
#MAINTAINER yinguohua "yinguohua@qiushibaike.com"
COPY app /app
#COPY cmd/web_template.json /usr/web_template.json
#CMD ["/usr/application", "-env", "prod", "-conf", "/usr/web_template.json"]
#CMD ["/usr/application", "-env", "prod", "-conf", "/usr/web_template.json"]
CMD ["/app"]
EXPOSE 8080

# docker run -d --restart=unless-stopped -p 8080:8080 --name=micro_user micro_user:latest
# /usr/application -env test -conf /usr/web_template.json