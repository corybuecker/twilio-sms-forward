from golang
maintainer Cory Buecker <email@corybuecker.com>

run go get github.com/corybuecker/twilio-sms-forward
run go install github.com/corybuecker/twilio-sms-forward

entrypoint /go/bin/twilio-sms-forward
