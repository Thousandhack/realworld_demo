package server

import (
	"log"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"

	"realworld_demo/internal/errors"
)

func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	// 记录错误
	log.Printf("API错误: %v", err)

	// 转换为HTTP错误
	se := errors.FromError(err)

	// 获取编解码器
	codec, _ := http.CodecForRequest(r, "Accept")

	// 打印错误详情
	log.Printf("HTTP错误: 状态码=%d, 错误=%v", se.Code, se.Errors)

	// 序列化错误
	body, err := codec.Marshal(se)
	if err != nil {
		log.Printf("错误序列化失败: %v", err)
		w.WriteHeader(500)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/"+codec.Name())

	// 设置状态码
	if se.Code > 99 && se.Code < 600 {
		w.WriteHeader(se.Code)
	} else {
		w.WriteHeader(500)
	}

	// 写入响应体
	_, _ = w.Write(body)

	// 打印响应体
	log.Printf("错误响应体: %s", string(body))
}
