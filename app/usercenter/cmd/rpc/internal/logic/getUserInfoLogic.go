package logic

import (
	"context"
	"go-zero-for-k8s/app/usercenter/cmd/rpc/usercenter"
	"go-zero-for-k8s/app/usercenter/model"
	"go-zero-for-k8s/common/xerr"
	"os"

	"go-zero-for-k8s/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-for-k8s/app/usercenter/cmd/rpc/pb"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var ErrUserNoExistsError = xerr.NewErrMsg("用户不存在")

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l.Logger.Infof("rpc-req: %+v", in)

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserInfo find user db err , id:%d , err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrUserNoExistsError, "id:%d", in.Id)
	}
	var respUser usercenter.User

	// 获取主机名并处理错误
	hostName, err := os.Hostname()
	if err != nil {
		l.Logger.Errorf("get hostname error: %v", err)
		hostName = "unknown-host"
	}
	_ = copier.Copy(&respUser, user)
	respUser.Nickname = respUser.Nickname + "  rpc ---- > > " + hostName

	return &usercenter.GetUserInfoResp{
		User: &respUser,
	}, nil

}
