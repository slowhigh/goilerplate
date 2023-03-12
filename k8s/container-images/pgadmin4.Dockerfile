FROM dpage/pgadmin4

USER root

RUN mkdir -p  /var/lib/pgadmin/storage/user_goilerplate.com

COPY ./k8s/config/servers.json /pgadmin4/servers.json

COPY ./k8s/config/pgpass /var/lib/pgadmin/storage/user_goilerplate.com/pgpass

RUN chmod 0600 /var/lib/pgadmin/storage/user_goilerplate.com/pgpass

ENTRYPOINT ["/entrypoint.sh"]