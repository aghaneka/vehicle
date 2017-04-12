package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// HDFC is a high level smart contract that HDFCs together business artifact based smart contracts
type HDFC struct {

}



// Application is for storing retreived Application

type Application struct{	
	ApplicationId string `json:"applicationId"`
	Status string `json:"status"`
	Title string `json:"title"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Gender string `json:"gender"`
	Dob string `json:"dob"`
	Age string `json:"age"`
	MartialStatus string `json:"martialStatus"`
	FatherName string `json:"fatherName"`
	MotherName string `json:"motherName"`
	Nationality string `json:"nationality"`
	ResidentialStatus string `json:"residentialStatus"`
	PlaceOfBirth string `json:"placeOfBirth"`
	PanNumber string `json:"panNumber"`
	AadharNumber string `json:"aadharNumber"`
	EducationalQualification string `json:"educationalQualification"`
	PoliticallyExposed string `json:"politicallyExposed"`
	DisablePersonPolicy string `json:"disablePersonPolicy"`
	AnyCriminalProceeding string `json:"anyCriminalProceeding"`
	LifeApprovalStatus string `json:"lifeApprovalStatus"`
	HealthApprovalStatus string `json:"healthApprovalStatus"`
	LifePlanId string `json:"lifePlanId"`
	SumAssuredLife string `json:"sumAssuredLife"`
	HealthPlanId string `json:"healthPlanId"`
	SumAssuredHealth string `json:"sumAssuredHealth"`
	HealthPremium string `json:"healthPremium"`
	LifePremium string `json:"lifePremium"`
	LifeRiderSumAssured string `json:"lifeRiderSumAssured"`
	LifeRiderTerm string `json:"lifeRiderTerm"`
	LifeRiderPremTerm string `json:"lifeRiderPremTerm"`
	LifeRiderPaymentOption string `json:"lifeRiderPaymentOption"`
	TermLength string `json:"termLength"`
	PremiumTerm string `json:"premiumTerm"`
	PremiumFrequency string `json:"premiumFrequency"`
	PremiumPaymentOption string `json:"premiumPaymentOption"`
	AppCreatedBy string `json:"appCreatedBy"`
}

	
// Quote is for storing retreived Quote

type Quote struct{	
	QuoteId string `json:"quoteId"`
	Dob string `json:"dob"`
	Gender string `json:"gender"`
	SumAssuredLife string `json:"sumAssuredLife"`
	LifeRiderList string `json:"lifeRiderList"`
	SumAssuredHealth string `json:"sumAssuredHealth"`
	HealthRiderList string `json:"healthRiderList"`
	TermLength string `json:"termLength"`
	PremiumTerm string `json:"premiumTerm"`
	PremiumFrequency string `json:"premiumFrequency"`
	PremiumPaymentOption string `json:"premiumPaymentOption"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	EmailId string `json:"emailId"`
	Mobile string `json:"mobile"`
	Nationality string `json:"nationality"`
	ResidentialStatus string `json:"residentialStatus"`
	Country string `json:"country"`
	State string `json:"state"`
	LifePlanId string `json:"lifePlanId"`
	HealthPlanId string `json:"healthPlanId"`
	Age string `json:"age"`
	Status string `json:"status"`
	HealthPremium string `json:"healthPremium"`
	LifePremium string `json:"lifePremium"`
	LifeRiderSumAssured string `json:"lifeRiderSumAssured"`
	LifeRiderTerm string `json:"lifeRiderTerm"`
	LifeRiderPremTerm string `json:"lifeRiderPremTerm"`
	LifeRiderPaymentOption string `json:"lifeRiderPaymentOption"`
	CreatedBy string `json:"createdBy"`
	HealthSI string `json:"healthSI"`
	LifeSI string `json:"lifeSI"`
	Tenure string `json:"tenure"`

}


// ListApplication is for storing retreived Application list with status

type ListApplication struct{	
	ApplicationId string `json:"applicationId"`
	Status string `json:"status"`
	LifeApprovalStatus string `json:"lifeApprovalStatus"`
	HealthApprovalStatus string `json:"healthApprovalStatus"`
}

// ListQuote is for storing retreived Quote list with status

type ListQuote struct{	
	QuoteId string `json:"quoteId"`
	Status string `json:"status"`
}

// CountApplication is for storing retreived Application count

type CountApplication struct{	
	Count int `json:"count"`
}

// Premium is for storing retreived Premium count

type Premium struct{	
	HealthPremium string `json:"healthPremium"`
	LifePremium string `json:"lifePremium"`	
}



// Init initializes the smart contracts

