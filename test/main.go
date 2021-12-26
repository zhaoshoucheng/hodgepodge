package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type RespInfo struct {
	Errno int `json:"errno"`
	Errmsg string `json:"errmsg"`
	Data struct{
		Avg []struct{
			Desc []struct{
				Name string `json:"name"`
				Value string `json:"value"`
			} `json:"desc"`
			X string `json:"x"`
		} `json:"avg"`
		LastAvg []struct{
			Desc []struct{
				Name string `json:"name"`
				Value string `json:"value"`
			} `json:"desc"`
		} `json:"last_avg"`
	} `json:"data"`
}

type data struct {
	ID int
	Source int
	SV int
}

func main() {
	//url1 := "http://172.22.41.32:8080/test_service_name/v1/do2"
	//url2 := "http://127.0.0.1:8080/test_service_name/v1/do2"
	//url3 := "http://localhost:8080/test_service_name/v1/do2"
	//go Get(url1,"172.22.41.32")

	str := `{"errno":0,"errmsg":"","data":{"2021-12-12":[{"x":"00:00","desc":[{"name":"PI","value":"23.09"},{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"15.21"},{"name":"停车次数","value":"0.48"}]},{"x":"00:30","desc":[{"name":"PI","value":"7.56"},{"name":"单次停车比率","value":"0.20"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"8.71"},{"name":"停车次数","value":"0.30"}]},{"x":"01:00","desc":[{"name":"PI","value":"17.35"},{"name":"单次停车比率","value":"0.32"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"13.47"},{"name":"停车次数","value":"0.34"}]},{"x":"01:30","desc":[{"name":"PI","value":"23.43"},{"name":"单次停车比率","value":"0.55"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"19.64"},{"name":"停车次数","value":"0.59"}]},{"x":"02:00","desc":[{"name":"PI","value":"18.59"},{"name":"单次停车比率","value":"0.44"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"8.59"},{"name":"停车次数","value":"0.28"}]},{"x":"02:30","desc":[{"name":"PI","value":"18.21"},{"name":"单次停车比率","value":"0.57"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"10.13"},{"name":"停车次数","value":"0.41"}]},{"x":"03:00","desc":[{"name":"PI","value":"15.23"},{"name":"单次停车比率","value":"0.27"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"16.06"},{"name":"停车次数","value":"0.42"}]},{"x":"03:30","desc":[{"name":"PI","value":"35.83"},{"name":"单次停车比率","value":"0.73"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"27.61"},{"name":"停车次数","value":"0.74"}]},{"x":"04:00","desc":[{"name":"PI","value":"31.79"},{"name":"单次停车比率","value":"0.86"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"21.50"},{"name":"停车次数","value":"0.80"}]},{"x":"05:30","desc":[{"name":"PI","value":"40.83"},{"name":"单次停车比率","value":"0.67"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"25.91"},{"name":"停车次数","value":"0.41"}]},{"x":"06:00","desc":[{"name":"PI","value":"43.04"},{"name":"单次停车比率","value":"0.64"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"32.68"},{"name":"停车次数","value":"0.53"}]},{"x":"06:30","desc":[{"name":"PI","value":"50.55"},{"name":"单次停车比率","value":"0.78"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"39.94"},{"name":"停车次数","value":"0.60"}]},{"x":"07:00","desc":[{"name":"PI","value":"33.30"},{"name":"单次停车比率","value":"0.51"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"30.96"},{"name":"停车次数","value":"0.55"}]},{"x":"07:30","desc":[{"name":"PI","value":"56.54"},{"name":"单次停车比率","value":"0.58"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"40.45"},{"name":"停车次数","value":"0.57"}]},{"x":"08:00","desc":[{"name":"PI","value":"39.25"},{"name":"单次停车比率","value":"0.55"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"37.31"},{"name":"停车次数","value":"0.54"}]},{"x":"08:30","desc":[{"name":"PI","value":"26.85"},{"name":"单次停车比率","value":"0.57"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"38.76"},{"name":"停车次数","value":"0.57"}]},{"x":"09:00","desc":[{"name":"PI","value":"32.68"},{"name":"单次停车比率","value":"0.60"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"39.04"},{"name":"停车次数","value":"0.60"}]},{"x":"09:30","desc":[{"name":"PI","value":"74.71"},{"name":"单次停车比率","value":"0.69"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"36.91"},{"name":"停车次数","value":"0.51"}]},{"x":"10:00","desc":[{"name":"PI","value":"24.95"},{"name":"单次停车比率","value":"0.52"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"37.12"},{"name":"停车次数","value":"0.55"}]},{"x":"10:30","desc":[{"name":"PI","value":"20.50"},{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"33.56"},{"name":"停车次数","value":"0.48"}]},{"x":"11:00","desc":[{"name":"PI","value":"23.88"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"30.77"},{"name":"停车次数","value":"0.44"}]},{"x":"11:30","desc":[{"name":"PI","value":"29.53"},{"name":"单次停车比率","value":"0.32"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"29.49"},{"name":"停车次数","value":"0.41"}]},{"x":"12:00","desc":[{"name":"PI","value":"53.50"},{"name":"单次停车比率","value":"0.56"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"29.57"},{"name":"停车次数","value":"0.42"}]},{"x":"12:30","desc":[{"name":"PI","value":"81.89"},{"name":"单次停车比率","value":"0.58"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.02"},{"name":"延误时间","value":"33.15"},{"name":"停车次数","value":"0.44"}]},{"x":"13:00","desc":[{"name":"PI","value":"26.11"},{"name":"单次停车比率","value":"0.42"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"29.45"},{"name":"停车次数","value":"0.42"}]},{"x":"13:30","desc":[{"name":"PI","value":"36.49"},{"name":"单次停车比率","value":"0.49"},{"name":"过饱和比率","value":"0.03"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"38.01"},{"name":"停车次数","value":"0.54"}]},{"x":"14:00","desc":[{"name":"PI","value":"30.43"},{"name":"单次停车比率","value":"0.36"},{"name":"过饱和比率","value":"0.03"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"32.40"},{"name":"停车次数","value":"0.48"}]},{"x":"14:30","desc":[{"name":"PI","value":"62.36"},{"name":"单次停车比率","value":"0.79"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"37.69"},{"name":"停车次数","value":"0.53"}]},{"x":"15:00","desc":[{"name":"PI","value":"51.92"},{"name":"单次停车比率","value":"0.69"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"31.00"},{"name":"停车次数","value":"0.46"}]},{"x":"15:30","desc":[{"name":"PI","value":"27.29"},{"name":"单次停车比率","value":"0.47"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"36.35"},{"name":"停车次数","value":"0.51"}]},{"x":"16:00","desc":[{"name":"PI","value":"36.04"},{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"36.57"},{"name":"停车次数","value":"0.53"}]},{"x":"16:30","desc":[{"name":"PI","value":"21.91"},{"name":"单次停车比率","value":"0.41"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"35.49"},{"name":"停车次数","value":"0.47"}]},{"x":"17:00","desc":[{"name":"PI","value":"21.40"},{"name":"单次停车比率","value":"0.49"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"36.92"},{"name":"停车次数","value":"0.51"}]},{"x":"17:30","desc":[{"name":"PI","value":"50.45"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"40.47"},{"name":"停车次数","value":"0.59"}]},{"x":"18:00","desc":[{"name":"PI","value":"22.10"},{"name":"单次停车比率","value":"0.47"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"36.40"},{"name":"停车次数","value":"0.51"}]},{"x":"18:30","desc":[{"name":"PI","value":"28.95"},{"name":"单次停车比率","value":"0.46"},{"name":"过饱和比率","value":"0.03"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"36.28"},{"name":"停车次数","value":"0.51"}]},{"x":"19:00","desc":[{"name":"PI","value":"22.74"},{"name":"单次停车比率","value":"0.43"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"32.84"},{"name":"停车次数","value":"0.47"}]},{"x":"19:30","desc":[{"name":"PI","value":"62.00"},{"name":"单次停车比率","value":"0.74"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"38.02"},{"name":"停车次数","value":"0.52"}]},{"x":"20:00","desc":[{"name":"PI","value":"29.74"},{"name":"单次停车比率","value":"0.44"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"34.25"},{"name":"停车次数","value":"0.48"}]},{"x":"20:30","desc":[{"name":"PI","value":"32.90"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"32.26"},{"name":"停车次数","value":"0.42"}]},{"x":"21:00","desc":[{"name":"PI","value":"19.30"},{"name":"单次停车比率","value":"0.42"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"27.41"},{"name":"停车次数","value":"0.43"}]},{"x":"21:30","desc":[{"name":"PI","value":"49.65"},{"name":"单次停车比率","value":"0.49"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":0},{"name":"延误时间","value":"27.31"},{"name":"停车次数","value":"0.37"}]},{"x":"22:00","desc":[{"name":"PI","value":"35.80"},{"name":"单次停车比率","value":"0.47"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"20.87"},{"name":"停车次数","value":"0.32"}]},{"x":"22:30","desc":[{"name":"PI","value":"17.88"},{"name":"单次停车比率","value":"0.26"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"23.52"},{"name":"停车次数","value":"0.33"}]},{"x":"23:00","desc":[{"name":"PI","value":"29.66"},{"name":"单次停车比率","value":"0.62"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"21.92"},{"name":"停车次数","value":"0.59"}]},{"x":"23:30","desc":[{"name":"PI","value":"17.72"},{"name":"单次停车比率","value":"0.37"},{"name":"过饱和比率","value":0},{"name":"溢流比率","value":0},{"name":"延误时间","value":"13.22"},{"name":"停车次数","value":"0.39"}]}],"avg":[{"x":"00:00","desc":[{"name":"PI","value":"23.09"},{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"15.21"},{"name":"停车次数","value":"0.48"}]},{"x":"00:30","desc":[{"name":"延误时间","value":"8.71"},{"name":"停车次数","value":"0.30"},{"name":"PI","value":"7.56"},{"name":"单次停车比率","value":"0.20"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"01:00","desc":[{"name":"PI","value":"17.35"},{"name":"单次停车比率","value":"0.32"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"13.47"},{"name":"停车次数","value":"0.34"}]},{"x":"01:30","desc":[{"name":"延误时间","value":"19.64"},{"name":"停车次数","value":"0.59"},{"name":"PI","value":"23.43"},{"name":"单次停车比率","value":"0.55"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"02:00","desc":[{"name":"PI","value":"18.59"},{"name":"单次停车比率","value":"0.44"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"8.59"},{"name":"停车次数","value":"0.28"}]},{"x":"02:30","desc":[{"name":"PI","value":"18.21"},{"name":"单次停车比率","value":"0.57"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"10.13"},{"name":"停车次数","value":"0.41"}]},{"x":"03:00","desc":[{"name":"PI","value":"15.23"},{"name":"单次停车比率","value":"0.27"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"16.06"},{"name":"停车次数","value":"0.42"}]},{"x":"03:30","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"27.61"},{"name":"停车次数","value":"0.74"},{"name":"PI","value":"35.83"},{"name":"单次停车比率","value":"0.73"},{"name":"过饱和比率","value":"0.00"}]},{"x":"04:00","desc":[{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"21.50"},{"name":"停车次数","value":"0.80"},{"name":"PI","value":"31.79"},{"name":"单次停车比率","value":"0.86"}]},{"x":"05:30","desc":[{"name":"单次停车比率","value":"0.67"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"25.91"},{"name":"停车次数","value":"0.41"},{"name":"PI","value":"40.83"}]},{"x":"06:00","desc":[{"name":"单次停车比率","value":"0.64"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"32.68"},{"name":"停车次数","value":"0.53"},{"name":"PI","value":"43.04"}]},{"x":"06:30","desc":[{"name":"停车次数","value":"0.60"},{"name":"PI","value":"50.55"},{"name":"单次停车比率","value":"0.78"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"39.94"}]},{"x":"07:00","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"30.96"},{"name":"停车次数","value":"0.55"},{"name":"PI","value":"33.30"},{"name":"单次停车比率","value":"0.51"},{"name":"过饱和比率","value":"0.00"}]},{"x":"07:30","desc":[{"name":"延误时间","value":"40.45"},{"name":"停车次数","value":"0.57"},{"name":"PI","value":"56.54"},{"name":"单次停车比率","value":"0.58"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.01"}]},{"x":"08:00","desc":[{"name":"PI","value":"39.25"},{"name":"单次停车比率","value":"0.55"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"37.31"},{"name":"停车次数","value":"0.54"}]},{"x":"08:30","desc":[{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"38.76"},{"name":"停车次数","value":"0.57"},{"name":"PI","value":"26.85"},{"name":"单次停车比率","value":"0.57"}]},{"x":"09:00","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"39.04"},{"name":"停车次数","value":"0.60"},{"name":"PI","value":"32.68"},{"name":"单次停车比率","value":"0.60"},{"name":"过饱和比率","value":"0.00"}]},{"x":"09:30","desc":[{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.91"},{"name":"停车次数","value":"0.51"},{"name":"PI","value":"74.71"},{"name":"单次停车比率","value":"0.69"}]},{"x":"10:00","desc":[{"name":"PI","value":"24.95"},{"name":"单次停车比率","value":"0.52"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"37.12"},{"name":"停车次数","value":"0.55"}]},{"x":"10:30","desc":[{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"33.56"},{"name":"停车次数","value":"0.48"},{"name":"PI","value":"20.50"},{"name":"单次停车比率","value":"0.48"}]},{"x":"11:00","desc":[{"name":"延误时间","value":"30.77"},{"name":"停车次数","value":"0.44"},{"name":"PI","value":"23.88"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"}]},{"x":"11:30","desc":[{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"29.49"},{"name":"停车次数","value":"0.41"},{"name":"PI","value":"29.53"},{"name":"单次停车比率","value":"0.32"}]},{"x":"12:00","desc":[{"name":"停车次数","value":"0.42"},{"name":"PI","value":"53.50"},{"name":"单次停车比率","value":"0.56"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"29.57"}]},{"x":"12:30","desc":[{"name":"PI","value":"81.89"},{"name":"单次停车比率","value":"0.58"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.02"},{"name":"延误时间","value":"33.15"},{"name":"停车次数","value":"0.44"}]},{"x":"13:00","desc":[{"name":"PI","value":"26.11"},{"name":"单次停车比率","value":"0.42"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"29.45"},{"name":"停车次数","value":"0.42"}]},{"x":"13:30","desc":[{"name":"单次停车比率","value":"0.49"},{"name":"过饱和比率","value":"0.03"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"38.01"},{"name":"停车次数","value":"0.54"},{"name":"PI","value":"36.49"}]},{"x":"14:00","desc":[{"name":"PI","value":"30.43"},{"name":"单次停车比率","value":"0.36"},{"name":"过饱和比率","value":"0.03"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"32.40"},{"name":"停车次数","value":"0.48"}]},{"x":"14:30","desc":[{"name":"延误时间","value":"37.69"},{"name":"停车次数","value":"0.53"},{"name":"PI","value":"62.36"},{"name":"单次停车比率","value":"0.79"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"15:00","desc":[{"name":"停车次数","value":"0.46"},{"name":"PI","value":"51.92"},{"name":"单次停车比率","value":"0.69"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"31.00"}]},{"x":"15:30","desc":[{"name":"PI","value":"27.29"},{"name":"单次停车比率","value":"0.47"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.35"},{"name":"停车次数","value":"0.51"}]},{"x":"16:00","desc":[{"name":"停车次数","value":"0.53"},{"name":"PI","value":"36.04"},{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.57"}]},{"x":"16:30","desc":[{"name":"PI","value":"21.91"},{"name":"单次停车比率","value":"0.41"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"35.49"},{"name":"停车次数","value":"0.47"}]},{"x":"17:00","desc":[{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.92"},{"name":"停车次数","value":"0.51"},{"name":"PI","value":"21.40"},{"name":"单次停车比率","value":"0.49"}]},{"x":"17:30","desc":[{"name":"PI","value":"50.45"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"40.47"},{"name":"停车次数","value":"0.59"}]},{"x":"18:00","desc":[{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.40"},{"name":"停车次数","value":"0.51"},{"name":"PI","value":"22.10"},{"name":"单次停车比率","value":"0.47"}]},{"x":"18:30","desc":[{"name":"延误时间","value":"36.28"},{"name":"停车次数","value":"0.51"},{"name":"PI","value":"28.95"},{"name":"单次停车比率","value":"0.46"},{"name":"过饱和比率","value":"0.03"},{"name":"溢流比率","value":"0.00"}]},{"x":"19:00","desc":[{"name":"单次停车比率","value":"0.43"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"32.84"},{"name":"停车次数","value":"0.47"},{"name":"PI","value":"22.74"}]},{"x":"19:30","desc":[{"name":"停车次数","value":"0.52"},{"name":"PI","value":"62.00"},{"name":"单次停车比率","value":"0.74"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"38.02"}]},{"x":"20:00","desc":[{"name":"PI","value":"29.74"},{"name":"单次停车比率","value":"0.44"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"34.25"},{"name":"停车次数","value":"0.48"}]},{"x":"20:30","desc":[{"name":"PI","value":"32.90"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"32.26"},{"name":"停车次数","value":"0.42"}]},{"x":"21:00","desc":[{"name":"PI","value":"19.30"},{"name":"单次停车比率","value":"0.42"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"27.41"},{"name":"停车次数","value":"0.43"}]},{"x":"21:30","desc":[{"name":"单次停车比率","value":"0.49"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"27.31"},{"name":"停车次数","value":"0.37"},{"name":"PI","value":"49.65"}]},{"x":"22:00","desc":[{"name":"单次停车比率","value":"0.47"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"20.87"},{"name":"停车次数","value":"0.32"},{"name":"PI","value":"35.80"}]},{"x":"22:30","desc":[{"name":"PI","value":"17.88"},{"name":"单次停车比率","value":"0.26"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"23.52"},{"name":"停车次数","value":"0.33"}]},{"x":"23:00","desc":[{"name":"单次停车比率","value":"0.62"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"21.92"},{"name":"停车次数","value":"0.59"},{"name":"PI","value":"29.66"}]},{"x":"23:30","desc":[{"name":"PI","value":"17.72"},{"name":"单次停车比率","value":"0.37"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"13.22"},{"name":"停车次数","value":"0.39"}]}],"last_avg":[{"x":"00:00","desc":[{"name":"PI","value":"14.04"},{"name":"单次停车比率","value":"0.31"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"12.10"},{"name":"停车次数","value":"0.34"}]},{"x":"00:30","desc":[{"name":"PI","value":"22.32"},{"name":"单次停车比率","value":"0.59"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"16.00"},{"name":"停车次数","value":"0.55"}]},{"x":"01:00","desc":[{"name":"PI","value":"8.19"},{"name":"单次停车比率","value":"0.20"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"7.22"},{"name":"停车次数","value":"0.24"}]},{"x":"01:30","desc":[{"name":"PI","value":"35.00"},{"name":"单次停车比率","value":"0.69"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"15.29"},{"name":"停车次数","value":"0.41"}]},{"x":"02:00","desc":[{"name":"PI","value":"26.64"},{"name":"单次停车比率","value":"0.60"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"24.23"},{"name":"停车次数","value":"0.58"}]},{"x":"02:30","desc":[{"name":"单次停车比率","value":"0.75"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"13.79"},{"name":"停车次数","value":"0.41"},{"name":"PI","value":"26.88"}]},{"x":"03:00","desc":[{"name":"延误时间","value":"31.04"},{"name":"停车次数","value":"0.58"},{"name":"PI","value":"37.50"},{"name":"单次停车比率","value":"0.80"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"04:00","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"24.35"},{"name":"停车次数","value":"0.68"},{"name":"PI","value":"49.42"},{"name":"单次停车比率","value":"1.00"},{"name":"过饱和比率","value":"0.00"}]},{"x":"04:30","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"26.11"},{"name":"停车次数","value":"0.72"},{"name":"PI","value":"37.17"},{"name":"单次停车比率","value":"0.73"},{"name":"过饱和比率","value":"0.00"}]},{"x":"05:00","desc":[{"name":"单次停车比率","value":"0.11"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"15.00"},{"name":"停车次数","value":"0.20"},{"name":"PI","value":"2.45"}]},{"x":"05:30","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"22.86"},{"name":"停车次数","value":"0.39"},{"name":"PI","value":"26.54"},{"name":"单次停车比率","value":"0.46"},{"name":"过饱和比率","value":"0.00"}]},{"x":"06:00","desc":[{"name":"停车次数","value":"0.46"},{"name":"PI","value":"63.70"},{"name":"单次停车比率","value":"0.52"},{"name":"过饱和比率","value":"0.04"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"29.83"}]},{"x":"06:30","desc":[{"name":"停车次数","value":"0.51"},{"name":"PI","value":"42.22"},{"name":"单次停车比率","value":"0.61"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"31.77"}]},{"x":"07:00","desc":[{"name":"PI","value":"48.26"},{"name":"单次停车比率","value":"0.61"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"37.99"},{"name":"停车次数","value":"0.61"}]},{"x":"07:30","desc":[{"name":"停车次数","value":"0.51"},{"name":"PI","value":"28.68"},{"name":"单次停车比率","value":"0.51"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"35.77"}]},{"x":"08:00","desc":[{"name":"PI","value":"55.77"},{"name":"单次停车比率","value":"0.61"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"48.19"},{"name":"停车次数","value":"0.66"}]},{"x":"08:30","desc":[{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"45.26"},{"name":"停车次数","value":"0.65"},{"name":"PI","value":"36.05"},{"name":"单次停车比率","value":"0.63"}]},{"x":"09:00","desc":[{"name":"停车次数","value":"0.53"},{"name":"PI","value":"18.78"},{"name":"单次停车比率","value":"0.51"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.32"}]},{"x":"09:30","desc":[{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"39.29"},{"name":"停车次数","value":"0.56"},{"name":"PI","value":"31.65"},{"name":"单次停车比率","value":"0.52"}]},{"x":"10:00","desc":[{"name":"停车次数","value":"0.55"},{"name":"PI","value":"28.05"},{"name":"单次停车比率","value":"0.53"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"35.41"}]},{"x":"10:30","desc":[{"name":"延误时间","value":"42.11"},{"name":"停车次数","value":"0.59"},{"name":"PI","value":"25.80"},{"name":"单次停车比率","value":"0.55"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"}]},{"x":"11:00","desc":[{"name":"PI","value":"25.92"},{"name":"单次停车比率","value":"0.46"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"35.59"},{"name":"停车次数","value":"0.50"}]},{"x":"11:30","desc":[{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"32.28"},{"name":"停车次数","value":"0.45"},{"name":"PI","value":"21.47"},{"name":"单次停车比率","value":"0.42"}]},{"x":"12:00","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"23.71"},{"name":"停车次数","value":"0.37"},{"name":"PI","value":"16.36"},{"name":"单次停车比率","value":"0.35"},{"name":"过饱和比率","value":"0.00"}]},{"x":"12:30","desc":[{"name":"PI","value":"31.36"},{"name":"单次停车比率","value":"0.47"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"31.68"},{"name":"停车次数","value":"0.49"}]},{"x":"13:00","desc":[{"name":"单次停车比率","value":"0.51"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"36.94"},{"name":"停车次数","value":"0.53"},{"name":"PI","value":"26.54"}]},{"x":"13:30","desc":[{"name":"延误时间","value":"36.67"},{"name":"停车次数","value":"0.53"},{"name":"PI","value":"22.28"},{"name":"单次停车比率","value":"0.51"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"14:00","desc":[{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"34.58"},{"name":"停车次数","value":"0.50"},{"name":"PI","value":"26.17"},{"name":"单次停车比率","value":"0.46"}]},{"x":"14:30","desc":[{"name":"单次停车比率","value":"0.36"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"25.19"},{"name":"停车次数","value":"0.41"},{"name":"PI","value":"16.28"}]},{"x":"15:00","desc":[{"name":"停车次数","value":"0.49"},{"name":"PI","value":"72.43"},{"name":"单次停车比率","value":"0.73"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"33.14"}]},{"x":"15:30","desc":[{"name":"延误时间","value":"34.53"},{"name":"停车次数","value":"0.53"},{"name":"PI","value":"22.89"},{"name":"单次停车比率","value":"0.50"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"16:00","desc":[{"name":"停车次数","value":"0.60"},{"name":"PI","value":"28.38"},{"name":"单次停车比率","value":"0.56"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"40.34"}]},{"x":"16:30","desc":[{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"39.35"},{"name":"停车次数","value":"0.51"},{"name":"PI","value":"30.86"}]},{"x":"17:00","desc":[{"name":"溢流比率","value":"0.01"},{"name":"延误时间","value":"37.29"},{"name":"停车次数","value":"0.55"},{"name":"PI","value":"51.43"},{"name":"单次停车比率","value":"0.32"},{"name":"过饱和比率","value":"0.02"}]},{"x":"17:30","desc":[{"name":"停车次数","value":"0.57"},{"name":"PI","value":"45.78"},{"name":"单次停车比率","value":"0.43"},{"name":"过饱和比率","value":"0.04"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"38.00"}]},{"x":"18:00","desc":[{"name":"PI","value":"68.88"},{"name":"单次停车比率","value":"0.71"},{"name":"过饱和比率","value":"0.02"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"34.14"},{"name":"停车次数","value":"0.50"}]},{"x":"18:30","desc":[{"name":"延误时间","value":"36.39"},{"name":"停车次数","value":"0.51"},{"name":"PI","value":"26.50"},{"name":"单次停车比率","value":"0.48"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"19:00","desc":[{"name":"PI","value":"20.13"},{"name":"单次停车比率","value":"0.37"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"26.41"},{"name":"停车次数","value":"0.41"}]},{"x":"19:30","desc":[{"name":"PI","value":"28.42"},{"name":"单次停车比率","value":"0.40"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"29.62"},{"name":"停车次数","value":"0.44"}]},{"x":"20:00","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"25.88"},{"name":"停车次数","value":"0.34"},{"name":"PI","value":"71.55"},{"name":"单次停车比率","value":"0.54"},{"name":"过饱和比率","value":"0.04"}]},{"x":"20:30","desc":[{"name":"PI","value":"23.71"},{"name":"单次停车比率","value":"0.38"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"29.09"},{"name":"停车次数","value":"0.41"}]},{"x":"21:00","desc":[{"name":"单次停车比率","value":"0.57"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"24.98"},{"name":"停车次数","value":"0.39"},{"name":"PI","value":"58.37"}]},{"x":"21:30","desc":[{"name":"PI","value":"28.01"},{"name":"单次停车比率","value":"0.34"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"25.82"},{"name":"停车次数","value":"0.35"}]},{"x":"22:00","desc":[{"name":"单次停车比率","value":"0.39"},{"name":"过饱和比率","value":"0.01"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"29.73"},{"name":"停车次数","value":"0.41"},{"name":"PI","value":"23.05"}]},{"x":"22:30","desc":[{"name":"延误时间","value":"22.60"},{"name":"停车次数","value":"0.31"},{"name":"PI","value":"37.65"},{"name":"单次停车比率","value":"0.45"},{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"}]},{"x":"23:00","desc":[{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"24.97"},{"name":"停车次数","value":"0.56"},{"name":"PI","value":"32.95"},{"name":"单次停车比率","value":"0.60"},{"name":"过饱和比率","value":"0.00"}]},{"x":"23:30","desc":[{"name":"过饱和比率","value":"0.00"},{"name":"溢流比率","value":"0.00"},{"name":"延误时间","value":"16.61"},{"name":"停车次数","value":"0.47"},{"name":"PI","value":"20.10"},{"name":"单次停车比率","value":"0.44"}]}]},"trace_id":"ac3601ba61b85af8c47d16545b4dc6b0"}`
	resp := RespInfo{}
	err := json.Unmarshal([]byte(str),&resp)
	if err != nil {
		fmt.Println(err)
		return
	}

	avgMap := make(map[string]float64)
	lastAvgMap := make(map[string]float64)
	for _, agvInfo := range resp.Data.Avg {
		for index, desc := range agvInfo.Desc {
			if index == 0 && desc.Name == "PI" {
				continue
			}
			if index == 1 && desc.Name == "单次停车比率" {
				continue
			}
			if index == 2 && desc.Name == "过饱和比率" {
				continue
			}
			if index == 3 && desc.Name == "溢流比率" {
				continue
			}
			if index == 4 && desc.Name == "延误时间" {
				continue
			}
			if index == 5 && desc.Name == "停车次数" {
				continue
			}
			fmt.Println("agvInfo", agvInfo)
			if _,exists := avgMap[desc.Name]; !exists {
				avgMap[desc.Name] = 0
			}
			value, err := strconv.ParseFloat(desc.Value,64)
			if err != nil {
				fmt.Println(err)
			}
			avgMap[desc.Name] = avgMap[desc.Name] + value
		}
	}

	fmt.Println("---------------------------------")
	for _, lastAvginfo := range resp.Data.LastAvg {
		for _, desc := range lastAvginfo.Desc {
			if _,exists := lastAvgMap[desc.Name]; !exists {
				lastAvgMap[desc.Name] = 0
				fmt.Println(desc.Name)
			}
			value, err := strconv.ParseFloat(desc.Value,64)
			if err != nil {
				fmt.Println(err)
			}
			lastAvgMap[desc.Name] = lastAvgMap[desc.Name] + value
		}
	}

	fmt.Println(len(resp.Data.Avg))

	for key , value := range avgMap {
		if lastValue, exists := lastAvgMap[key];exists {
			fmt.Println(key, " ",value/float64(len(resp.Data.Avg)), " ",lastValue/float64(len(resp.Data.Avg)))
		}
	}
	//WebSocketClient()
	//go Get(url2,"127.0.0.1")
	//go Get(url3,"localhost")
	//select {}



}
func Get(url string, host string){
	method := "GET"
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Authorization","Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzYwMzI4MjIsImlzcyI6ImFwcF9pZF9hIn0.Qbt48atk5sCF1CggvMKB-5H2Hm3AxRd80f5EvVDPGxA")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 200;i++ {
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(host," ",time.Now().Minute(), " ", i, " ", string(body))
		time.Sleep(time.Second)
	}
}

func WebSocketClient() {
	fmt.Println("Client started")
	for {
		conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://127.0.0.1:8888/v1/websocket")
		if err != nil {
			fmt.Println("Cannot connect: " + err.Error())
			time.Sleep(time.Duration(5) * time.Second)
			continue
		}
		fmt.Println("Connected to server")
		for i := 0; i < 10; i++ {
			randomNumber := strconv.Itoa(rand.Intn(100))
			msg := []byte(randomNumber)
			err = wsutil.WriteClientMessage(conn, ws.OpText, msg)
			if err != nil {
				fmt.Println("Cannot send: " + err.Error())
				continue
			}
			fmt.Println("Client message send with random number " + randomNumber)
			msg, _, err := wsutil.ReadServerData(conn)
			if err != nil {
				fmt.Println("Cannot receive data: " + err.Error())
				continue
			}
			fmt.Println("Server message received with random number: " + string(msg))
			time.Sleep(time.Duration(5) * time.Second)
		}
		err = conn.Close()
		if err != nil {
			fmt.Println("Cannot close the connection: " + err.Error())
			os.Exit(1)
		}
		fmt.Println("Disconnected from server")
	}
}