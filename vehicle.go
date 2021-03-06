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
	VinId string `json:"vinId"`
	Make string `json:"make"`
	Model string `json:"model"`
}

// ListApplication is for storing retreived Application list with status

type ListCars struct{	
	VinId string `json:"vinId"`
	Make string `json:"make"`
	Model string `json:"model"`
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

	fmt.Println("no error")

	
	res2E := Car{}
	
    res2E.VinId = row.Columns[0].GetString_()
	res2E.Make = row.Columns[1].GetString_()
	res2E.Model = row.Columns[2].GetString_()

    
    fmt.Println("row 1"+row.Columns[0].GetString_())
    fmt.Println("row 2"+row.Columns[1].GetString_())
    fmt.Println("row 3"+row.Columns[2].GetString_())

	json_byte, err := json.Marshal(res2E)

		if err != nil {

			return nil, err
			panic(err)
		}
		return json_byte, nil

}

func (t *VEHICLE) getAllCars(stub shim.ChaincodeStubInterface) ([]byte, error) {
    	
	var columns []shim.Column
	
	rows, err := stub.GetRows("Cars", columns)
	
	
	res2E := []*Car{}

		for {
			select {

			case row, ok := <-rows:

				if !ok {
					rows = nil
				} else {

					u := new(Car)
					u.VinId = row.Columns[0].GetString_()
					u.Make = row.Columns[1].GetString_()
					u.Model = row.Columns[2].GetString_()
					res2E = append(res2E, u)
				}
			}
			if rows == nil {
				break
			}
		}

		jsonRows, err := json.Marshal(res2E)

		if err != nil {
			return nil, fmt.Errorf("getcars operation failed. Error marshaling JSON: %s", err)
		}

		return jsonRows, nil	
	
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
	if function == "getAllCars" {
		t := VEHICLE{}
		return t.getAllCars(stub)		
	}	
	
	return nil, nil
}

func main() {
	err := shim.Start(new(VEHICLE))
	if err != nil {
		fmt.Printf("Error starting VEHICLE: %s", err)
	}
} 
