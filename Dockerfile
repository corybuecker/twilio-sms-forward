from busybox
maintainer Cory Buecker <email@corybuecker.com>

run mkdir /go
add twilio-sms-forward /go/twilio-sms-forward
cmd ["/go/twilio-sms-forward"]
