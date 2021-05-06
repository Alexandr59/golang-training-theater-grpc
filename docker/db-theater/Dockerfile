FROM postgres:alpine
COPY ./db-theater/script.sql /docker-entrypoint-initdb.d/init.sql
EXPOSE 5432