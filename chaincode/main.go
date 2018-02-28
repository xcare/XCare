// main
package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type XChaincode struct {
}

func (t *XChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return t.init(stub)
}

func (t *XChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "userApprove" { 
		return t.userApprove(stub, args)
	} else if function == "bid" { 
		return t.bid(stub, args)
	} else if function == "bidApprove" { 
		//		return t.bid(stub, args)
	} else if function == "queryBidsByTender" { 
		return t.queryBidsByTender(stub, args)
	}

	return shim.Error("Invalid invoke function name.")
}

func main() {
	err := shim.Start(new(XChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

