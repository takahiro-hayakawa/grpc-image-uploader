package main

import (
	"bytes"
	"context"
	"google.golang.org/grpc"
	"image/upload/gen"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
)

func readN(r io.Reader, n int) []byte {
	buf := make([]byte, n)
	cnt, err := r.Read(buf)
	if err != nil {
		return []byte{}
	}
	if n != cnt {
		return buf[0:cnt]
	}
	return buf
}

func chunkedUpload(c gen.ImageUploadServiceClient, filePath string) {
	fileName := filepath.Base(filePath)

	log.Println("upload start")
	log.Printf("sent name=%v\n", fileName)

	stream, err := c.Upload(context.Background())
	if err != nil {
		log.Fatalf("error while calling upload: %v", err)
	}

	err = stream.Send(&gen.ImageUploadRequest{File: &gen.ImageUploadRequest_FileMeta_{FileMeta: &gen.ImageUploadRequest_FileMeta{FileName: "nekochan.jpg"}}})
	if err != nil {
		log.Fatalf("error while calling upload: %v", err)
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}
	r := bytes.NewReader(buf)
	for {
		data := readN(r, 1024*100)
		dataSize := len(data)
		if dataSize <= 0 {
			break
		}
		err = stream.Send(&gen.ImageUploadRequest{File: &gen.ImageUploadRequest_Data{Data: data}})
		if err != nil {
			log.Fatalf("error while sent image: %v", err)
		}
		log.Printf("sent %v\n", dataSize)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error close: %v", err)
	}

	log.Printf("Response from Server: %v\n", res)
}

func main() {
	var opts []grpc.DialOption

	// セキュア通信設定
	tls := false
	if tls {
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	cc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	filePath, _ := filepath.Abs("client/img/nekochan.jpg")
	c := gen.NewImageUploadServiceClient(cc)
	chunkedUpload(c, filePath)
}
