package initialize

func Run() {
	LoadConfig()
	InitLogger()
	InitMySql()
	InitRedis()
	InitKafka()

	r := InitRouter()
	r.Run(":8888")
}
