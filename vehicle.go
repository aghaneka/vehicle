package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type VEHICLE struct {

}
 


// Application is for storing retreived Application

type Car struct{	
	vinId string `json:"vinId"`
	make string `json:"make"`
	model string `json:"model"`
}

// ListApplication is for storing retreived Application list with status

type ListCars struct{	
	vinId string `json:"vinId"`
	make string `json:"make"`
	model string `json:"model"`
}


// Init initializes the smart contracts

func (t *VEHICLE) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("Cars")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("Cars", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "vinId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "make", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "model", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	
	if err != nil {
		return nil, errors.New("Failed creating Car table.")
	}
	
	// setting up the users role
/*	stub.PutState("user_type1_1", []byte("hdfcSales"))
	stub.PutState("user_type1_2", []byte("healthSales"))
	stub.PutState("user_type1_3", []byte("hdfcUW"))
	stub.PutState("user_type1_4", []byte("healthUW"))	
*/	
	
	
	return nil, nil
}

//============================================Quote=============================================//

func (t *VEHICLE) submitCar(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
        fmt.Println("inserting...."+args[0])
	
		vinId:=args[0]
		make:=args[1]
		model:=args[2]
	
		// Insert a row
		ok, err := stub.InsertRow("Cars", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: vinId}},
				&shim.Column{Value: &shim.Column_String_{String_: make}},
				&shim.Column{Value: &shim.Column_String_{String_: model}},

			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
		
		return nil, nil
}


//get the application(depends on the role)
func (t *VEHICLE) getCar(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
    fmt.Println("get car called NEW...."+args[0])
	vinId := args[0]
	
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: vinId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("Cars", columns)
	if err != nil {
	    fmt.Println("get car erroe")
		jsonResp := "{\"Error\":\"Failed to get the data for the vinId " + vinId + "\"}"
		return []byte("111"), errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
	    fmt.Println("get car empty")
		jsonResp := "{\"Error\":\"Failed to get the data for the vinId " + vinId + "\"}"
		return []byte("222"), errors.New(jsonResp)
	}

	
	
	res2E := Car{}
	
    res2E.vinId = row.Columns[0].GetString_()
	res2E.make = row.Columns[1].GetString_()
	res2E.model = row.Columns[2].GetString_()

	mapB, _ := json.Marshal(res2E)
	fmt.Println("mapB")
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// Invoke invokes the chaincode
func (t *VEHICLE) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "submitCar" {
		t := VEHICLE{}
		return t.submitCar(stub, args)
	}  

	return nil, errors.New("Invalid invoke function name.")

}

func (t *VEHICLE) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getCar" {
		t := VEHICLE{}
		return t.getCar(stub, args)		
	}
	
	return nil, nil
}

func main() {
	err := shim.Start(new(VEHICLE))
	if err != nil {
		fmt.Printf("Error starting VEHICLE: %s", err)
	}
} 
