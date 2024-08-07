package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	Client struct {
		MessageBuf    int64
		ProducerCache int64
		ConsumerCache int64
		Upgrade       struct {
			HandshakeTimeout int64
			ReadBufferSize   int64
			WriteBufferSize  int64
		}
	}
	GroupMsgKqConf kq.KqConf
	log            logx.LogConf
	UserRpcService zrpc.RpcClientConf
	NewUserKqConf  kq.KqConf

	//redis
	Redis struct {
		Host string
		Type string
		Pass string
	}
}
