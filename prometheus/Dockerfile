FROM prom/prometheus:latest

# Copy the Prometheus configuration file into the image
COPY prometheus.yml /etc/prometheus/prometheus.yml

# Expose the Prometheus UI port
EXPOSE 9090

# Set the default command to run Prometheus with the necessary flags
ENTRYPOINT [ "prometheus", \
            "--config.file=/etc/prometheus/prometheus.yml", \
            "--storage.tsdb.path=/prometheus", \
            "--storage.tsdb.retention.size=2GB", \
            "--storage.tsdb.retention.time=15d" ]