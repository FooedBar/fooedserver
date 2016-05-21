sudo apt-get update -y
sudo apt-get -y upgrade
sudo curl -O https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
sudo tar -xvf go1.6.linux-amd64.tar.gz
echo export\ PATH=\$PATH:/usr/local/go/bin >> .profile
echo export\ GOPATH=~/gopkg >> .profile
source .profile
sudo apt-get install git -y
go --version
go get github.com/FooedBar/fooedserver
cd gopkg/src/github.com/FooedBar/fooedserver/