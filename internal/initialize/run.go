package initialize

func Run() {
	LoadConfig()
	// m := global.Config.MySQL
	// fmt.Println("loaded configuration mysql:", m)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()
	r.Run(":8888")
}
