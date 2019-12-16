#!/bin/bash
echo "================restart================"
./teardown.sh
echo
echo
echo "finish teardown"
sleep 1
docker network rm net_basic
echo
echo
echo "finish network remove"
sleep 1
./start.sh
./cc_pret.sh
echo
echo
echo "ready to work..."
echo "================script has finished================"