package paddleOCRclient

import (
	"C-OCR/paddleOCRclient/proto"
	"context"
	"fmt"
	"github.com/heejinzzz/consulTool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var consulAddress = "127.0.0.1:8500"
var Client proto.C_OCRClient

func init() {
	InitPaddleOCRclient()
}

func InitPaddleOCRclient() {
	serverAddr := consulTool.GetServerAddress(consulAddress, "paddleOCR")

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	Client = proto.NewC_OCRClient(conn)
}

// DoOCR 对传入的图片进行OCR文字识别，将识别出的文字内容和结果图片返回
func DoOCR(imageBytes []byte, imageType string) (resultText string, resultImageBytes []byte) {
	req := &proto.Request{
		ImageType:  imageType,
		ImageBytes: imageBytes,
	}
	res, err := Client.Ocr(context.Background(), req)
	if err != nil {
		InitPaddleOCRclient()
		res, err = Client.Ocr(context.Background(), req)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	resultText = res.Text
	resultImageBytes = res.ImageBytes
	return
}
