#!/bin/bash
export o1p1=0
export o1p2=1

export o2p1=0
export o2p2=1

export o3p1=0
export o3p2=1

#peer0.org1.pretzel.com
echo ">>>>>>>>>>>peer0.org1.pretzel.com install chaincode"
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.pretzel.com/users/Admin@org1.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org1.pretzel.com:7051
if [ $o1p1 -eq 1 ]
then
echo ">>>>>>>>>>>peer0.org1.pretzel.com.o1p1"
peer chaincode install -n pretzel -v 1.0 -p github.com/pretzel
sleep 1
fi
if [ $o1p2 -eq 1 ];then
echo ">>>>>>>>>>>peer0.org1.pretzel.com.o1p2"
peer chaincode install -n pretzel2 -v 1.0 -p github.com/pretzel2
sleep 1
fi
peer chaincode list --installed

#peer0.org2.pretzel.com
echo ">>>>>>>>>>>peer0.org2.pretzel.com install chaincode"
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org2.pretzel.com/users/Admin@org2.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org2.pretzel.com:8051
if [ $o2p1 -eq 1 ];then
echo ">>>>>>>>>>>peer0.org2.pretzel.com.o2p1"
peer chaincode install -n pretzel -v 1.0 -p github.com/pretzel
sleep 1
fi
if [ $o2p2 -eq 1 ];then
echo ">>>>>>>>>>>peer0.org2.pretzel.com.o2p2"
peer chaincode install -n pretzel2 -v 1.0 -p github.com/pretzel2
sleep 1
fi
peer chaincode list --installed


#peer0.org3.pretzel.com
echo ">>>>>>>>>>>peer0.org3.pretzel.com install chaincode"
CORE_PEER_LOCALMSPID=Org3MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org3.pretzel.com/users/Admin@org3.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org3.pretzel.com:9051
if [ $o3p1 -eq 1 ];then
echo ">>>>>>>>>>>peer0.org3.pretzel.com.o3p1"
peer chaincode install -n pretzel -v 1.0 -p github.com/pretzel
sleep 1
fi
if [ $o3p2 -eq 1 ];then
echo ">>>>>>>>>>>peer0.org3.pretzel.com.o3p2"
peer chaincode install -n pretzel2 -v 1.0 -p github.com/pretzel2
sleep 1
fi
peer chaincode list --installed
sleep 5

#chaincode instantiate on pretzelchannel
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.pretzel.com/users/Admin@org1.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org1.pretzel.com:7051
echo "instantiating..."
# peer chaincode instantiate -n pretzel -v 1.0 -C pretzelchannel -c '{"Args":[]}'
# peer chaincode instantiate -n pretzel -v 1.0 -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")'
peer chaincode instantiate -n pretzel -v 1.0 -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5
# peer chaincode instantiate -n pretzel2 -v 1.0 -C pretzelchannel -c '{"Args":[]}'
# peer chaincode instantiate -n pretzel2 -v 1.0 -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")'
peer chaincode instantiate -n pretzel2 -v 1.0 -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel2/collections_config.json
sleep 5
peer chaincode list --instantiated -C pretzelchannel
#"OR('Org1MSP.member','Org2MSP.member','Org3MSP.member')"
echo "<<<<<<<<<<<<<<<<<<  org1 - invoke, query >>>>>>>>>>>>>>>>>>>>"
set -x
peer chaincode invoke -n pretzel2 -C pretzelchannel -c '{"Args":["inputWS","user1","10"]}'
peer chaincode query -n pretzel2 -C pretzelchannel -c '{"Args":["readWS","user1"]}'
sleep 3
set +x
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org2.pretzel.com/users/Admin@org2.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org2.pretzel.com:8051
echo "<<<<<<<<<<<<<<<<<<  org2 - invoke, query >>>>>>>>>>>>>>>>>>>>"
set -x
peer chaincode invoke -n pretzel2 -C pretzelchannel -c '{"Args":["inputWS","user2","20"]}'
peer chaincode query -n pretzel2 -C pretzelchannel -c '{"Args":["readWS","user2"]}'
sleep 3
set +x
CORE_PEER_LOCALMSPID=Org3MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org3.pretzel.com/users/Admin@org3.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org3.pretzel.com:9051
echo "<<<<<<<<<<<<<<<<<<  org2 - invoke, query >>>>>>>>>>>>>>>>>>>>"
set -x
peer chaincode invoke -n pretzel2 -C pretzelchannel -c '{"Args":["inputWS","user3","30"]}'
peer chaincode query -n pretzel2 -C pretzelchannel -c '{"Args":["readWS","user3"]}'
sleep 3
set +x
set -e
peer chaincode query -n pretzel2 -C pretzelchannel -c '{"Args":["readWS","user1"]}'
peer chaincode query -n pretzel2 -C pretzelchannel -c '{"Args":["readWS","user2"]}'
peer chaincode query -n pretzel2 -C pretzelchannel -c '{"Args":["readWS","user3"]}'
set +e