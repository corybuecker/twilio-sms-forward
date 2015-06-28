from busybox:latest
maintainer Cory Buecker <email@corybuecker.com>

add twilio-sms-forward /usr/bin/twilio-sms-forward
entrypoint ["/usr/bin/twilio-sms-forward"]
