package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{"id":"18ed9e59-3be9-41b0-8573-1897ad79e9ef","source":{"tid":"18ed9e59-3be9-41b0-8573-1897ad79e9ef"},"tmax":1000,"imp":[{"id":"Div1","ext":{"pubmatic":{"publisherId":"  5890  ","adSlot":"       /43743431/DMDemo@728x90:0","pmzoneid":"zone1","kadpageurl":"www.yahoo.com","lat":"40.712775  ","lon":"   -74.005973","yob":"  1982  ","gender":"   M   ","kadfloor":"  1.75   "},"appnexus":{"use_pmt_rule":false,"placement_id":13144370}},"banner":{"format":[{"w":300,"h":250},{"w":728,"h":90},{"w":160,"h":600}]}}],"test":1,"ext":{"prebid":{"targeting":{"includewinners":true,"includebidderkeys":false}}},"site":{"publisher":{"id":"1"},"page":"http://ebay.com/inte/automation/pubmatic_in_prebid_v1/prebidserver-singleslot.html?pbjs_debug=true"},"device":{"w":726,"h":793}}`
	// data := `{
	// 	"id": 11,
	// 	"name": "Irshad",
	// 	"department": "IT",
	// 	"designation": "Product Manager",
	// 	"address": {
	// 		"city": "Mumbai",
	// 		"state": "Maharashtra",
	// 		"country": "India"
	// 	}
	// }`
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("unmarshal error", err)
		return
	}
	add := result["imp"].(map[string]interface{})
	for k := range add {
		fmt.Println(k)
	}

}
