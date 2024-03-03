package grpc

import (
	"fmt"
	"google.golang.org/grpc/encoding"
	"google.golang.org/protobuf/proto"
)

func Codec() encoding.Codec {
	return WithCodecParent(&protoCodec{})
}

// WithCodecParent 一个协议无感知的codec实现，返回一个encoding.Codec类型的实例；该函数尝试讲grpc消息当作raw bytes来实现，当尝试失败后，会有fallback作为一个后退的codec
func WithCodecParent(fallback encoding.Codec) encoding.Codec {
	return &rawCodec{fallback}
}

// rawCodec 自定义codec类型，实现encoding.Codec接口中的Marshal和Unmarshal
type rawCodec struct {
	parentCodec encoding.Codec // parentCodec用于当自定义Marshal和Unmarshal失败时的回退codec
}
type frame struct {
	payload []byte
}

// Marshal 序列化函数
func (r *rawCodec) Marshal(v interface{}) ([]byte, error) {
	// 尝试将消息转换为*frame类型，
	out, ok := v.(*frame)
	if !ok {
		// 若失败，则采用变量parentCodec中的Marshal进行序列化
		return r.parentCodec.Marshal(v)
	}
	return out.payload, nil
}

// Unmarshal 反序列化函数
func (r *rawCodec) Unmarshal(data []byte, v interface{}) error {
	// 尝试将消息转为*frame类型，提取出payload到[]byte, 实现反序列化
	dst, ok := v.(*frame)
	if !ok {
		// 若失败，则采用变量parentCodec中的Unmarshal进行反序列化
		return r.parentCodec.Unmarshal(data, v)
	}
	dst.payload = data
	return nil
}
func (r *rawCodec) Name() string {
	return fmt.Sprintf("proxy-grpc-%s", r.parentCodec.Name())
}

// protoCodec 实现protobuf的默认codec
type protoCodec struct{}

func (protoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}
func (protoCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}
func (protoCodec) Name() string {
	return "proto"
}
