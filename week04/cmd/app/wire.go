//+build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-practice-homework/week04/internal/app/biz"
	"go-practice-homework/week04/internal/app/data"
	"go-practice-homework/week04/internal/app/service"
)

//go:generate wire .
func initApp(*gin.Engine) *service.AppService {
	panic(wire.Build(biz.ProviderSet, service.ProviderSet, data.NewDataRepo))
}