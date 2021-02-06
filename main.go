package main

import "bitbucket.it.ittreasury.com/gitops/demo/server"

func main() {
	server.NewGinServer().Run()
}
