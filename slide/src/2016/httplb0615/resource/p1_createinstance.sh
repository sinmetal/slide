gcloud compute instances create www1 www2 --zone us-central1-b \
 --tags http-server --metadata startup-script-url=gs://cpo200demo1.appspot.com/startup.sh

Created [https://www.googleapis.com/compute/v1/projects/cp300demo1/zones/us-central1-b/instances/www2].
Created [https://www.googleapis.com/compute/v1/projects/cp300demo1/zones/us-central1-b/instances/www1].
NAME  ZONE           MACHINE_TYPE   PREEMPTIBLE  INTERNAL_IP  EXTERNAL_IP      STATUS
www2  us-central1-b  n1-standard-1               10.240.0.2   104.197.253.46   RUNNING
www1  us-central1-b  n1-standard-1               10.240.0.4   104.154.101.154  RUNNING
