package SalseTracker

import (
	constants "SalesTracker/constants"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	DB "SalesTracker/DB"
)

type Response struct{
	Status  string            `json:"status"`
	ErrMsg  string            `json:"errMsg"`
}


func SyncDataApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {
		log.Println("SyncDataApi (+)")
		var lResp Response
		lErr := SyncData() 
		if lErr != nil {
			log.Println(lErr.Error() + " Error(STSDA001) ")
			lResp.Status = constants.ErrorCode
			lResp.ErrMsg = lErr.Error()
		} else {
			lResp.Status = constants.SuccessCode
			lResp.ErrMsg = ""
		}

		lData, lErr := json.Marshal(lResp)
		if lErr != nil {
			fmt.Fprint(w, "Error taking data Error(STSDA002) " +lErr.Error())
		} else {
			fmt.Fprint(w, string(lData))
		}

		log.Println("SyncDataApi (-)")
	}
}


func SyncData() error {
	log.Println("SyncData (+)")
	var lData [][]string
	var lUpdateId int64
	var lErr error

	var lProduct map[string]string  // is used to avoid the duplicate recod to be inserted
	var lCustomer map[string]string  // is used to avoid the duplicate recod to be inserted
	var lsales map[string]string  // is used to avoid the duplicate recod to be inserted

	lUpdateId,lErr = LogSync("IN","",0)
	if lErr!=nil{
		LogSync("OUT","E",lUpdateId)
		log.Println(" Error(STSD001.0) :"+lErr.Error())
		return errors.New(" Error(STSD001.0) :"+lErr.Error())
	}

	lData,lErr = Readfile("./SalesData/salesData.csv")
	if lErr!=nil{
		LogSync("OUT","E",lUpdateId)
		log.Println(" Error(STSD001) :"+lErr.Error())
		return errors.New(" Error(STSD001) :"+lErr.Error())
	}

	lProduct,lErr = GetExistID("select  p.ProductID from ProductMaster p")
	if lErr!=nil{
		LogSync("OUT","E",lUpdateId)
		log.Println(" Error(STSD001.0) :"+lErr.Error())
		return errors.New(" Error(STSD001.0) :"+lErr.Error())
	}

	lCustomer,lErr = GetExistID("select  c.CustomerID from CoustomerDetalis c")
	if lErr!=nil{
		LogSync("OUT","E",lUpdateId)
		log.Println(" Error(STSD001.0) :"+lErr.Error())
		return errors.New(" Error(STSD001.0) :"+lErr.Error())
	}

	lsales,lErr = GetExistID("select  s.OrderID from Sales s")
	if lErr!=nil{
		LogSync("OUT","E",lUpdateId)
		log.Println(" Error(STSD001.0) :"+lErr.Error())
		return errors.New(" Error(STSD001.0) :"+lErr.Error())
	}

	lSalesInsertQry := `insert into Sales 
	(OrderID, DateofSale, QuantitySold, UnitPrice, Discount, ShippingCost , CustomerID, ProductID, PaymentMethod)
	values `
	lSaleCount:=0
	lSalesInsertVal:=``
	
	lProductInsertQry := `insert into ProductMaster 
	(ProductID, ProductName, Category)
	values `
	lProductInsertVal:=``
	lProductCount:=0

	lCustomertInsertQry := `insert into CoustomerDetalis 
	(CustomerID, CustomerName, CustomerEmail, CustomerAddress, Region)
	values `
	lCustomertInsertVal:=``
	lCustomerCount:=0

	for lIdx,lRows  := range lData {
		if lIdx == 0{
			continue
		}

		var lProductstructRec ProductMaster
		var lCoustmerDtlsRec CoustomerDetalis
		var lSalesRec Sales

		
		lProductstructRec.ProductID =strings.ReplaceAll(strings.TrimSpace(lRows[1]), "'", "''") // avoiding the escapeing charter and spaceing
		lProductstructRec.ProductName = strings.ReplaceAll(strings.TrimSpace(lRows[3]), "'", "''")
		lProductstructRec.Category = strings.ReplaceAll(strings.TrimSpace(lRows[4]), "'", "''")
		
		lCoustmerDtlsRec.CustomerID = strings.ReplaceAll(strings.TrimSpace(lRows[2]), "'", "''")
		lCoustmerDtlsRec.CustomerName = strings.ReplaceAll(strings.TrimSpace(lRows[12]), "'", "''")
		lCoustmerDtlsRec.CustomerEmail = strings.ReplaceAll(strings.TrimSpace(lRows[13]), "'", "''")
		lCoustmerDtlsRec.CustomerAddress = strings.ReplaceAll(strings.TrimSpace(lRows[14]), "'", "''")
		lCoustmerDtlsRec.Region = strings.ReplaceAll(strings.TrimSpace(lRows[5]), "'", "''")
		
		lSalesRec.OrderID = strings.ReplaceAll(strings.TrimSpace(lRows[0]), "'", "''")
		lSalesRec.DateofSale = strings.TrimSpace(lRows[6])
		lSalesRec.Discount,_ = strconv.ParseFloat(strings.TrimSpace(lRows[9]), 64)
		lSalesRec.PaymentMethod = strings.ReplaceAll(strings.TrimSpace(lRows[11]), "'", "''")
		lSalesRec.QuantitySold,_ = strconv.ParseFloat(strings.TrimSpace(lRows[7]), 64)
		lSalesRec.UnitPrice,_ = strconv.ParseFloat(strings.TrimSpace(lRows[8]), 64)
		lSalesRec.ShippingCost,_ = strconv.ParseFloat(strings.TrimSpace(lRows[10]), 64)

		if _,lExist:=lProduct[lProductstructRec.ProductID]  ;!lExist{
			lProductInsertVal+=`(`+
			`'`+lProductstructRec.ProductID+`',`+
			`'`+lProductstructRec.ProductName+`',`+
			`'`+lProductstructRec.Category+`'`+
			`),`
			lProductCount++
		}

		if _,lExist:=lCustomer[lCoustmerDtlsRec.CustomerID]  ;!lExist{
			lCustomertInsertVal+=`(`+
			`'`+lCoustmerDtlsRec.CustomerID+`',`+
			`'`+lCoustmerDtlsRec.CustomerName+`',`+
			`'`+lCoustmerDtlsRec.CustomerEmail +`',`+
			`'`+lCoustmerDtlsRec.CustomerAddress +`',`+
			`'`+lCoustmerDtlsRec.Region +`'`+
			`),`
			lCustomerCount++
		}

		if _,lExist:=lsales[lSalesRec.OrderID]  ;!lExist{
			lSaleCount++
			lSalesInsertVal+=`(`+
			`'`+lSalesRec.OrderID+`',`+
			`'`+ lSalesRec.DateofSale+`',`+
			`'`+strconv.FormatFloat(lSalesRec.QuantitySold, 'f', 2, 64)+`',`+
			`'`+strconv.FormatFloat(lSalesRec.UnitPrice, 'f', 2, 64)+`',`+
			`'`+strconv.FormatFloat(lSalesRec.Discount, 'f', 2, 64)+`',`+
			`'`+strconv.FormatFloat(lSalesRec.ShippingCost, 'f', 2, 64)+`',`+
			`'`+ lCoustmerDtlsRec.CustomerID+`',`+
			`'`+ lProductstructRec.ProductID+`',`+
			`'`+lSalesRec.PaymentMethod+`'`+
			`),`
		}



		if lSaleCount>=1000{
			lErr = DB.InsertBulkData(DB.GDbConn.MariaDB,lSalesInsertVal,lSalesInsertQry)
			if lErr!=nil{
				LogSync("OUT","E",lUpdateId)
				log.Println(" Error(STSD002) :"+lErr.Error())
				return errors.New(" Error(STSD002) :"+lErr.Error())
			}
			lSaleCount=0
			lSalesInsertVal=``
		}



		if lProductCount >= 1000{
			lErr = DB.InsertBulkData(DB.GDbConn.MariaDB,lProductInsertVal,lProductInsertQry)
			if lErr!=nil{
				LogSync("OUT","E",lUpdateId)
				log.Println(" Error(STSD003) :"+lErr.Error())
				return errors.New(" Error(STSD003) :"+lErr.Error())
			}
			lProductCount = 0   
			lProductInsertVal = ``
		}
		


		if lCustomerCount >= 1000{
			lErr = DB.InsertBulkData(DB.GDbConn.MariaDB,lCustomertInsertVal,lCustomertInsertQry)
			if lErr!=nil{
				LogSync("OUT","E",lUpdateId)
				log.Println(" Error(STSD004) :"+lErr.Error())
				return errors.New(" Error(STSD004) :"+lErr.Error())
			}
			lCustomerCount = 0 
			lCustomertInsertVal = ``
		}
		lProduct[lProductstructRec.ProductID]="Y"
		lCustomer[lCoustmerDtlsRec.CustomerID]="Y"
		lsales[lSalesRec.OrderID]="Y"
	}
	if lCustomerCount > 0{
		lErr = DB.InsertBulkData(DB.GDbConn.MariaDB,lCustomertInsertVal,lCustomertInsertQry)
		if lErr!=nil{
			LogSync("OUT","E",lUpdateId)
			log.Println(" Error(STSD005) :"+lErr.Error())
			return errors.New(" Error(STSD005) :"+lErr.Error())
		} 
		lCustomerCount=0
		lCustomertInsertVal=``
	}

	if lProductCount > 0{
	lErr = DB.InsertBulkData(DB.GDbConn.MariaDB,lProductInsertVal,lProductInsertQry)
		if lErr!=nil{
			LogSync("OUT","E",lUpdateId)
			log.Println(" Error(STSD006) :"+lErr.Error())
			return errors.New(" Error(STSD006) :"+lErr.Error())
		}
		lProductCount=0
		lProductInsertVal=`` 
	}

	if lSaleCount > 0{
	lErr = DB.InsertBulkData(DB.GDbConn.MariaDB,lSalesInsertVal,lSalesInsertQry)
		if lErr!=nil{
			LogSync("OUT","E",lUpdateId)
			log.Println(" Error(STSD007) :"+lErr.Error())
			return errors.New(" Error(STSD007) :"+lErr.Error())
		}
		lSaleCount=0
		lSalesInsertVal=``
	}
	LogSync("OUT","S",lUpdateId)

	log.Println("SyncData (-)")
	return  nil
}

