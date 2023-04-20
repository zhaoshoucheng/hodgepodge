package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/zhaoshoucheng/hodgepodge/net"
	"time"
)

const getESData = `
{
  "query": {
    "range": {
      "@timestamp": {
        "gte": %d,
        "lte": %d,
        "format": "epoch_millis"
      }
    }
  }
}
`

var localCache *cache.Cache

const searchUrl = "https://lkfe.lkcoffee.com/elasticsearch/waf-*/_search"

func InitLoaclCache() {
	localCache = cache.New(30*time.Second, time.Minute)
}
func SearchDataFromES(ctx context.Context, start int64, end int64) ([]*Source, error) {
	body := fmt.Sprintf(getESData, start, end)
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := net.PostQueryWithHeaders(ctx, searchUrl, "", string(bodyStr), "application/json", nil)
	if err != nil {
		return nil, err
	}
	respData := ESDataResp{}
	err = json.Unmarshal([]byte(resp), &respData)
	if err != nil {
		return nil, err
	}
	var sourceList []*Source
	for _, hit := range respData.Hits.Hits {
		_, find := localCache.Get(hit.ID)
		if find {
			continue
		}
		localCache.Set(hit.ID, struct{}{}, time.Minute*5)
		sourceList = append(sourceList, &hit.Source)
	}
	return sourceList, nil
}
