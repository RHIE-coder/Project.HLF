package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type transferHistory struct {
	fromBank     string `json:"fromBank"`
	fromUsername string `json:"fromUsername"`
	toBank       string `json:"toBank"`
	toUsername   string `json:"toUsername"`
	remittance   int64  `json:"remittance"`
}
type AccountInfo struct {
	username string `json:"username"`
	bank     string `json:bank`
	account  string `json:"account"`
	balance  int64  `json:"balance"`
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

	//initAccount(), transfer(), getTransferHistory(), getAccount()
	switch function {
	case "initAccount1":
		return t.initAccount("initAccount1", stub, args)
	case "initAccount2":
		return t.initAccount("initAccount2", stub, args)
	case "initAccount3":
		return t.initAccount("initAccount3", stub, args)
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
}

func (t *SimpleChaincode) initAccount(funcName string, stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// ==== Input sanitation ====
	fmt.Println("---start " + funcName + "---")
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Private account data must be passed in transient map.")
	}
	transMap, err := stub.GetTransient()
	if err != nil {
		return shim.Error("Error getting transient: " + err.Error())
	}
	if _, ok := transMap["AccountInfo"]; !ok {
		return shim.Error("AccountInfo must be a key in the transient map")
	}
	if len(transMap["AccountInfo"]) == 0 {
		return shim.Error("AccountInfo value in the transient map must be a non-empty JSON string")
	}

	//ready to control JsonObject
	var inputAccount AccountInfo

	err = json.Unmarshal(transMap["AccountInfo"], &inputAccount)
	if err != nil {
		return shim.Error("Failed to decode JSON of: " + string(transMap["AccountInfo"]))
	}

	//validation field
	if len(inputAccount.username) == 0 {
		return shim.Error("username field must be a non-empty string")
	}
	if len(inputAccount.bank) == 0 {
		return shim.Error("bank field must be a non-empty string")
	}
	if inputAccount.balance < 10000 {
		return shim.Error("balance field must be a positive integer and over 10000 won")
	}
	if len(inputAccount.account) == 0 {
		return shim.Error("account field must be a non-empty string")
	}

	// ==== Check if account already exists ====

	var accountAsBytes []byte

	if funcName == "initAccount1" {
		accountAsBytes, err = stub.GetPrivateData("collectionAccountInfoPrivateDetails1", inputAccount.account)
	} else if funcName == "initAccount2" {
		accountAsBytes, err = stub.GetPrivateData("collectionAccountInfoPrivateDetails2", inputAccount.account)
	} else if funcName == "initAccount3" {
		accountAsBytes, err = stub.GetPrivateData("collectionAccountInfoPrivateDetails3", inputAccount.account)
	}

	if err != nil {
		return shim.Error("Failed to get account: " + err.Error())
	} else if accountAsBytes != nil {
		fmt.Println("This account already exists: " + inputAccount.username)
		return shim.Error("This account already exists: " + inputAccount.username)
	}

	// ==== Create object, marshal to JSON, and save to private data ====
	account := &AccountInfo{
		username: inputAccount.username,
		bank:     inputAccount.bank,
		account:  inputAccount.account,
		balance:  inputAccount.balance,
	}
	accountJSONasBytes, err := json.Marshal(account)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("new account name " + inputAccount.username)

	// === Save account to private data ===
	if funcName == "initAccount1" {
		err = stub.PutPrivateData("collectionAccountInfoPrivateDetails1", inputAccount.username, accountJSONasBytes)
	} else if funcName == "initAccount2" {
		err = stub.PutPrivateData("collectionAccountInfoPrivateDetails2", inputAccount.username, accountJSONasBytes)
	} else if funcName == "initAccount3" {
		err = stub.PutPrivateData("collectionAccountInfoPrivateDetails3", inputAccount.username, accountJSONasBytes)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== account saved and indexed. Return success ====
	fmt.Println("---end init account---")
	return shim.Success(nil)

}

func (t *SimpleChaincode) transfer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("start transfer")

	if len(args) != 5 {
		shim.Error("Incorrect arguments number")
	}

	fromBank := args[0]
	fromUsername := args[2]
	toBank := args[3]
	toUsername := args[4]
	remittance, _ := strconv.ParseInt(args[5], 0, 64)

	//setting private data location
	var pdLocFrom string
	var pdLocTo string

	if fromBank == "A" {
		pdLocFrom = "collectionAccountInfoPrivateDetails1"
	} else if fromBank == "B" {
		pdLocFrom = "collectionAccountInfoPrivateDetails2"
	} else if fromBank == "C" {
		pdLocFrom = "collectionAccountInfoPrivateDetails3"
	}

	if toBank == "A" {
		pdLocTo = "collectionAccountInfoPrivateDetails1"
	} else if toBank == "B" {
		pdLocTo = "collectionAccountInfoPrivateDetails2"
	} else if toBank == "C" {
		pdLocTo = "collectionAccountInfoPrivateDetails3"
	}
	fromPdAsBytes, err := stub.GetPrivateData(pdLocFrom, fromUsername)
	toPdAsBytes, err := stub.GetPrivateData(pdLocTo, toUsername)

	fromInfo := AccountInfo{}
	toInfo := AccountInfo{}
	err = json.Unmarshal(fromPdAsBytes, &fromInfo)
	if err != nil {
		shim.Error("fromInfo error")
	}
	err = json.Unmarshal(toPdAsBytes, &toInfo)
	if err != nil {
		shim.Error("toInfo error")
	}

	fromInfo.balance -= remittance
	toInfo.balance += remittance

	fromPdJSONasBytes, _ := json.Marshal(fromInfo)
	toPdJSONasBytes, _ := json.Marshal(toInfo)
	err = stub.PutPrivateData(pdLocFrom, fromInfo.account, fromPdJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutPrivateData(pdLocTo, toInfo.account, toPdJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) getTransferHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//bank 이름 기준
	if len(args[0]) != 1 {
		shim.Error("Incorrect argument number")
	}
	fromBank := args[0]
	historyAsBytes, _ := stub.GetState(fromBank)
	return shim.Success(historyAsBytes)

}
func (t *SimpleChaincode) getAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var username string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments.")
	}

	username = args[0]

	//무슨 기준으로?
	accountAsbytes, err := stub.GetPrivateData("collectionAccountInfoPrivateDetails1", username)
	if err != nil {
		return shim.Error("Error occured : " + err.Error())
	} else if accountAsbytes == nil {
		return shim.Error("no value error : " + err.Error())
	}

	return shim.Success(accountAsbytes)
}
