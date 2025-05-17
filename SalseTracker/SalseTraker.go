package SalseTracker

import (
	constants "SalesTracker/constants"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Response struct{
	Status  string            `json:"status"`
	ErrMsg  string            `json:"errMsg"`
}


func DemoApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {
		log.Println("DemoApi (+)")


		lVars := mux.Vars(r)
		lId := lVars["id"]
		log.Println(lId)
		lResp, lErr := Demo() 
		if lErr != nil {
			log.Println(lErr.Error() + " Error(TDA001) ")
			lResp.Status = constants.ErrorCode
			lResp.ErrMsg = lErr.Error()
		} else {
			lResp.Status = constants.SuccessCode
			lResp.ErrMsg = ""
		}

		lData, lErr := json.Marshal(lResp)
		if lErr != nil {
			fmt.Fprint(w, "Error taking data Error(TDA002) " +lErr.Error())
		} else {
			fmt.Fprint(w, string(lData))
		}

		log.Println("DemoApi (-)")
	}
}

func Demo()(lResp Response,lErr error){
	log.Println("Demo (+)")

	lResp.Status = constants.SuccessCode

	log.Println("Demo (-)")
	return lResp, lErr
}



func SyncData() ( error) {
	log.Println("SyncData (+)")
	

	lData,lErr := Readfile("./SalesData/salesData.csv")
	if lErr!=nil{
		return lErr
	}

	for lIdx,lRows  := range lData {
		if lIdx==0{
			continue
		}

		var lProductstructRec ProductMaster
		var lCoustmerDtlsRec CoustomerDetalis
		var lSalesRec Sales

		lSalesRec.OrderID = strings.TrimSpace(lRows[0])
		lSalesRec.DateofSale = strings.TrimSpace(lRows[6])
		lSalesRec.Discount,_ = strconv.ParseFloat(strings.TrimSpace(lRows[9]), 64)
		lSalesRec.PaymentMethod = strings.TrimSpace(lRows[11])
		lSalesRec.QuantitySold,_ = strconv.ParseFloat(strings.TrimSpace(lRows[7]), 64)
		lSalesRec.UnitPrice,_ = strconv.ParseFloat(strings.TrimSpace(lRows[8]), 64)
		lSalesRec.ShippingCost,_ = strconv.ParseFloat(strings.TrimSpace(lRows[10]), 64)

		lProductstructRec.ProductID = strings.TrimSpace(lRows[1])
		lProductstructRec.ProductName = strings.TrimSpace(lRows[3])
		lProductstructRec.Category = strings.TrimSpace(lRows[4])

		lCoustmerDtlsRec.CustomerID = strings.TrimSpace(lRows[2])
		lCoustmerDtlsRec.CustomerName = strings.TrimSpace(lRows[12])
		lCoustmerDtlsRec.CustomerEmail = strings.TrimSpace(lRows[13])
		lCoustmerDtlsRec.CustomerAddress = strings.TrimSpace(lRows[14])
		lCoustmerDtlsRec.Region = strings.TrimSpace(lRows[5])

		// lQueryString1:=``




		
	}

	log.Println("SyncData (-)")
	return  nil
}




