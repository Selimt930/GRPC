package main

import (
	"MailService/grpcservice"
)

func main() {
	grpcservice.NewService().InitService().Run()
}
