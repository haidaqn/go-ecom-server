package initialize

import (
	"fmt"
	"time"

	"github.com/haidaqn/go-ecommerce-backend-api/global"
	"github.com/haidaqn/go-ecommerce-backend-api/internal/po"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, errorString string) {
	if err != nil {
		global.Logger.Error(errorString, zap.Error(err))
		panic(err)
	}
}

func InitMySql() {
	mysqlSetting := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, mysqlSetting.User, mysqlSetting.Password, mysqlSetting.Host, mysqlSetting.Port, mysqlSetting.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
	})
	checkErrorPanic(err, "Failed to connect to MySQL database")
	global.Logger.Info("Connected to MySQL database")
	global.MySQL = db
	SetPool()
	genTablesDAO()
	// MigrateTables()
}

func SetPool() {
	sqlDB, err := global.MySQL.DB()
	checkErrorPanic(err, "Failed to get SQL DB from GORM")
	mysqlSetting := global.Config.MySQL

	// giới hạn số lượng connection idle được giữ lại
	sqlDB.SetMaxIdleConns(mysqlSetting.MaxIdleConns)
	// giới hạn số lượng connection mở tối đa
	sqlDB.SetMaxOpenConns(mysqlSetting.MaxOpenConns)
	// giới hạn thời gian sống của connection tối đa
	sqlDB.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifetime) * time.Second)
	global.Logger.Info("MySQL connection pool settings applied")
}

func MigrateTables() {
	err := global.MySQL.AutoMigrate(&po.User{}, &po.Role{})
	checkErrorPanic(err, "Failed to migrate tables")
	global.Logger.Info("Tables migrated successfully")
}

func genTablesDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(global.MySQL) // reuse your gorm db
	g.GenerateModel("users")
	g.Execute()

}
