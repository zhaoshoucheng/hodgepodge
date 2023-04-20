package main

type GeoIP struct {
	IspName     string `json:"isp_name"` //
	CountryName string `json:"country_name"`
}

type ScoreTag struct {
	Name string `json:"name"`
}
type Threaten struct {
	Type      string   `json:"type"`
	RiskLevel string   `json:"risk_level"` //作为结果
	ScoreTag  ScoreTag `json:"score_tag"`
}
type Source struct {
	SourceIP      string   `json:"sourceip"`
	GeoIP         GeoIP    `json:"geoip"`
	HttpUserAgent string   `json:"http_user_agent"`
	HttpReferer   string   `json:"http_referer"`
	Reason        string   `json:"reason"`
	Timestamp     string   `json:"@timestamp"`
	Threaten      Threaten `json:"threaten"`
}

type ESDataResp struct {
	Hits struct {
		Total int `json:"total"`
		Hits  []struct {
			ID     string `json:"_id"`
			Source Source `json:"_source"`
		}
	} `json:"hits"`
}
