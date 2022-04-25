FROM scylladb/scylla-nightly:latest

RUN echo -e "#!/usr/bin/env bash\n\
set -eu -o pipefail\n\
exec /usr/bin/scylla-manager-agent $@" > /scylla-manager-agent-docker-entrypoint.sh

RUN echo "[program:scylla-manager-agent]\n\
command=/scylla-manager-agent-docker-entrypoint.sh\n\
autorestart=true\n\
stdout_logfile=/dev/stdout\n\
stdout_logfile_maxbytes=0\n\
stderr_logfile=/dev/stderr\n\
stderr_logfile_maxbytes=0" > /etc/supervisord.conf.d/scylla-manager-agent.conf

RUN apt-get update && \
    apt-get install -y gnupg2 && \
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 5e08fbd8b5d6ec9c && \
    wget -O /etc/apt/sources.list.d/scylla-manager.list http://downloads.scylladb.com/deb/ubuntu/scylladb-manager-2.6.list && \
    apt-get update && apt-get install scylla-manager-agent

RUN echo -e "auth_token: token\n\
s3:\n\
    access_key_id: minio\n\
    secret_access_key: minio123\n\
    endpoint: http://minio:9000" > /etc/scylla-manager-agent/scylla-manager-agent.yaml \
