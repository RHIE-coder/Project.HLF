#peer0.org3.pretzel.com
echo "peer0.org1.pretzel.com install chaincode"
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.pretzel.com/users/Admin@org1.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org1.pretzel.com:7051
peer chaincode install -n pretzel -v 1.0 -p github.com/pretzel
sleep 1
peer chaincode install -n pretzel2 -v 1.0 -p github.com/pretzel2
sleep 1
peer chaincode list --installed

#peer0.org3.pretzel.com
echo "peer0.org2.pretzel.com install chaincode"
CORE_PEER_LOCALMSPID=Org2MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org2.pretzel.com/users/Admin@org2.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org2.pretzel.com:7051
peer chaincode install -n pretzel -v 1.0 -p github.com/pretzel
sleep 1
peer chaincode install -n pretzel2 -v 1.0 -p github.com/pretzel2
sleep 1
peer chaincode list --installed


#peer0.org3.pretzel.com
echo "peer0.org3.pretzel.com install chaincode"
CORE_PEER_LOCALMSPID=Org3MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org3.pretzel.com/users/Admin@org3.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org3.pretzel.com:7051
peer chaincode install -n pretzel -v 1.0 -p github.com/pretzel
sleep 1
peer chaincode install -n pretzel2 -v 1.0 -p github.com/pretzel2
sleep 1
peer chaincode list --installed
sleep 5

#chaincode instantiate on pretzelchannel
CORE_PEER_LOCALMSPID=Org1MSP
CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/crypto/peerOrganizations/org1.pretzel.com/users/Admin@org1.pretzel.com/msp
CORE_PEER_ADDRESS=peer0.org1.pretzel.com:7051
echo "instantiating..."
peer chaincode instantiate -n pretzel -v 1.0 -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5
peer chaincode instantiate -n pretzel2 -v 1.0 -C pretzelchannel -c '{"Args":[]}' -P 'OR ("Org1MSP.member", "Org2MSP.member","Org3MSP.member")' --collections-config /opt/gopath/src/github.com/pretzel/collections_config.json
sleep 5
peer chaincode list --instantiated -C pretzelchannel