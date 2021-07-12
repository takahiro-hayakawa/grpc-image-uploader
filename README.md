# grpc-image-uploader
スターティングgRPCの第6章を単方向ストリーミングの画像アップロード機能をGoで実装

## protoファイルコンパイル
```shell
 protoc -Iproto --go_out=plugins=grpc:. proto/*
```

## 稼働確認
### サーバー
```shell
$ go run server/main.go 
2021/07/13 02:08:23 start gRPC server port: 50051

```

### クライアント
```shell
$ go run client/main.go 
2021/07/13 02:57:24 upload start
2021/07/13 02:57:24 sent name=nekochan.jpg
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 102400
2021/07/13 02:57:24 sent 61457
2021/07/13 02:57:24 Response from Server: uuid:"1f99d661-aabc-4f42-a4b7-4d0789e16ee9"  size:778257  content_type:"image/jpeg"  file_name:"nekochan.jpg"
```
