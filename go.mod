module github.com/zhaoshoucheng/hodgepodge

go 1.16

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
	gorm.io/driver/mysql v1.1.0
	gorm.io/gorm v1.21.10
)

replace gorm.io/gorm v1.21.10 => github.com/go-gorm/gorm v1.21.11

replace github.com/go-gorm/gorm v1.21.11 => gorm.io/gorm v1.21.10
