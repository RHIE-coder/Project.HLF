package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	shim "github.com/hyperledger/fabric/core/chaincode/shim"
	pe "github.com/hyperledger/fabric/protos/peer"
)

//Chaincode for implementation
type pretzelChaincode struct {
}

//exampleData
type exampleData struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
}

type examplePD struct {
	Money int `json:"money"`
}

func main() {
	err := shim.Start(new(pretzelChaincode))
	if err != nil {
		fmt.Printf("Error starting pretzel chaincode: %s", err)
	}
}

func (pc *pretzelChaincode) Init(stub shim.ChaincodeStubInterface) pe.Response {
	doc := `
	{
		"username":"maria",
		"age":10
	}
	`
	return shim.Success([]byte(doc))
}

func (pc *pretzelChaincode) Invoke(stub shim.ChaincodeStubInterface) pe.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running" + function)
	fmt.Println("args : ", args)

	switch function {
	case "inputWS":
		return pc.inputExampleData(stub, args)
	case "inputPD":
		return pc.inputExamplePD(stub, args)
	case "readWS":
		return pc.readExampleData(stub, args)
	case "readPD":
		return pc.readExamplePD(stub, args)
	case "M":
		return pc.checkMultiData(stub, args)
	case "S":
		return pc.checkSingleiData(stub, args)
	default:
		fmt.Println("invoke did not find func:" + function)
		return shim.Error("Received unknown function invocation : " + function)
	}
}

//username, age
func (pc *pretzelChaincode) inputExampleData(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	username := args[0]
	age, _ := strconv.ParseInt(args[1], 0, 0)

	edAsBytes, err := stub.GetState(args[0])
	if edAsBytes == nil {
		fmt.Println("no data. You can input the new data")
	} else {
		return shim.Error("there are already Data")
	}
	if err != nil {
		return shim.Error("inputExampleData() Error")
	}

	edws := exampleData{}
	edws.Username = username
	edws.Age = int(age)
	jsonBytesObj, _ := json.Marshal(edws)

	stub.PutState(username, jsonBytesObj) //error return
	return shim.Success(jsonBytesObj)
}

//money
func (pc *pretzelChaincode) inputExamplePD(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	username := args[0]
	money, _ := strconv.ParseInt(args[1], 0, 0)
	pdName := args[2]
	epdAsBytes, err := stub.GetPrivateData(pdName, username)
	if epdAsBytes == nil {
		fmt.Println("no data. You can input the new data")
	} else {
		return shim.Error("there are already Data")
	}
	if err != nil {
		return shim.Error("inputExamplePD() Error")
	}
	epd := examplePD{}
	epd.Money = int(money)
	jsonBytesObj, _ := json.Marshal(epd)
	stub.PutPrivateData(pdName, username, jsonBytesObj) //error return
	return shim.Success(jsonBytesObj)
}

func (pc *pretzelChaincode) readExampleData(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	data, err := stub.GetState(args[0])
	if err != nil {
		shim.Error("readExampleData() Error")
	}
	return shim.Success(data)
}
func (pc *pretzelChaincode) readExamplePD(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	pdName := args[1]
	data, err := stub.GetPrivateData(pdName, args[0])
	if err != nil {
		shim.Error("readExamplePD() Error")
	}
	return shim.Success(data)
}

func (pc *pretzelChaincode) checkMultiData(stub shim.ChaincodeStubInterface, args []string) pe.Response {
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

func (pc *pretzelChaincode) checkSingleiData(stub shim.ChaincodeStubInterface, args []string) pe.Response {
	a := args[0]
	type checker struct {
		A string `json:"a"`
	}
	ch := checker{}
	ch.A = a
	obj, _ := json.Marshal(ch)
	return shim.Success([]byte(obj))
}
