package internal

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/api/sheets/v4"
)

func Append(columns, id string, rows []Request) ([]byte, error) {
	var (
		body []byte
		resp *sheets.AppendValuesResponse
		err  error
		data [][]interface{}
	)

	sheetRange := "All In One" + columns

	for _, row := range rows {
		_, err = time.Parse("01/02/2006", row.Date)
		_, err = time.Parse("2006-01-02", row.Date)
		if err != nil {
			return nil, err
		}
		tmp := []interface{}{row.Date, row.Income, row.Spend, row.Remark}
		data = append(data, tmp)

	}
	// fmt.Println("Data ", data, r)
	values := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         data,
	}
	fmt.Println("reached here")
	resp, err = appendRow(id, sheetRange, values)
	if err != nil {
		return nil, err
	}
	body, err = resp.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func GetAll(id string) ([]byte, error) {
	res, err := getAllData(id, "All In One!A:D")
	if err != nil {
		return nil, err
	}
	var rawRes []Response
	for _, row := range res.Values {
		date, ok := row[0].(string)
		if !ok {
			date = "invalid date"
		}
		income, err := strconv.ParseFloat(row[1].(string), 64)
		if err != nil {
			income = 0
		}
		spend, err := strconv.ParseFloat(row[2].(string), 64)
		if err != nil {
			spend = 0
		}
		remark, ok := row[3].(string)
		if !ok {
			remark = "N/A"
		}
		rawRes = append(rawRes, Response{
			Date:   date,
			Income: float32(income),
			Spend:  float32(spend),
			Remark: remark,
		})
	}

	return json.Marshal(rawRes[1:])
}

func GetCategories(id string) ([]byte, error) {
	resp, err := getAllData(id, "Category!A:A")
	if err != nil {
		return nil, err
	}
	var categories []string
	for _, v := range resp.Values[1:] {
		categories = append(categories, v[0].(string))
	}
	return json.Marshal(categories)
}