func (t *HDFC) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("ApplicationTable")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("ApplicationTable", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "applicationId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "title", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "age", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "martialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "fatherName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "motherName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "residentialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "placeOfBirth", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "panNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "aadharNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "educationalQualification", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "politicallyExposed", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "disablePersonPolicy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "anyCriminalProceeding", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeApprovalStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthApprovalStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifePlanId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sumAssuredLife", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthPlanId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sumAssuredHealth", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthPremium", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifePremium", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderSumAssured", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderTerm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderPremTerm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderPaymentOption", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "termLength", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "premiumTerm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "premiumFrequency", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "premiumPaymentOption", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "appCreatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	
	if err != nil {
		return nil, errors.New("Failed creating ApplicationTable.")
	}
	
	// Check if table already exists
	_, err = stub.GetTable("Quote")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	
	// Create Quote Table
	err = stub.CreateTable("Quote", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "quoteId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sumAssuredLife", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderList", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "sumAssuredHealth", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthRiderList", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "termLength", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "premiumTerm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "premiumFrequency", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "premiumPaymentOption", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "emailId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "mobile", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "nationality", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "residentialStatus", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "country", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "state", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifePlanId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthPlanId", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "age", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthPremium", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifePremium", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderSumAssured", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderTerm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderPremTerm", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeRiderPaymentOption", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "createdBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "healthSI", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "lifeSI", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "tenure", Type: shim.ColumnDefinition_STRING, Key: false},		
	})
	if err != nil {
		return nil, errors.New("Failed creating Quote.")
	}
	
	
	// setting up the users role
	stub.PutState("user_type1_1", []byte("hdfcSales"))
	stub.PutState("user_type1_2", []byte("healthSales"))
	stub.PutState("user_type1_3", []byte("hdfcUW"))
	stub.PutState("user_type1_4", []byte("healthUW"))	
	
	
	
	return nil, nil
}

//============================================Quote=============================================//

//submit the quote, store and calculate the premium and return the premium
func (t *HDFC) submitQuote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 31 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 31. Got: %d.", len(args))
		}

		
		quoteId:=args[0]
		dob:=args[1]
		gender:=args[2]
		sumAssuredLife:=args[3]
		lifeRiderList:=args[4]
		sumAssuredHealth:=args[5]
		healthRiderList:=args[6]
		termLength:=args[7]
		premiumTerm:=args[8]
		premiumFrequency:=args[9]
		premiumPaymentOption:=args[10]
		firstName:=args[11]
		lastName:=args[12]
		emailId:=args[13]
		mobile:=args[14]
		nationality:=args[15]
		residentialStatus:=args[16]
		country:=args[17]
		state:=args[18]
		lifePlanId:=args[19]
		healthPlanId:=args[20]
		age:=args[21]
		status:=args[22]
		lifeRiderSumAssured:=args[23]
		lifeRiderTerm:=args[24]
		lifeRiderPremTerm:=args[25]
		lifeRiderPaymentOption:=args[26]		
		healthSI:=args[27]
		lifeSI:=args[28]
		tenure:=args[29]
		assignerRole:=args[30]
		
		
		//Dummy Logic for Premium calculation
		
		newTenure, _ := strconv.ParseInt(tenure, 10, 0)
		newHealthSI, _ := strconv.ParseInt(healthSI, 10, 0)
		newLifeSI, _ := strconv.ParseInt(lifeSI, 10, 0)
		
		newHealthPremium := int(int(newTenure) * (int(newHealthSI) / 50000) * 50)
		newLifePremium := int((int(newLifeSI) / 100000) * int(newTenure) * 26)
			
		healthPremium :=strconv.Itoa(newHealthPremium)
		lifePremium := strconv.Itoa(newLifePremium)
		
		
		
		// The metadata will contain the role of the users 
		//assignerRole, err := stub.ReadCertAttribute("userid")
		//fmt.Printf("Assiger role is %v\n", assignerRole)

		//if err != nil {
		//	return nil, fmt.Errorf("Failed getting metadata, [%v]", err)
		//}

		//if len(assignerRole) == 0 {
		//	return nil, errors.New("Invalid assigner role. Empty.")
		//}

		
		createdBy:=assignerRole
		
		// Insert a row
		ok, err := stub.InsertRow("Quote", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: quoteId}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredLife}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderList}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredHealth}},
				&shim.Column{Value: &shim.Column_String_{String_: healthRiderList}},
				&shim.Column{Value: &shim.Column_String_{String_: termLength}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumFrequency}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: emailId}},
				&shim.Column{Value: &shim.Column_String_{String_: mobile}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: country}},
				&shim.Column{Value: &shim.Column_String_{String_: state}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderSumAssured}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPremTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: createdBy}},
				&shim.Column{Value: &shim.Column_String_{String_: healthSI}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeSI}},
				&shim.Column{Value: &shim.Column_String_{String_: tenure}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		
		res2E := Premium{}
		res2E.HealthPremium = healthPremium
		res2E.LifePremium = lifePremium
		mapB, _ := json.Marshal(res2E)
		fmt.Println(string(mapB))
		
		return mapB, nil

}


