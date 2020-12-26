FROM alpine:latest
#MAINTAINER yinguohua "yinguohua@qiushibaike.com"
COPY app /app
COPY conf.toml /conf.toml

#COPY cmd/web_template.json /usr/web_template.json
#CMD ["/usr/application", "-env", "prod", "-conf", "/usr/web_template.json"]
#CMD ["/usr/application", "-env", "prod", "-conf", "/usr/web_template.json"]
CMD ["/app","-conf","conf.toml"]
EXPOSE 8081

# docker run -d --restart=unless-stopped -p 8081:8081 --name=micro_api micro_api:latest
# /usr/application -env test -conf /usr/web_template.json