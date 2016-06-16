gcloud compute forwarding-rules create http-rule --address 107.178.254.251 \
 --global --target-http-proxy web-proxy --port-range 80

Created [https://www.googleapis.com/compute/v1/projects/cp300demo1/global/forwardingRules/http-rule].
---
IPAddress: 107.178.254.251
IPProtocol: TCP
creationTimestamp: '2016-06-15T02:22:37.443-07:00'
description: ''
id: '8685163442422456498'
kind: compute#forwardingRule
name: http-rule
portRange: 80-80
selfLink: https://www.googleapis.com/compute/v1/projects/cp300demo1/global/forwardingRules/http-rule
target: web-proxy
