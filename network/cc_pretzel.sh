#!/bin/bash
version=$1

set -ev

#chaincode install
docker exec cli peer chaincode install -n pretzel -v $version -p github.com/pretzel
sleep 5
#chaincode instatiate
docker exec cli peer chaincode instantiate -n pretzel -v $version -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5
echo '-------------------------------------END-------------------------------------'