func LogSync(pPunch, pSyncStatus string,pId int64)(lLastInserted int64,lErr error){
	log.Println("LogSync (+)")

	lQueryString := ``
	if strings.EqualFold(pPunch,"IN")  {
		lQueryString = `insert into  SyncTracking (StartPunch) values (now())`
	}else if strings.EqualFold(pPunch,"OUT"){
		lQueryString = `
		update  synctracking  
		set SyncSuccess='`+pSyncStatus+`', EndPunch =now()
		where  id = `+ fmt.Sprint(pId)
	}

	lResult, lErr := DB.GDbConn.MariaDB.Exec(lQueryString)
	if lErr != nil {
		log.Println(lErr)
		return lLastInserted , lErr
	}
	
	lLastInserted, lErr = lResult.LastInsertId()
	if lErr != nil {
		log.Println(lErr)
		return lLastInserted , lErr
	}

	log.Println("LogSync (-)")
	return lLastInserted ,  lErr
}



func GetExistID(pQueryString string)( map[string]string, error){
	log.Println("GetExistID (+)")

	var lErr error
	lResult:= make(map[string]string)

		lRows, lErr := DB.GDbConn.MariaDB.Query(pQueryString)
		if lErr != nil {
			log.Println(" Error(STGP001) :"+lErr.Error())
			return lResult, errors.New(" Error(STGP001) :"+lErr.Error())
		} 
		defer lRows.Close()
		for lRows.Next() {
			var lID string
			lErr := lRows.Scan(&lID)
			if lErr != nil {
				log.Println(" Error(STGP002) :"+lErr.Error())
				return lResult, errors.New(" Error(STGP002) :"+lErr.Error())
			} 
			lResult[lID]="Y"			
		}

		if lErr := lRows.Err();lErr != nil{
			log.Println(" Error(STGP003) :"+lErr.Error())
			return lResult, errors.New(" Error(STGP003) :"+lErr.Error())
		}

	log.Println("GetExistID (-)")
	return lResult, lErr
}




