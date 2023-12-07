package main

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var (
	ACCESS_TOKEN = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC90ZXN0LWFkbWluLmRvdXl1eGluZ2NoZW4uY29tXC9hcGlcL2Vucm9sbFwvcGFzc3BvcnQtbG9naW4iLCJpYXQiOjE2ODQyMjUyMTgsImV4cCI6MTY4NTA4OTIxOCwibmJmIjoxNjg0MjI1MjE4LCJqdGkiOiJMYVlaNGxCbUk0eDQxUWlkIiwic3ViIjoxMDk5LCJwcnYiOiI0YmZkOGJjNmU3YjRmYTY2MWVhMmJkMzRiMThjZWI0YTQ5ODA3Y2VjIn0.cuBj5ofRr6Kt62DQIaWZjB2waotMvOg7O6_SX58IJbY"
	URL          = "https://test-admin.douyuxingchen.com/api/bus-order/sale-after/apply"
)

func main() {
	orderList := []string{"20230317112939900353521"}
	for i := 0; i < len(orderList); i++ {
		params := map[string]interface{}{
			"after_sale_type": 1,
			"bus_order_no":    orderList[i],
			"reason_remark":   "系统批量处理",
		}
		resp, err := httpPost(context.Background(), URL, params)
		if err != nil {

		}
	}
}

func httpPost(ctx context.Context, path string, params map[string]interface{}) ([]byte, error) {
	p, _ := json.Marshal(params)
	req, err := http.NewRequest(http.MethodPost, path, strings.NewReader(string(p)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", ACCESS_TOKEN)
	req = req.WithContext(ctx)
	client := &http.Client{Timeout: time.Second * 60}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	buffer := make([]byte, 2048)
	length, _ := resp.Body.Read(buffer)
	return buffer[:length], nil
}
