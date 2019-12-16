#!/bin/bash
set -ev
#chaincode invoke user1
docker exec cli peer chaincode invoke -n pretzel -C pretzelchannel -c '{"Args":["inputWS","user1","30"]}'
sleep 5
#chaincode query user1
docker exec cli peer chaincode query -n pretzel -C pretzelchannel -c '{"Args":["readWS","user1"]}'
sleep 5
echo '-------------------------------------END-------------------------------------'
