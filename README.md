# twilio-sms-forward
A Go-based web application to forward SMS messages from a Twilio number.

To compile statically for the Docker container.

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v --ldflags '-extldflags -static -w -s' -installsuffix nocgo .

[![Circle CI](https://circleci.com/gh/corybuecker/twilio-sms-forward.svg?style=svg)](https://circleci.com/gh/corybuecker/twilio-sms-forward)
