module github.com/zhaoshoucheng/hodgepodge

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/HdrHistogram/hdrhistogram-go v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/go-lintpack/lintpack v0.5.2 // indirect
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golangci/errcheck v0.0.0-20181223084120-ef45e06d44b6 // indirect
	github.com/golangci/go-tools v0.0.0-20190318055746-e32c54105b7c // indirect
	github.com/golangci/goconst v0.0.0-20180610141641-041c5f2b40f3 // indirect
	github.com/golangci/gocyclo v0.0.0-20180528134321-2becd97e67ee // indirect
	github.com/golangci/golangci-lint v1.42.1 // indirect
	github.com/golangci/gosec v0.0.0-20190211064107-66fb7fc33547 // indirect
	github.com/golangci/ineffassign v0.0.0-20190609212857-42439a7714cc // indirect
	github.com/golangci/prealloc v0.0.0-20180630174525-215b22d4de21 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/klauspost/cpuid v1.2.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/shirou/gopsutil v0.0.0-20180427012116-c95755e4bcd7 // indirect
	github.com/shirou/w32 v0.0.0-20160930032740-bb4de0191aa4 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.8.0 // indirect
	go.uber.org/dig v1.11.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
	sourcegraph.com/sqs/pbtypes v0.0.0-20180604144634-d3ebe8f20ae4 // indirect
)

replace gorm.io/gorm v1.21.10 => github.com/go-gorm/gorm v1.21.11

replace github.com/go-gorm/gorm v1.21.11 => gorm.io/gorm v1.21.10
