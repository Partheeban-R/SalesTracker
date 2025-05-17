package SalseTracker

import (
	"encoding/csv"
	"os"
)

func Readfile(pfilePath string) (lRecords[][]string,lErr error){
	lFile, lErr := os.Open(pfilePath)
	if lErr != nil {
		return lRecords, lErr
	}
	defer lFile.Close()

	lReader := csv.NewReader(lFile)
	lRecords, lErr = lReader.ReadAll()
	
	if lErr != nil {
		return lRecords,lErr
	}
	return lRecords,nil
}