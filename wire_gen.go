// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/go-eagle/eagle/pkg/app"
	"github.com/go-microservice/ins-api/internal/repository"
	"github.com/go-microservice/ins-api/internal/server"
	"github.com/go-microservice/ins-api/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

func InitApp(cfg *app.Config, config *app.ServerConfig) (*app.App, error) {
	userServiceClient := repository.NewUserClient()
	userServiceServer := service.NewUserServiceServer(userServiceClient)
	httpServer := server.NewHTTPServer(config, userServiceServer)
	appApp := newApp(cfg, httpServer)
	return appApp, nil
}
