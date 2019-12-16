#!/bin/bash
version=$1

set -ev

#chaincode install on peer0.org1.pretzel.com
docker exec cli peer chaincode install -n pretzel -v $version -p github.com/pretzel
sleep 5
#chaincode instatiate on peer0.org1.pretzel.com
docker exec cli peer chaincode instantiate -n pretzel -v $version -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5


#chaincode install on peer0.org2.pretzel.com
export CORE_PEER_LOCALMSPID=Org2MSP
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org2.pretzel.com/users/Admin@org2.pretzel.com/msp
export CORE_PEER_ADDRESS=peer0.org2.pretzel.com:7051
peer chaincode install -n pretzel -v $version -p github.com/pretzel
sleep 5
#chaincode instatiate on peer0.org2.pretzel.com
peer chaincode instantiate -n pretzel -v $version -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5


#chaincode install on peer0.org3.pretzel.com
export CORE_PEER_LOCALMSPID=Org3MSP
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org3.pretzel.com/users/Admin@org3.pretzel.com/msp
export CORE_PEER_ADDRESS=peer0.org3.pretzel.com:7051
peer chaincode install -n pretzel -v $version -p github.com/pretzel
sleep 5
#chaincode instatiate on peer0.org3.pretzel.com
docker exec cli peer chaincode instantiate -n pretzel -v $version -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5

#cli point to orderer (peer0.org1.pretzel.com)
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.pretzel.com/users/Admin@org1.pretzel.com/msp
export CORE_PEER_ADDRESS=peer0.org1.pretzel.com:7051

echo '-------------------------------------END-------------------------------------'