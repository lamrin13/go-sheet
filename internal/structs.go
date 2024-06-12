package internal

type Request struct {
	Date   string  `json:"date"`
	Income float32 `json:"income"`
	Spend  float32 `json:"spend"`
	Remark string  `json:"remark"`
}

type Response struct {
	Date   string  `json:"date"`
	Income float32 `json:"income"`
	Spend  float32 `json:"spend"`
	Remark string  `json:"remark"`
}