//get the quote (depends on the role,only sales guy can see)

func (t *HDFC) getQuote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting quoteId to query")
	}

	quoteId := args[0]
	assignerRole := args[1]
	
	
	// Get the row pertaining to this quoteId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: quoteId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("Quote", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the quote " + quoteId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the quoteId " + quoteId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := Quote{}
	
	
	res2E.QuoteId = row.Columns[0].GetString_()
	res2E.Dob = row.Columns[1].GetString_()
	res2E.Gender = row.Columns[2].GetString_()
	res2E.SumAssuredLife = row.Columns[3].GetString_()
	res2E.LifeRiderList = row.Columns[4].GetString_()
	res2E.SumAssuredHealth = row.Columns[5].GetString_()
	res2E.HealthRiderList = row.Columns[6].GetString_()
	res2E.TermLength = row.Columns[7].GetString_()
	res2E.PremiumTerm = row.Columns[8].GetString_()
	res2E.PremiumFrequency = row.Columns[9].GetString_()
	res2E.PremiumPaymentOption = row.Columns[10].GetString_()
	res2E.FirstName = row.Columns[11].GetString_()
	res2E.LastName = row.Columns[12].GetString_()
	res2E.EmailId = row.Columns[13].GetString_()
	res2E.Mobile = row.Columns[14].GetString_()
	res2E.Nationality = row.Columns[15].GetString_()
	res2E.ResidentialStatus = row.Columns[16].GetString_()
	res2E.Country = row.Columns[17].GetString_()
	res2E.State = row.Columns[18].GetString_()
	res2E.LifePlanId = row.Columns[19].GetString_()
	res2E.HealthPlanId = row.Columns[20].GetString_()
	res2E.Age = row.Columns[21].GetString_()
	res2E.Status = row.Columns[22].GetString_()
	res2E.HealthPremium = row.Columns[23].GetString_()
	res2E.LifePremium = row.Columns[24].GetString_()
	res2E.LifeRiderSumAssured = row.Columns[25].GetString_()
	res2E.LifeRiderTerm = row.Columns[26].GetString_()
	res2E.LifeRiderPremTerm = row.Columns[27].GetString_()
	res2E.LifeRiderPaymentOption = row.Columns[28].GetString_()
	res2E.CreatedBy = row.Columns[29].GetString_()
	res2E.HealthSI = row.Columns[30].GetString_()
	res2E.LifeSI = row.Columns[31].GetString_()
	res2E.Tenure = row.Columns[32].GetString_()

    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	
	// restrict Under Writer to view the quote
	
	//getting the role of the user
	//assignerRole, err := stub.ReadCertAttribute("userid")
	//fmt.Printf("Assiger role is %v\n", string(assignerRole))

	//if err != nil {
	//	return nil, fmt.Errorf("Failed getting metadata, [%v]", err)
	//}

	//if len(assignerRole) == 0 {
	//	return nil, errors.New("Invalid assigner role. Empty.")
	//}
	
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	quoteCreateOrg1, err := stub.GetState(res2E.CreatedBy)
	quoteCreateOrg := string (quoteCreateOrg1)
	
	if assignerOrg=="hdfcUW" && quoteCreateOrg=="healthSales"{
		return nil, fmt.Errorf("You are not authorized")
	}else if assignerOrg=="healthUW" && quoteCreateOrg=="hdfcSales"{
		return nil, fmt.Errorf("You are not authorized")
	}
	
	return mapB, nil

}

