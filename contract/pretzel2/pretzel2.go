package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pe "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for implementation
type pretzelChaincode2 struct {
}

//exampleData2
type exampleData2 struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type examplePD2 struct {
	Money int `json:"money"`
}

func main() {
	err := shim.Start(new(pretzelChaincode2))
	if err != nil {
		fmt.Printf("Error starting pretzel chaincode: %s", err)
	}
}

func (pc *pretzelChaincode2) Init(stub shim.ChaincodeStubInterface) pe.Response {
	return shim.Success(nil)
}

func (pc *pretzelChaincode2) Invoke(stub shim.ChaincodeStubInterface) pe.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running" + function)
	fmt.Println("args : ", args)

	switch function {
	case "inputWS":
		return pc.inputexampleData2(stub, args)
	case "inputPD":
		return pc.inputexamplePD2(stub, args)
	case "readWS":
		return pc.readexampleData2(stub, args)
	case "readPD":
		return pc.readexamplePD2(stub, args)
	case "M":
		return pc.checkMultiData2(stub, args)
	case "S":
		return pc.checkSingleiData2(stub, args)
	default:
		fmt.Println("invoke did not find func:" + function)
		return shim.Error("Received unknown function invocation : " + function)
	}
}

//username, age
func (pc *pretzelChaincode2) inputexampleData2(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	username := args[0]
	age, _ := strconv.ParseInt(args[1], 0, 0)

	edAsBytes, err := stub.GetState(args[0])
	if edAsBytes == nil {
		fmt.Println("no data. You can input the new data")
	} else {
		return shim.Error("there are already Data")
	}
	if err != nil {
		return shim.Error("inputexampleData2() Error")
	}

	edws := exampleData2{}
	edws.Username = username
	edws.Age = int(age)
	jsonBytesObj, _ := json.Marshal(edws)

	stub.PutState(username, jsonBytesObj) //error return
	return shim.Success(jsonBytesObj)
}

//money
func (pc *pretzelChaincode2) inputexamplePD2(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	var err error
	pdName := args[0]
	transMap, err := stub.GetTransient()
	type pdInput struct {
		Username string `json:"username"`
		Money    int    `json:"money"`
	}
	if err != nil {
		return shim.Error("inputexamplePD2() Error")
	}
	pdin := pdInput{}
	err = json.Unmarshal(transMap["data"], &pdin)
	// epdAsBytes, err := stub.GetPrivateData(pdName, username)
	// if epdAsBytes == nil {
	// 	fmt.Println("no data. You can input the new data")
	// } else {
	// 	return shim.Error("there are already Data")
	// }
	// if err != nil {
	// 	return shim.Error("inputexamplePD2() Error")
	// }
	epd := examplePD2{}
	username := pdin.Username
	epd.Money = pdin.Money
	jsonBytesObj, _ := json.Marshal(epd)
	stub.PutPrivateData(pdName, username, jsonBytesObj) //error return

	return shim.Success(jsonBytesObj)
}

func (pc *pretzelChaincode2) readexampleData2(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	data, err := stub.GetState(args[0])
	if err != nil {
		shim.Error("readexampleData2() Error")
	}
	return shim.Success(data)
}
func (pc *pretzelChaincode2) readexamplePD2(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	pdName := args[1]
	data, err := stub.GetPrivateData(pdName, args[0])
	if err != nil {
		shim.Error("readexamplePD2() Error")
	}
	return shim.Success(data)
}

func (pc *pretzelChaincode2) checkMultiData2(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	a := args[0]
	b := args[1]
	type checker struct {
		A string `json:"a"`
		B string `json:"b"`
	}
	ch := checker{}
	ch.A = a
	ch.B = b
	obj, _ := json.Marshal(ch)
	return shim.Success([]byte(obj))
}

func (pc *pretzelChaincode2) checkSingleiData2(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	a := args[0]
	type checker struct {
		A string `json:"a"`
	}
	ch := checker{}
	ch.A = a
	obj, _ := json.Marshal(ch)
	return shim.Success([]byte(obj))
}
