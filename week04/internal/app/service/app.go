package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go-practice-homework/week04/internal/app/biz"
	"strconv"
)
import pb "go-practice-homework/week04/api/app/v1"

type AppService struct {
	pb.UnimplementedAppServiceServer
	router *gin.Engine
	data   *biz.DataUsecase
}

var ProviderSet = wire.NewSet(NewAppService)

func NewAppService(gin *gin.Engine, data *biz.DataUsecase) *AppService {
	return &AppService{router: gin, data: data}
}

// 路由注册，首先需要 gin.IRouter 接口用于注册
// 其次需要获取到 SayHello 方法对应的 http method 和 path
func (s *AppService) RegisterService() {
	s.router.Handle("GET", "/data", s.GinGetData)
}

func (s *AppService) GinGetData(g *gin.Context) {
	var in pb.GetDataRequest
	in.Id, _ = strconv.ParseInt(g.Query("id"), 10, 64)

	out, err := s.GetData(g, &in)
	if err != nil {
		g.JSON(404, "Data not found.")
		return
	}

	g.JSON(200, out.String())
	return
}

func (s *AppService) GetData(ctx context.Context, req *pb.GetDataRequest) (*pb.GetDataReply, error) {
	d, err := s.data.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetDataReply{Data: &pb.Data{Id: d.Id, Data: d.Data}}, nil
}
