package chain

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

// SmartContract used to representate ledger
type SmartContract struct {
	contractapi.Contract
}

// QueryResult used to representate query result data
type QueryResult struct {
	Key    string `json:"key"`
	Record *interface{}
}
