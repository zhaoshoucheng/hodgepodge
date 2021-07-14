module github.com/zhaoshoucheng/hodgepodge

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.8.0 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace gorm.io/gorm v1.21.10 => github.com/go-gorm/gorm v1.21.11

replace github.com/go-gorm/gorm v1.21.11 => gorm.io/gorm v1.21.10
