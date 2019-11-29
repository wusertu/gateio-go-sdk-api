package gateio

import "fmt"

// ErrMsg 错误格式
type ErrMsg struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

// QueryResponse 单次请求, 订阅后回包
type QueryResponse struct {
	ID     int64       `json:"id"`
	Err    *ErrMsg     `json:"error"`
	Result interface{} `json:"result"`
}

// Valid 检验对象是否合法
func (t *QueryResponse) Valid() bool {
	return t.ID > 0 && t.Err == nil && t.Result != nil
}

// SubscribeResponse 服务器主动推送数据格式
type SubscribeResponse struct {
	ID     int64         `json:"id"`
	Method string        `json:"method"`
	Result []interface{} `json:"params"`
}

type Ticker struct {
	Period      int    `json:"period"`
	Open        string `json:"open"`
	Close       string `json:"close"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Last        string `json:"last"`
	Change      string `json:"change"`
	QuoteVolume string `json:"quoteVolume"`
	BaseVolume  string `json:"baseVolume"`
}

type Trade struct {
	ID     int     `json:"id"`
	Time   float64 `json:"time"`
	Price  string  `json:"price"`
	Amount string  `json:"amount"`
	Type   string  `json:"type"`
}

type KLine struct {
	Time   float64 `json:"time"`
	Open   string  `json:"open"`
	Close  string  `json:"close"`
	High   string  `json:"high"`
	Low    string  `json:"low"`
	Volume string  `json:"volume"`
	Amount string  `json:"amount"`
	Symbol string  `json:"symbol"`
}

// Valid 检验对象是否合法
func (t *SubscribeResponse) Valid() bool {
	return len(t.Method) > 0 && t.Result != nil
}

func loadResponse(msg []byte) (interface{}, error) {
	var err error

	subscribeResp := SubscribeResponse{}
	err = JsonBytes2Struct(msg, &subscribeResp)
	if err == nil && subscribeResp.Valid() {
		return &subscribeResp, nil
	}

	queryResp := QueryResponse{}
	err = JsonBytes2Struct(msg, &queryResp)
	if err == nil && queryResp.Valid() {
		return &queryResp, nil
	}

	fmt.Println("invalid data", string(msg))

	return nil, nil
}

// ReceivedDataCallback 接收 ws 数据的回调方法 interface
type ReceivedDataCallback func(interface{}) error

// DefaultPrintData 测试时用于输出服务器推送数据
func DefaultPrintData(obj interface{}) error {
	switch obj.(type) {
	case string:
		fmt.Println(obj)
	default:
		msg, err := Struct2JsonString(obj)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(msg)

	}
	return nil
}