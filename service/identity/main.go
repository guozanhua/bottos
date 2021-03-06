package main

import (
	"log"
	"github.com/micro/go-micro"
	 proto "./proto"
	"golang.org/x/net/context"
	"github.com/emicklei/go-restful"
)

type User struct{}

func (u *User) Register(ctx context.Context, req *proto.RegisterRequest, rsp *proto.RegisterResponse) error {
	rsp.Code = 1
	rsp.Msg = "OK"
	return nil
}

func (u *User) Login(ctx context.Context, req *proto.LoginRequest, rsp *proto.LoginResponse) error {
	rsp.Code = 1
	rsp.Msg = "OK"
	rsp.Token = req.Username + req.Signature
	return nil
}


func WriteChain(req *restful.Request, rsp *restful.Response) error {
	log.Print("Received Say.Anything API request")
	rsp.WriteEntity(map[string]string{
		"message": "Hi, this is the Greeter API",
	})
	return nil
}

func GetBolckNum() (int, int, error) {
	req := curl.NewRequest()
	resp, err := req.SetUrl(GET_INFO_URL).Get()
	if err != nil {
		return 0, resp.Raw.StatusCode, err
	}
	if resp.IsOk() {
		js, _ := simplejson.NewJson([]byte(resp.Body))
		block_num := js.Get("head_block_num").MustInt()
		return block_num, resp.Raw.StatusCode, err
	} else {
		return 0, resp.Raw.StatusCode, err
	}
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("2.0.0"),
	)

	service.Init()

	proto.RegisterUserHandler(service.Server(), new(User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
