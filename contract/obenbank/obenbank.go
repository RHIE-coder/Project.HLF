package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type transferHistory struct {
	from		string	`json:"from"`
	to			string	`json:"to"`
	remittance	int64	`json:"remittance"`
}

type AccountInfo struct {
	username	string	`json:"username"`
	account		string	`json:"account"`
	balance		int64	`json:"balance"`
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	//transfer(), getTransferHistory(), getAccount()
	switch function {
	case "initAccount1":
		return t.transfer(stub, args)
	case "initAccount2":
		return t.transfer(stub, args)
	case "initAccount3":
		return t.transfer(stub, args)
	case "transfer":
		return t.transfer(stub, args)
	case "getTransferHistory":
		return t.getTransferHistory(stub, args)
	case "getAccount":
		return t.getAccount(stub, args)
	default:
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")
	}

	func (t *SimpleChaincode) initAccount1(stub shim.ChaincodeStubInterface, args []string) pb.Response{

	}
}