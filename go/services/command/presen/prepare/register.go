package prepare

import (
	"crypto/tls"
	"embed"

	"github.com/grpc-tutorial/go/services/command/proto/pb"

	"io/fs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var content embed.FS

type CommandServer struct {
	Server *grpc.Server
}

func loadPem() credentials.TransportCredentials {
	cert, _ := fs.ReadFile(content, "commandservice.pem")
	key, _ := fs.ReadFile(content, "commandservice-key.pem")

	if certificate, err := tls.X509KeyPair(cert, key); err != nil {
		panic(err)
	} else {
		creds := credentials.NewServerTLSFromCert(&certificate)
		return creds
	}
}

func NewCommandServer(category pb.CategoryCommandServer, product pb.ProductCommandServer) *CommandServer {
	//grpc serverを作成する
	server := grpc.NewServer(
		grpc.Creds(loadPem()),
	)

	pb.RegisterCategoryCommandServer(server, category)
	pb.RegisterProductCommandServer(server, product)
	return &CommandServer{Server: server}
}