//all quote id with overall status(irrespective of the role)
func (t *HDFC) listAllQuote(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	var columns []shim.Column

	rows, err := stub.GetRows("Quote", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
		
	res2E:= []*ListQuote{}	
	
	for row := range rows {		
		newApp:= new(ListQuote)
		newApp.QuoteId = row.Columns[0].GetString_()
		newApp.Status = row.Columns[22].GetString_()
		res2E=append(res2E,newApp)				
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}


//===================================application============================================//

func (t *HDFC) submitApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 34 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 34. Got: %d.", len(args))
		}

		applicationId := args[0]		
		title := args[1]
		firstName := args[2]
		lastName := args[3]
		gender := args[4]
		dob := args[5]
		age := args[6]
		martialStatus := args[7]
		fatherName := args[8]
		motherName := args[9]
		nationality := args[10]
		residentialStatus := args[11]
		placeOfBirth := args[12]
		panNumber := args[13]
		aadharNumber := args[14]
		educationalQualification := args[15]
		politicallyExposed := args[16]
		disablePersonPolicy := args[17]
		anyCriminalProceeding := args[18]		
		lifePlanId := args[19]
		sumAssuredLife := args[20]
		healthPlanId := args[21]
		sumAssuredHealth := args[22]
		healthPremium := args[23]
		lifePremium := args[24]
		lifeRiderSumAssured := args[25]
		lifeRiderTerm := args[26]
		lifeRiderPremTerm := args[27]
		lifeRiderPaymentOption := args[28]
		termLength := args[29]
		premiumTerm := args[30]
		premiumFrequency := args[31]
		premiumPaymentOption := args[32]
		assignerRole := args[33]
		
		status := "SUBMITTED"
		lifeApprovalStatus := "AUTO"
		healthApprovalStatus := "AUTO"
		
			
		//Setting up the status for the application
		
		userAge, _ := strconv.ParseInt(age, 10, 0)
		
		if userAge >=40 && userAge<=60{
			status = "MEDICAL_REQUIRED"
			lifeApprovalStatus = "AUTO"
			healthApprovalStatus = "AUTO"
		} else if userAge >=30 && userAge<40{
			status = "MANUAL_UW"
			lifeApprovalStatus = "UW_PENDING"
			healthApprovalStatus = "UW_PENDING"
		}else if userAge >=25 && userAge<30{
			status = "APPROVED"
			lifeApprovalStatus = "AUTO"
			healthApprovalStatus = "AUTO"
		} else{
			status = "REJECTED"
			lifeApprovalStatus = "AUTO"
			healthApprovalStatus = "AUTO"
		}
		
		
		//getting the user role
		//assignerRole, err := stub.ReadCertAttribute("userid")
		//fmt.Printf("Assiger role is %v\n", string(assignerRole))

		//if err != nil {
		//	return nil, fmt.Errorf("Failed getting metadata, [%v]", err)
		//}

		//if len(assignerRole) == 0 {
		//	return nil, errors.New("Invalid assigner role. Empty.")
		//}
				
		appCreatedBy := assignerRole
		
		// Insert a row
		ok, err := stub.InsertRow("ApplicationTable", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredLife}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredHealth}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderSumAssured}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPremTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: termLength}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumFrequency}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: appCreatedBy}},				
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}

		return nil, err
		
}


//get the application(depends on the role)
func (t *HDFC) getApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting applicationid to query")
	}

	applicationId := args[0]
	assignerRole := args[1]
	

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + applicationId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the application " + applicationId + "\"}"
		return nil, errors.New(jsonResp)
	}

	
	
	res2E := Application{}
	
	res2E.ApplicationId = row.Columns[0].GetString_()
	res2E.Status = row.Columns[1].GetString_()
	res2E.Title = row.Columns[2].GetString_()
	res2E.FirstName = row.Columns[3].GetString_()
	res2E.LastName = row.Columns[4].GetString_()
	res2E.Gender = row.Columns[5].GetString_()
	res2E.Dob = row.Columns[6].GetString_()
	res2E.Age = row.Columns[7].GetString_()
	res2E.MartialStatus = row.Columns[8].GetString_()
	res2E.FatherName = row.Columns[9].GetString_()
	res2E.MotherName = row.Columns[10].GetString_()
	res2E.Nationality = row.Columns[11].GetString_()
	res2E.ResidentialStatus = row.Columns[12].GetString_()
	res2E.PlaceOfBirth = row.Columns[13].GetString_()
	res2E.PanNumber = row.Columns[14].GetString_()
	res2E.AadharNumber = row.Columns[15].GetString_()
	res2E.EducationalQualification = row.Columns[16].GetString_()
	res2E.PoliticallyExposed = row.Columns[17].GetString_()
	res2E.DisablePersonPolicy = row.Columns[18].GetString_()
	res2E.AnyCriminalProceeding = row.Columns[19].GetString_()
	res2E.LifeApprovalStatus = row.Columns[20].GetString_()
	res2E.HealthApprovalStatus = row.Columns[21].GetString_()
	res2E.LifePlanId = row.Columns[22].GetString_()
	res2E.SumAssuredLife = row.Columns[23].GetString_()
	res2E.HealthPlanId = row.Columns[24].GetString_()
	res2E.SumAssuredHealth = row.Columns[25].GetString_()
	res2E.HealthPremium = row.Columns[26].GetString_()
	res2E.LifePremium = row.Columns[27].GetString_()
	res2E.LifeRiderSumAssured = row.Columns[28].GetString_()
	res2E.LifeRiderTerm = row.Columns[29].GetString_()
	res2E.LifeRiderPremTerm = row.Columns[30].GetString_()
	res2E.LifeRiderPaymentOption = row.Columns[31].GetString_()
	res2E.TermLength = row.Columns[32].GetString_()
	res2E.PremiumTerm = row.Columns[33].GetString_()
	res2E.PremiumFrequency = row.Columns[34].GetString_()
	res2E.PremiumPaymentOption = row.Columns[35].GetString_()
	res2E.AppCreatedBy = row.Columns[36].GetString_()
	
	
    
	
	//getting the role of the user
	//assignerRole, err := stub.ReadCertAttribute("userid")
	//fmt.Printf("Assiger role is %v\n", string(assignerRole))

	//if err != nil {
	//	return nil, fmt.Errorf("Failed getting metadata, [%v]", err)
	//}

	//if len(assignerRole) == 0 {
	//	return nil, errors.New("Invalid assigner role. Empty.")
	//}
	
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	//appCreateOrg1, err := stub.GetState(res2E.AppCreatedBy)
	//appCreateOrg := string(appCreateOrg1)
	
	
	if assignerOrg=="hdfcUW"{
		if res2E.Status == "MANUAL_UW" || res2E.Status == "UW_PENDING" || res2E.Status ==  "UW_REJECT"{
			res2E.HealthPremium = "NA"
		} else{
			return nil, errors.New("You are not authorized")
		}
	}else if assignerOrg=="healthUW"{
		if res2E.Status == "MANUAL_UW" || res2E.Status == "UW_PENDING" || res2E.Status ==  "UW_REJECT"{
			res2E.LifePremium = "NA"
		}else{
			return nil, errors.New("You are not authorized")
		}
	}
	
	
	mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}



