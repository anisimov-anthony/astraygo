#!/bin/sh
set -e

# Замена переменных окружения в конфиге
sed "s/\${REMOTE_APP_SERVER_HOST}/${REMOTE_APP_SERVER_HOST}/g" \
  /etc/prometheus/prometheus.yml.template > /etc/prometheus/prometheus.yml

echo "Prometheus configuration generated:"
cat /etc/prometheus/prometheus.yml

# Запуск Prometheus
exec /bin/prometheus \
  --config.file=/etc/prometheus/prometheus.yml \
  --storage.tsdb.path=/prometheus \
  --web.console.libraries=/usr/share/prometheus/console_libraries \
  --web.console.templates=/usr/share/prometheus/consoles \
  --web.enable-lifecycle
