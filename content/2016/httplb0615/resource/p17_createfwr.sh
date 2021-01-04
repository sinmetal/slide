gcloud compute firewall-rules create allow-130-211-0-0-22 \
    --source-ranges 130.211.0.0/22 \
    --target-tags http-server \
    --allow tcp:80