//update the under writer status and change the overall status accordingly(depends on the role)
func (t *HDFC) UpdateStatusUW(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}

	applicationId := args[0]
	newStatus := args[1]
	assignerRole := args[2]
	

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()

	
	lifeApprovalStatus:=row.Columns[20].GetString_()
	healthApprovalStatus:=row.Columns[21].GetString_()
	
	
	//getting the role of the user
	//assignerRole, err := stub.ReadCertAttribute("userid")
	//fmt.Printf("Assiger role is %v\n", string(assignerRole))

	//if err != nil {
	//	return nil, fmt.Errorf("Failed getting metadata, [%v]", err)
	//}

	//if len(assignerRole) == 0 {
	//	return nil, errors.New("Invalid assigner role. Empty.")
	//}
	
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	
	if assignerOrg=="hdfcUW"{
		if newStatus=="approved"{
			lifeApprovalStatus = "HDFC_APPROVED"
		} else if newStatus=="rejected"{
			lifeApprovalStatus = "HDFC_REJECT"
		}
		healthApprovalStatus = row.Columns[21].GetString_()
	}else if assignerOrg=="healthUW"{
		lifeApprovalStatus = row.Columns[20].GetString_()
		if newStatus=="approved"{
			healthApprovalStatus = "PARTNER_APPROVED"
		} else if newStatus=="rejected"{
			healthApprovalStatus = "PARTNER_REJECT"
		}
	} else{
		return nil, errors.New("You are not authorized user.")
	}
	
		
	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}	
	
	
	
	//applicationId := row.Columns[0].GetString_()
	status := row.Columns[1].GetString_()
	
	if lifeApprovalStatus== "HDFC_REJECT" || healthApprovalStatus== "PARTNER_REJECT"{
		status = "UW_REJECT"
	} else if lifeApprovalStatus== "HDFC_APPROVED" && healthApprovalStatus== "PARTNER_APPROVED"{
		status = "APPROVED"
	} else if lifeApprovalStatus== "HDFC_APPROVED" || healthApprovalStatus== "PARTNER_APPROVED"{
		status = "UW_PENDING"
	}
	
	
	title := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	gender := row.Columns[5].GetString_()
	dob := row.Columns[6].GetString_()
	age := row.Columns[7].GetString_()
	martialStatus := row.Columns[8].GetString_()
	fatherName := row.Columns[9].GetString_()
	motherName := row.Columns[10].GetString_()
	nationality := row.Columns[11].GetString_()
	residentialStatus := row.Columns[12].GetString_()
	placeOfBirth := row.Columns[13].GetString_()
	panNumber := row.Columns[14].GetString_()
	aadharNumber := row.Columns[15].GetString_()
	educationalQualification := row.Columns[16].GetString_()
	politicallyExposed := row.Columns[17].GetString_()
	disablePersonPolicy := row.Columns[18].GetString_()
	anyCriminalProceeding := row.Columns[19].GetString_()
	//lifeApprovalStatus:=newStatus
	//healthApprovalStatus:=row.Columns[21].GetString_()
	lifePlanId:=row.Columns[22].GetString_()
	sumAssuredLife:=row.Columns[23].GetString_()
	healthPlanId:=row.Columns[24].GetString_()
	sumAssuredHealth:=row.Columns[25].GetString_()
	healthPremium:=row.Columns[26].GetString_()
	lifePremium:=row.Columns[27].GetString_()
	lifeRiderSumAssured:=row.Columns[28].GetString_()
	lifeRiderTerm:=row.Columns[29].GetString_()
	lifeRiderPremTerm:=row.Columns[30].GetString_()
	lifeRiderPaymentOption:=row.Columns[31].GetString_()
	termLength:=row.Columns[32].GetString_()
	premiumTerm:=row.Columns[33].GetString_()
	premiumFrequency:=row.Columns[34].GetString_()
	premiumPaymentOption:=row.Columns[35].GetString_()
	appCreatedBy:=row.Columns[36].GetString_()
	
	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredLife}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredHealth}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderSumAssured}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPremTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: termLength}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumFrequency}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: appCreatedBy}},
		}})
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}


