/*
 * Copyright tongdun technology All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Define the BlacklistContract structure
type demoContract struct {
}

// Define the Blacklist structure, with 3 properties.
type demo struct {
	Index     int         `json:"Index"`
	Timestamp interface{} `json:"Timestamp"`
	BLHash    string      `json:"BLHash"`
}

/*
 * The Init method is called when the Smart Contract "fabblack" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (B *demoContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (B *demoContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()

	// Route to the appropriate handler function to interact with the ledger appropriately
	switch function {
	case "querydemo":
		return B.querydemo(stub, args)
	case "Createdemo":
		return B.Createdemo(stub, args)
	default:
		return shim.Error("no such function!")
	}
}

//query blacklist by key(That is "Index")
func (B *demoContract) querydemo(stub shim.ChaincodeStubInterface, Args []string) pb.Response {
	if len(Args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	demo, _ := stub.GetState(Args[0])
	return shim.Success(demo)
}

func (B *demoContract) Createdemo(stub shim.ChaincodeStubInterface, Args []string) pb.Response {
	if len(Args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	time, err := stub.GetTxTimestamp()
	if err != nil {
		return shim.Error(err.Error())
	}
	index, _ := strconv.Atoi(Args[0])
	var demo = demo{Index: index, Timestamp: time, BLHash: Args[1]}
	demoAsBytes, _ := json.Marshal(demo)
	stub.PutState(Args[0], demoAsBytes)

	return shim.Success(nil)
}

func (B *demoContract) GetDataFromIPFS(stub shim.ChaincodeStubInterface, Args []string) pb.Response {
	if len(Args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

}

func main() {
	err := shim.Start(new(demoContract))
	if err != nil {
		fmt.Println("Error create new Smart Contract: %s", err)
	}
}
