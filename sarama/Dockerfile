FROM alpine
RUN mkdir -p /app/config/prod
COPY consume /app
WORKDIR /app
#COPY config/prod/app.toml /app/config/prod/app.toml
#COPY templates /app/templates
RUN chmod 777 /app/consume
EXPOSE 8081
ENV go_env=prod
ENTRYPOINT ["/app/consume"]