//all applicationid with overall status(irrespective of the role)
func (t *HDFC) listAllApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	assignerRole := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	
	
	// getting the org
	assignerOrg1, err := stub.GetState(assignerRole)
	assignerOrg := string(assignerOrg1)
	
	
	res2E:= []*ListApplication{}	
	
	for row := range rows {
		newApp:= new(ListApplication)
		newApp.ApplicationId = row.Columns[0].GetString_()
		newApp.Status = row.Columns[1].GetString_()
		
		if assignerOrg=="hdfcUW"{
			newApp.LifeApprovalStatus = row.Columns[20].GetString_()
			newApp.HealthApprovalStatus = "NA"
		} else if assignerOrg=="healthUW"{
			newApp.LifeApprovalStatus = "NA"
			newApp.HealthApprovalStatus = row.Columns[21].GetString_()
		} else{
			newApp.LifeApprovalStatus = row.Columns[20].GetString_()
			newApp.HealthApprovalStatus = row.Columns[21].GetString_()
		}
		res2E=append(res2E,newApp)
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}


//application count(irrespective of the role)
func (t *HDFC) getNumApplications(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	var columns []shim.Column

	contractCounter := 0

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	for row := range rows {
		if len(row.Columns) != 0 {
			contractCounter++
		}
	}

	res2E := CountApplication{}
	res2E.Count = contractCounter
	mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
		
	return mapB, nil
}


//update the overall status of the application(irrespective of the role)
func (t *HDFC) UpdateStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3.")
	}

	applicationId := args[0]
	newStatus := args[1]

	// Get the row pertaining to this applicationId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: applicationId}}
	columns = append(columns, col1)

	row, err := stub.GetRow("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Error: Failed retrieving application with applicationId %s. Error %s", applicationId, err.Error())
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		return nil, nil
	}


	//currStatus := row.Columns[1].GetString_()



	//End- Check that the currentStatus to newStatus transition is accurate
	// Delete the row pertaining to this applicationId
	err = stub.DeleteRow(
		"ApplicationTable",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	//applicationId := row.Columns[0].GetString_()
	status := newStatus
	title := row.Columns[2].GetString_()
	firstName := row.Columns[3].GetString_()
	lastName := row.Columns[4].GetString_()
	gender := row.Columns[5].GetString_()
	dob := row.Columns[6].GetString_()
	age := row.Columns[7].GetString_()
	martialStatus := row.Columns[8].GetString_()
	fatherName := row.Columns[9].GetString_()
	motherName := row.Columns[10].GetString_()
	nationality := row.Columns[11].GetString_()
	residentialStatus := row.Columns[12].GetString_()
	placeOfBirth := row.Columns[13].GetString_()
	panNumber := row.Columns[14].GetString_()
	aadharNumber := row.Columns[15].GetString_()
	educationalQualification := row.Columns[16].GetString_()
	politicallyExposed := row.Columns[17].GetString_()
	disablePersonPolicy := row.Columns[18].GetString_()
	anyCriminalProceeding := row.Columns[19].GetString_()
	lifeApprovalStatus:=row.Columns[20].GetString_()
	healthApprovalStatus:=row.Columns[21].GetString_()
	lifePlanId:=row.Columns[22].GetString_()
	sumAssuredLife:=row.Columns[23].GetString_()
	healthPlanId:=row.Columns[24].GetString_()
	sumAssuredHealth:=row.Columns[25].GetString_()
	healthPremium:=row.Columns[26].GetString_()
	lifePremium:=row.Columns[27].GetString_()
	lifeRiderSumAssured:=row.Columns[28].GetString_()
	lifeRiderTerm:=row.Columns[29].GetString_()
	lifeRiderPremTerm:=row.Columns[30].GetString_()
	lifeRiderPaymentOption:=row.Columns[31].GetString_()
	termLength:=row.Columns[32].GetString_()
	premiumTerm:=row.Columns[33].GetString_()
	premiumFrequency:=row.Columns[34].GetString_()
	premiumPaymentOption:=row.Columns[35].GetString_()
	appCreatedBy:=row.Columns[36].GetString_()
	
	//Insert the row pertaining to this applicationId with new status
	_, err = stub.InsertRow(
		"ApplicationTable",
		shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: applicationId}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
				&shim.Column{Value: &shim.Column_String_{String_: title}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: age}},
				&shim.Column{Value: &shim.Column_String_{String_: martialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: fatherName}},
				&shim.Column{Value: &shim.Column_String_{String_: motherName}},
				&shim.Column{Value: &shim.Column_String_{String_: nationality}},
				&shim.Column{Value: &shim.Column_String_{String_: residentialStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: placeOfBirth}},
				&shim.Column{Value: &shim.Column_String_{String_: panNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: aadharNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: educationalQualification}},
				&shim.Column{Value: &shim.Column_String_{String_: politicallyExposed}},
				&shim.Column{Value: &shim.Column_String_{String_: disablePersonPolicy}},
				&shim.Column{Value: &shim.Column_String_{String_: anyCriminalProceeding}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: healthApprovalStatus}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredLife}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPlanId}},
				&shim.Column{Value: &shim.Column_String_{String_: sumAssuredHealth}},
				&shim.Column{Value: &shim.Column_String_{String_: healthPremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifePremium}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderSumAssured}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPremTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: lifeRiderPaymentOption}},
				&shim.Column{Value: &shim.Column_String_{String_: termLength}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumTerm}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumFrequency}},
				&shim.Column{Value: &shim.Column_String_{String_: premiumPaymentOption}},	
				&shim.Column{Value: &shim.Column_String_{String_: appCreatedBy}},	
		}})
		 
	if err != nil {
		return nil, errors.New("Failed inserting row.")
	}

	return nil, nil

}


