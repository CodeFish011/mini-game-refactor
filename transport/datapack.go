package transport

import (
	"bytes"
	"encoding/binary"
	"io"
	"unsafe"

	"google.golang.org/protobuf/proto"
	"sunborn.com/mini-game-refactor/pb"
)

// 定义前后端交互数据包的格式

type Datapack struct {
	ID   pb.MessageID // 标明当前数据包执行什么操作
	Len  uint32       // 标明 Data 部分的数据长度
	Data []byte       // 由 proto.Message 序列化之后生成的 []byte 填充
}

func (d *Datapack) GetHeadLen() int64 {
	return int64(unsafe.Sizeof(d.ID) + unsafe.Sizeof(d.Len))
}

// 将 ID 和 message 打包成 datapack
func Pack(id pb.MessageID, message proto.Message) (*Datapack, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		return nil, err
	}
	return &Datapack{
		ID:   id,
		Len:  uint32(len(data)),
		Data: data,
	}, nil
}

//将 message 从 datapack 中拆解出来
func (d *Datapack) Unpack(message proto.Message) error {
	err := proto.Unmarshal(d.Data, message)
	if err != nil {
		return err
	}
	return nil
}

// 将 datapack 序列化
func (d *Datapack) Marshal() ([]byte, error) {

	var err error

	writer := bytes.NewBuffer([]byte{})

	err = binary.Write(writer, binary.LittleEndian, d.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Write(writer, binary.LittleEndian, d.Len)
	if err != nil {
		return nil, err
	}

	err = binary.Write(writer, binary.LittleEndian, d.Data)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil

}

// 将 datapack 反序列化
// 此处传入的应该为 net.Conn
func Unmarshal(binaryData io.Reader) (*Datapack, error) {

	var err error

	d := &Datapack{}

	headData := make([]byte, d.GetHeadLen())

	_, err = io.ReadFull(binaryData, headData)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(headData)

	err = binary.Read(reader, binary.LittleEndian, &d.ID)
	if err != nil {
		return nil, err
	}

	err = binary.Read(reader, binary.LittleEndian, &d.Len)
	if err != nil {
		return nil, err
	}

	d.Data = make([]byte, d.Len)
	_, err = io.ReadFull(binaryData, d.Data)
	if err != nil {
		return nil, err
	}

	return d, nil
}
