gcloud compute instances create www3 www4 --zone us-central1-c \
 --tags http-server --metadata startup-script-url=gs://cpo200demo1.appspot.com/startup.sh

Created [https://www.googleapis.com/compute/v1/projects/cp300demo1/zones/us-central1-c/instances/www4].
Created [https://www.googleapis.com/compute/v1/projects/cp300demo1/zones/us-central1-c/instances/www3].
NAME  ZONE           MACHINE_TYPE   PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP    STATUS
www4  us-central1-c  n1-standard-1               10.240.0.6   23.251.145.70  RUNNING
www3  us-central1-c  n1-standard-1               10.240.0.5   146.148.68.4   RUNNING