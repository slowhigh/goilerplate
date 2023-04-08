FROM dpage/pgadmin4:6.21

USER root

RUN mkdir -p /var/lib/pgadmin/storage/user_goilerplate.com
COPY ./docker/pgAdmin4/servers.json /pgadmin4/servers.json
COPY ./docker/pgAdmin4/pgpass /var/lib/pgadmin/storage/user_goilerplate.com/pgpass
RUN chmod 0600 /var/lib/pgadmin/storage/user_goilerplate.com/pgpass

ENV PGADMIN_DEFAULT_EMAIL "user@goilerplate.com"
ENV PGADMIN_DEFAULT_PASSWORD "1234"
ENV PGADMIN_LISTEN_PORT "80"

ENTRYPOINT ["/entrypoint.sh"]