package main

import (
	"github.com/stark-industries/config"
	"github.com/stark-industries/pkg/router"
)

func main() {
	srv := config.New()
	router.Routes(srv)
	srv.Logger.Fatal(srv.Start(":8095"))
}

//go get k8s.io/api@kubernetes-1.12.9
//go get k8s.io/apimachinery@kubernetes-1.12.9
//go get k8s.io/client-go@kubernetes-1.12.9
