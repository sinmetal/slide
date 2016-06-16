gcloud compute backend-services add-backend web-service \
    --balancing-mode UTILIZATION \
    --max-utilization 0.8 \
    --capacity-scaler 1 \
    --instance-group us-group-c \
    --instance-group-zone us-central1-c

Updated [https://www.googleapis.com/compute/v1/projects/cp300demo1/global/backendServices/web-service].