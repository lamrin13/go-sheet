package internal

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

var sheetSrv *sheet

func init() {
	tmp, err := getService()
	if err != nil {
		log.Fatal(err)
	}
	sheetSrv = tmp
}

func querySheet(id string, sheetRange string) (*sheets.ValueRange, error) {
	resp, err := sheetSrv.Srv.Spreadsheets.Values.Get(id, sheetRange).
		DateTimeRenderOption("FORMATTED_STRING").
		ValueRenderOption("UNFORMATTED_VALUE").
		Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func appendRow(id string, sheetRange string, values *sheets.ValueRange) (*sheets.AppendValuesResponse, error) {
	resp, err := sheetSrv.Srv.Spreadsheets.Values.Append(id, sheetRange, values).
		ValueInputOption("USER_ENTERED").
		Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func searchRow(id string, searchReq *sheets.BatchGetValuesByDataFilterRequest) (*sheets.BatchGetValuesByDataFilterResponse, error) {
	resp, err := sheetSrv.Srv.Spreadsheets.Values.BatchGetByDataFilter(id, searchReq).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getSheets(id string) (*sheets.Spreadsheet, error) {
	resp, err := sheetSrv.Srv.Spreadsheets.Get(id).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getAllData(id string, tab string) (*sheets.ValueRange, error) {
	return sheetSrv.Srv.Spreadsheets.Values.Get(id, tab).Do()
}
