yes | gcloud compute forwarding-rules delete http-rule --global
yes | gcloud compute target-http-proxies delete web-proxy
yes | gcloud compute url-maps delete web-map
yes | gcloud compute backend-services delete web-service
yes | gcloud compute http-health-checks delete basic-check
yes | gcloud compute firewall-rules delete allow-130-211-0-0-22
yes | gcloud compute firewall-rules delete default-allow-http
yes | gcloud compute instance-groups unmanaged delete "us-group-b" \
 --zone "us-central1-b"
yes | gcloud compute instance-groups unmanaged delete "us-group-c" \
 --zone "us-central1-c"
yes | gcloud compute instances delete www1 www2 --zone us-central1-b
yes | gcloud compute instances delete www3 www4 --zone us-central1-c
yes | gcloud compute addresses delete lb-ip-cr --global