//get the application by pan number (depends on the role)
func (t *HDFC) getApplicationByPanNumber(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	lastPan := args[0]
	assignerRole := args[1]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E := Application{}
	
	for row := range rows {
		fetchedPan := row.Columns[14].GetString_()
		
		if fetchedPan == lastPan{
			res2E.ApplicationId = row.Columns[0].GetString_()
			res2E.Status = row.Columns[1].GetString_()
			res2E.Title = row.Columns[2].GetString_()
			res2E.FirstName = row.Columns[3].GetString_()
			res2E.LastName = row.Columns[4].GetString_()
			res2E.Gender = row.Columns[5].GetString_()
			res2E.Dob = row.Columns[6].GetString_()
			res2E.Age = row.Columns[7].GetString_()
			res2E.MartialStatus = row.Columns[8].GetString_()
			res2E.FatherName = row.Columns[9].GetString_()
			res2E.MotherName = row.Columns[10].GetString_()
			res2E.Nationality = row.Columns[11].GetString_()
			res2E.ResidentialStatus = row.Columns[12].GetString_()
			res2E.PlaceOfBirth = row.Columns[13].GetString_()
			res2E.PanNumber = row.Columns[14].GetString_()
			res2E.AadharNumber = row.Columns[15].GetString_()
			res2E.EducationalQualification = row.Columns[16].GetString_()
			res2E.PoliticallyExposed = row.Columns[17].GetString_()
			res2E.DisablePersonPolicy = row.Columns[18].GetString_()
			res2E.AnyCriminalProceeding = row.Columns[19].GetString_()
			res2E.LifeApprovalStatus = row.Columns[20].GetString_()
			res2E.HealthApprovalStatus = row.Columns[21].GetString_()
			res2E.LifePlanId = row.Columns[22].GetString_()
			res2E.SumAssuredLife = row.Columns[23].GetString_()
			res2E.HealthPlanId = row.Columns[24].GetString_()
			res2E.SumAssuredHealth = row.Columns[25].GetString_()
			res2E.HealthPremium = row.Columns[26].GetString_()
			res2E.LifePremium = row.Columns[27].GetString_()
			res2E.LifeRiderSumAssured = row.Columns[28].GetString_()
			res2E.LifeRiderTerm = row.Columns[29].GetString_()
			res2E.LifeRiderPremTerm = row.Columns[30].GetString_()
			res2E.LifeRiderPaymentOption = row.Columns[31].GetString_()
			res2E.TermLength = row.Columns[32].GetString_()
			res2E.PremiumTerm = row.Columns[33].GetString_()
			res2E.PremiumFrequency = row.Columns[34].GetString_()
			res2E.PremiumPaymentOption = row.Columns[35].GetString_()
			res2E.AppCreatedBy = row.Columns[36].GetString_()
			
			//getting the role of the user
			//assignerRole, err := stub.ReadCertAttribute("userid")
			//fmt.Printf("Assiger role is %v\n", string(assignerRole))

			//if err != nil {
			//	return nil, fmt.Errorf("Failed getting metadata, [%v]", err)
			//}

			//if len(assignerRole) == 0 {
			//	return nil, errors.New("Invalid assigner role. Empty.")
			//}
			
			assignerOrg1, err := stub.GetState(assignerRole)
			assignerOrg := string(assignerOrg1)
			if err != nil {
				return nil, fmt.Errorf("Failed getting the assignee role, [%v]", err)
			}
			//appCreateOrg1, err := stub.GetState(res2E.AppCreatedBy)
			//appCreateOrg := string(appCreateOrg1)
			
			
			if assignerOrg=="hdfcUW"{
				if res2E.Status == "MANUAL_UW" || res2E.Status == "UW_PENDING" || res2E.Status ==  "UW_REJECT"{
					res2E.HealthPremium = "NA"
				} else{
					return nil, errors.New("You are not authorized")
				}
			}else if assignerOrg=="healthUW"{
				if res2E.Status == "MANUAL_UW" || res2E.Status == "UW_PENDING" || res2E.Status ==  "UW_REJECT"{
					res2E.LifePremium = "NA"
				}else{
					return nil, errors.New("You are not authorized")
				}
			}
			
			
			mapB, _ := json.Marshal(res2E)
			fmt.Println(string(mapB))
			
			return mapB, nil
	}
	}
	return nil, errors.New("There is no application with the specified Pan Number.")

}


