package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := `{"errorCode":"200","errorMessage":"success","currentTime":1693831953120,"success":true,"data":{"code":200,"message":"success","requestId":"1693831953120","data":{"result":[{"id":"901829635483013120","orderId":276179899609,"parentId":0,"orderTime":"2023-09-04 13:49:00","finishTime":"","modifyTime":"2023-09-04 13:49:00","orderEmt":2,"plus":1,"unionId":1001133999,"skuId":10034328079196,"skuName":"闽光（MINGUANG）【懒人神器】小白鞋清洁剂去污皮鞋洗鞋刷鞋神器洗懒人型洗鞋神器 小白鞋清洁剂2瓶装","skuNum":1,"skuReturnNum":0,"skuFrozenNum":0,"price":39.90,"commissionRate":30.0000,"subSideRate":90.0000,"subsidyRate":0.0000,"finalRate":90.0000,"estimateCosPrice":9.90,"estimateFee":2.67,"actualCosPrice":0.00,"actualFee":0.00,"validCode":15,"traceType":2,"positionId":3100626496,"siteId":1535574920,"unionAlias":"","pid":"","cid1":null,"cid2":null,"cid3":null,"subUnionId":"","unionTag":"00000000000000000000000000000010","popId":11718677,"ext1":"b4f7f489c3f2d21e|__TRACK_ID__","payMonth":"0","cpActId":0,"unionRole":1,"giftCouponOcsAmount":0.00,"giftCouponKey":"","balanceExt":"","proPriceAmount":0.00,"rid":0,"goodsInfo":{"imageUrl":"http://img11.360buyimg.com/n0/jfs/t1/149211/7/25355/158783/622f234eEd63632f6/2d7a20a9fa54bc7e.jpg","owner":"p","mainSkuId":10034328079196,"productId":10021305632177,"shopName":"Min悦个护家清专营店","shopId":11452011},"categoryInfo":{"cid1":null,"cid1Name":"家庭清洁/纸品","cid2":null,"cid2Name":"皮具护理","cid3":null,"cid3Name":"小白鞋清洗剂"}}],"pageSize":500,"pageNo":1,"totalItems":1,"hasNext":false},"totalCount":1,"success":true}}`
	bytes := []byte(a)
	response, _ := NewFxOrderRowQueryResponse(bytes)
	//fmt.Println(err, response.OrderFxRowQueryResponse.Datastr)
	//os.Exit(2)
	//dataList := response.OrderFxRowQueryResponse.GetFxOrderRowQueryResult()
	//for _, i := range dataList.Data {
	//	fmt.Println(i)
	//}

	for _, result := range response.OrderFxRowQueryResponse.Datastr.Result {
		fmt.Println(result)
	}
}

func NewOrderRowQueryResponse(body []byte) (*OrderRowQueryResponse, error) {
	r := &OrderRowQueryResponse{}
	err := json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func NewFxOrderRowQueryResponse(body []byte) (*FxOrderRowQueryResponse, error) {
	r := &FxOrderRowQueryResponse{}
	s := string(body)
	err := json.Unmarshal(body, r)
	//fmt.Println(err, r)
	if err != nil || s == "" {
		return nil, err
	}
	//fmt.Println(123, r)
	return r, nil
}

type OrderRowQueryResponse struct {
	OrderRowQueryResponse *OrderRowQueryResponse_ `json:"jd_union_open_order_row_query_responce"`
}

type FxOrderRowQueryResponse struct {
	OrderFxRowQueryResponse *OrderFxRowQueryResponse_ `json:"data"`
}

type OrderFxRowQueryResponse_ struct {
	Datastr *Data `json:"data"`
}

type Data struct {
	Result   []interface{} `json:"result"`
	PageSize int           `json:"pageSize"`
	PageNo   int           `json:"pageNo"`
}

type OrderRowQueryResponse_ struct {
	Code           string `json:"code"`
	QueryResultRaw string `json:"queryResult"`
}

func (o *OrderRowQueryResponse_) GetOrderRowQueryResult() *OrderRowQueryResult {
	r := &OrderRowQueryResult{}
	err := json.Unmarshal([]byte(o.QueryResultRaw), r)
	if err != nil {
		return nil
	}
	return r
}

//func (o *OrderFxRowQueryResponse_) GetFxOrderRowQueryResult() *OrderRowQueryResult {
//	r := &OrderRowQueryResult{}
//	err := json.Unmarshal([]byte(o.Datastr), r)
//	if err != nil {
//		return nil
//	}
//	return r
//}

type OrderRowQueryResult struct {
	RequestId string `json:"requestId"`
	Code      int    `json:"code"`
	Data      string `json:"data"`
	HasMore   bool   `json:"hasMore"`
	Message   string `json:"message"`
}
