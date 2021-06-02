package gorm

//mysql 测试表
/*
CREATE TABLE `its_mapdata`.`base_city` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
	`city_id` INT(11) NOT NULL COMMENT '城市ID',
	`province_id` INT(11) NOT NULL COMMENT '省份ID',
	`name` VARCHAR(20) NOT NULL COMMENT '城市名',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建起始时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
) ENGINE = InnoDB COMMENT '基础城市表';

CREATE TABLE `its_mapdata`.`base_province` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
	`province_id` INT(11) NOT NULL COMMENT '省份ID',
	`name` VARCHAR(20) NOT NULL COMMENT '省份名',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建起始时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
) ENGINE = InnoDB COMMENT '基础省份表';

CREATE TABLE `its_mapdata`.`base_districts` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增id',
	`districts_id` INT(11) NOT NULL COMMENT '行政区ID',
	`name` VARCHAR(20) NOT NULL COMMENT '行政区名',
	`city_id` INT(11) COMMENT '城市ID',
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建起始时间',
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
) ENGINE = InnoDB COMMENT '基础行政区表';
*/

var CreatCity = "CREATE TABLE base_city ( `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增id', `city_id` INT(11) NOT NULL COMMENT '城市ID', `province_id` INT(11) COMMENT '省份ID', `name` VARCHAR(100) NOT NULL COMMENT '城市名', `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建起始时间', `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' , PRIMARY KEY (`id`)) ENGINE = InnoDB COMMENT '基础城市表';"
var CreatProvince = "CREATE TABLE base_province ( `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增id', `province_id` INT(11) NOT NULL COMMENT '省份ID', `name` VARCHAR(100) NOT NULL COMMENT '省份名', `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建起始时间', `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' , PRIMARY KEY (`id`)) ENGINE = InnoDB COMMENT '基础省份表';"
var CreatDistrices = "CREATE TABLE base_districts ( `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键自增id', `districts_id` INT(11) NOT NULL COMMENT '行政区ID', `name` VARCHAR(100) NOT NULL COMMENT '行政区名', `city_id` INT(11) COMMENT '城市ID', `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建起始时间', `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间' , PRIMARY KEY (`id`)) ENGINE = InnoDB COMMENT '基础行政区表';"

type BaseCity struct {
	CityID     int    `json:"city_id" orm:"column(city_id)`
	ProvinceID int    `json:"province_id" orm:"column(province_id)`
	Name       string `json:"name" orm:"column(city_name)`
}

type BaseProvince struct {
	ProvinceID int    `json:"province_id" orm:"column(province_id)`
	Name       string `json:"name" orm:"column(province_name)`
}

type BaseDistricts struct {
	DistrictsId int    `json:"districts_id" orm:"column(districts_id)`
	Name        string `json:"name" orm:"column(districts_name)`
	CityID      int    `json:"city_id" orm:"column(city_id)`
}

type MySQLConf struct {
	DriverName      string `mapstructure:"driver_name"`
	DataSourceName  string `mapstructure:"data_source_name"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

func InitMysqlConf() *MySQLConf {
	return &MySQLConf{
		DriverName:     "mysql",
		DataSourceName:  "root:didi0124@tcp(127.0.0.1:3306)/diditest1",
		MaxOpenConn:     20,
		MaxIdleConn:     10,
		MaxConnLifeTime: 100,
	}
}
