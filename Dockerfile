from debian:jessie
maintainer Cory Buecker <email@corybuecker.com>

env GOPATH $HOME/go
env GOROOT /usr/local/go
env PATH $PATH:$GOROOT/bin

run apt-get update
run apt-get install curl git ca-certificates -y --no-install-recommends

workdir /root
run curl -o go1.4.2.linux-amd64.tar.gz http://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz
run echo "5020af94b52b65cc9b6f11d50a67e4bae07b0aff go1.4.2.linux-amd64.tar.gz" | sha1sum --strict --check -

run tar -zxvf go1.4.2.linux-amd64.tar.gz -C /usr/local

run git clone https://github.com/corybuecker/twilio-sms-forward.git $HOME/go/github.com/corybuecker/twilio-sms-forward

workdir /root/go

run go get github.com/corybuecker/twilio-sms-forward/

run go build -o /root/twilio-sms-forward github.com/corybuecker/twilio-sms-forward/

run rm -rf /usr/local/go /root/go
run apt-get purge -y --auto-remove curl git ca-certificates

cmd ["/root/twilio-sms-forward"]