//get all applicationid with overall status by status (irrespective of the role)
func (t *HDFC) listAllApplicationByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	argStatus := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E:= []*ListApplication{}	
	
	for row := range rows {
		fetchedStatus := row.Columns[1].GetString_()
		
		if fetchedStatus == argStatus{
			newApp:= new(ListApplication)
			newApp.ApplicationId = row.Columns[0].GetString_()
			newApp.Status = row.Columns[1].GetString_()
			res2E=append(res2E,newApp)
		}
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}


//get all applicationid with overall status by last name (irrespective of the role)
func (t *HDFC) listAllApplicationByLastName(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1.")
	}

	argLastName := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("ApplicationTable", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}

	res2E:= []*ListApplication{}	
	
	for row := range rows {
		fetchedLastName := row.Columns[4].GetString_()
		
		if fetchedLastName == argLastName{
			newApp:= new(ListApplication)
			newApp.ApplicationId = row.Columns[0].GetString_()
			newApp.Status = row.Columns[1].GetString_()
			res2E=append(res2E,newApp)
		}
	}
	
	res2F, _ := json.Marshal(res2E)
    fmt.Println(string(res2F))
	return res2F, nil

}





// Invoke invokes the chaincode
func (t *HDFC) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "submitApplication" {
		t := HDFC{}
		return t.submitApplication(stub, args)
	}else if function == "UpdateStatusUW" { 
		t := HDFC{}
		return t.UpdateStatusUW(stub, args)
	} else if function == "submitQuote" { 
		t := HDFC{}
		return t.submitQuote(stub, args)
	}  else if function == "UpdateStatus" { 
		t := HDFC{}
		return t.UpdateStatus(stub, args)
	}  

	return nil, errors.New("Invalid invoke function name.")

}

func (t *HDFC) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getApplication" {
		t := HDFC{}
		return t.getApplication(stub, args)		
	}else if function == "listAllApplication" { 
		t := HDFC{}
		return t.listAllApplication(stub, args)
	}else if function == "getNumApplications" { 
		t := HDFC{}
		return t.getNumApplications(stub, args)
	}else if function == "getApplicationByPanNumber" { 
		t := HDFC{}
		return t.getApplicationByPanNumber(stub, args)
	}else if function == "listAllApplicationByStatus" { 
		t := HDFC{}
		return t.listAllApplicationByStatus(stub, args)
	}else if function == "listAllApplicationByLastName" { 
		t := HDFC{}
		return t.listAllApplicationByLastName(stub, args)
	}else if function == "getQuote" { 
		t := HDFC{}
		return t.getQuote(stub, args)
	}else if function == "listAllQuote" { 
		t := HDFC{}
		return t.listAllQuote(stub, args)
	}
	
	return nil, nil
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(HDFC))
	if err != nil {
		fmt.Printf("Error starting HDFC: %s", err)
	}
} 
