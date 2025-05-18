package SalseTracker

import (
	"SalesTracker/DB"
	constants "SalesTracker/constants"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func GetRevenueApi(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET")
	(w).Header().Set("Access-Control-Allow-Headers", "FROMDT,TODT,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	if r.Method == "GET" {
		log.Println("GetRevenueApi (+)")
		lVars := mux.Vars(r)
		lId := lVars["id"]
		log.Println(lId)

		lFromDt := r.Header.Get("FROMDT")
		lToDt := r.Header.Get("TODT")

		log.Println(lFromDt)
		log.Println(lToDt)


		lResp,lErr := GetRevenue(lId,lFromDt,lToDt)
		if lErr != nil {
			log.Println(lErr.Error() + " Error(STGRA001) ")
			lResp.Status = constants.ErrorCode
			lResp.ErrMsg = lErr.Error()
		} else {
			lResp.Status = constants.SuccessCode
			lResp.ErrMsg = ""
		}

		lData, lErr := json.Marshal(lResp)
		if lErr != nil {
			fmt.Fprint(w, "Error taking data Error(STGRA002) "+lErr.Error())
		} else {
			fmt.Fprint(w, string(lData))
		}

		log.Println("GetRevenueApi (-)")
	}
}

func GetRevenue(pType, pFromDt ,pToDt string) (lResp RevenueResp, lErr error) {
	log.Println("GetRevenue (+)")
	var lRec Revenue

	lQueryString := `
	`
	switch strings.ToUpper(pType) {

	case "TOTAL":	// finding the total revenue
		var lTotal float64
		lQueryString = `
		select  round(sum((QuantitySold * UnitPrice)- Discount+ ShippingCost),2) totalRevenue from  sales s 
		where  DateofSale >= ? and DateofSale <= ? 
		`
			lRows, lErr := DB.GDbConn.MariaDB.Query(lQueryString,pFromDt,pToDt)
		if lErr != nil {
			log.Println(" Error(STGR001) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR001) :"+lErr.Error())
		} 
		defer lRows.Close()
		if lRows.Next() {
			lErr := lRows.Scan(&lTotal)
			if lErr != nil {
				log.Println(" Error(STGR002) :"+lErr.Error())
				return lResp, errors.New(" Error(STGR002) :"+lErr.Error())
			} 
			lResp.TotalRevenue=	lTotal	
		}

		if lErr := lRows.Err();lErr != nil{
			log.Println(" Error(STGR003) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR003) :"+lErr.Error())
		}
		log.Println("GetRevenue (-)")
		return lResp, lErr

	case "PRODUCT":	// finding product wise total revenue
		lQueryString = `
		select 
		Nvl((select  p.ProductName from productmaster p where p.ProductID =s.ProductID),'') productName,
		round(sum((QuantitySold * UnitPrice)- Discount+ ShippingCost),2) totalRevenue
		from  sales s 
		where  DateofSale >= ? and DateofSale <= ?
		group by ProductID 
		`
		lRows, lErr := DB.GDbConn.MariaDB.Query(lQueryString,pFromDt,pToDt)
		if lErr != nil {
			log.Println(" Error(STGR004) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR004) :"+lErr.Error())
		} 
		defer lRows.Close()
		for lRows.Next() {
			lErr := lRows.Scan(&lRec.Product,&lRec.Revenue)
			if lErr != nil {
				log.Println(" Error(STGR005) :"+lErr.Error())
				return lResp, errors.New(" Error(STGR005) :"+lErr.Error())
			} 
			lResp.ProductRevenue = append(lResp.ProductRevenue, lRec)
			
		}

		if lErr := lRows.Err();lErr != nil{
			log.Println(" Error(STGR006) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR006) :"+lErr.Error())
		}
		log.Println("GetRevenue (-)")
		return lResp, lErr
	
	case "CATEGORY":	// finding category wise total revenue
		lQueryString = `
		select  Nvl(p.Category,'') Category ,
		round(sum((QuantitySold * UnitPrice)- Discount+ ShippingCost),2) totalRevenue
		from  sales s 
		inner join productmaster p on p.ProductID = s.ProductID 
		where  DateofSale >= ? and DateofSale <= ?
		group by p.Category
		`
		lRows, lErr := DB.GDbConn.MariaDB.Query(lQueryString,pFromDt,pToDt)
		if lErr != nil {
			log.Println(" Error(STGR007) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR007) :"+lErr.Error())
		} 
		defer lRows.Close()
		for lRows.Next() {
			lErr := lRows.Scan(&lRec.Category,&lRec.Revenue)
			if lErr != nil {
				log.Println(" Error(STGR008) :"+lErr.Error())
				return lResp, errors.New(" Error(STGR008) :"+lErr.Error())
			} 
			lResp.CategoryRevenue = append(lResp.CategoryRevenue, lRec)
			
		}

		if lErr := lRows.Err();lErr != nil{
			log.Println(" Error(STGR009) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR009) :"+lErr.Error())
		}
		log.Println("GetRevenue (-)")
		return lResp, lErr

	case "REGION":	// finding region wise total revenue
		lQueryString = `
		select  Nvl(c.Region,'') Region ,
		round(sum((QuantitySold * UnitPrice)- Discount+ ShippingCost),2) totalRevenue
		from  sales s 
		inner join coustomerdetalis c  on c.CustomerID= s.CustomerID
		where  DateofSale >= ? and DateofSale <= ?
		group by c.Region  
		`
			lRows, lErr := DB.GDbConn.MariaDB.Query(lQueryString,pFromDt,pToDt)
		if lErr != nil {
			log.Println(" Error(STGR0010) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR0010) :"+lErr.Error())
		} 
		defer lRows.Close()
		for lRows.Next() {
			lErr := lRows.Scan(&lRec.Region,&lRec.Revenue)
			if lErr != nil {
				log.Println(" Error(STGR0011) :"+lErr.Error())
				return lResp, errors.New(" Error(STGR0011) :"+lErr.Error())
			} 
			lResp.RegionRevenue = append(lResp.RegionRevenue, lRec)
			
		}

		if lErr := lRows.Err();lErr != nil{
			log.Println(" Error(STGR0012) :"+lErr.Error())
			return lResp, errors.New(" Error(STGR0012) :"+lErr.Error())
		}
		log.Println("GetRevenue (-)")
		return lResp, lErr
	
default:
	
	log.Println("GetRevenue (-)")
	return lResp, errors.New(" Invalid Revenue Type")
	}

}
