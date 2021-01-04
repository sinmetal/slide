gcloud compute backend-services create web-service \
 --protocol HTTP --http-health-check basic-check

Created [https://www.googleapis.com/compute/v1/projects/cp300demo1/global/backendServices/web-service].
NAME         BACKENDS  PROTOCOL
web-service            HTTP