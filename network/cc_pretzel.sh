#!/bin/bash
version=$1

set -ev

#chaincode install
docker exec cli peer chaincode install -n pretzel -v $version -p github.com/pretzel
#chaincode instatiate
docker exec cli peer chaincode instantiate -n pretzel -v $version -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")'
sleep 5
#chaincode invoke user1
docker exec cli peer chaincode invoke -n pretzel -C pretzelchannel -c '{"Args":["inputExampleData","user1","30"]}'
sleep 5
#chaincode query user1
docker exec cli peer chaincode query -n pretzel -C pretzelchannel -c '{"Args":["readExampleData","user1"]}'

echo '-------------------------------------END-------------------------------------'
