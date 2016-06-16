gcloud compute firewall-rules create default-allow-http \
    --source-ranges 0.0.0.0/0 \
    --target-tags http-server \
    --allow tcp:80