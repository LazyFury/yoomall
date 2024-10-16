// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"yoomall/apps/auth"
	"yoomall/apps/auth/handler"
	"yoomall/apps/auth/middleware"
	"yoomall/apps/auth/service"
	"yoomall/config"
	"yoomall/core/http"
)

// Injectors from wire.go:

func NewApp() httpserver.HttpServer {
	viper := config.NewConfig()
	db := NewDB(viper)
	authService := authservice.NewAuthService(db)
	authMiddlewareGroup := authmiddleware.NewAuthMiddlewareGroup(db)
	userHandler := handler.NewUserHandler(db, viper, authService, authMiddlewareGroup)
	userRoleHandler := handler.NewUserRoleHandler(db, authMiddlewareGroup)
	userTokenHandler := handler.NewUserTokenHandler(db, authMiddlewareGroup)
	permissionHandler := handler.NewPermissionHandler(db, authMiddlewareGroup)
	authApp := auth.NewAuthApp(viper, db, authService, userHandler, userRoleHandler, userTokenHandler, permissionHandler)
	httpServer := NewHttpServer(viper, authApp)
	return httpServer
}
