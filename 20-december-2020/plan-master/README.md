# Plancoin v2

Plancoin is a blockchain built on the Cosmos SDK with ability to generate a passive income (posmining) and create your own coins!

## Requirements

 - At least 50 GB disk space
 - Golang >= 1.13 with the $GOPATH and $GOBIN variables set
 - *Nix system
 
## Install

mkdir go_lang
cd go_lang
wget https://dl.google.com/go/go1.15.linux-amd64.tar.gz
sha256sum go1.15.linux-amd64.tar.gz
sha256sum go1.15.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.linux-amd64.tar.gz
go -version
nano -c ~/.profile
source ~/.profile
go -version
go version
cd ~
ls -l
pwd
mkdir -p go/src
cd go/src/31-october-2020/plan-master/
go get ./...
mkdir /usr/local/plancoin
cd cmd/plancoind/
ls -l
go build
cp plancoind /usr/local/plancoin/
cd ../plancoincli/
go build
cp plancoincli /usr/local/plancoin/
cd ~
plancoind -help
plancoind --help
plancoincli --help
ls -l /usr/local/plancoin/plancoin-configure
ls -l /usr/local/plancoin/
ls -l /usr/local/plancoin/plancoincli-configure
chmod u+w /usr/local/plancoin/plancoincli-configure
nano -c /usr/local/plancoin/plancoincli-configure
chmod u-w /usr/local/plancoin/plancoincli-configure
plancoincli-configure



## Running the node

Just start it by `plancoind start` and you should see it synchronizing the blocks. 

If you want to become a validator, please check the docs <https://>

## Resources

* Website: <https://plan-coin.com/>
* Explorer: <https://>
* Docs: <https://>
* Telegram channel: <https://t.me/plancoin>
