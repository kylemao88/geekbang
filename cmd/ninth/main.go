// goim decoder
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

///big-end
type improto struct {
	packlen [4]byte
	headlen [2]byte
	version [2]byte
	opr     [4]byte
	seq     [4]byte
	body    []byte // packlen - headlen
}

type imdecoder struct {
	buff      string
	currproto improto
}

func (d *imdecoder) parse() error {
	if len(d.buff) < 16 {
		return fmt.Errorf("invalid buff, len:%d", len(d.buff))
	}

	//
	packlen := []byte(d.buff[0:4])
	d.currproto.packlen[0] = packlen[0]
	d.currproto.packlen[1] = packlen[1]
	d.currproto.packlen[2] = packlen[2]
	d.currproto.packlen[3] = packlen[3]
	packLenNum, err := bytesToIntU(packlen)
	if err != nil {
		return fmt.Errorf("conver packlen err:%+v, packlen:%s", err, string(packlen))
	}

	//
	headlen := []byte(d.buff[4:6])
	d.currproto.headlen[0] = headlen[0]
	d.currproto.headlen[1] = headlen[1]
	headLenNum, err := bytesToIntU(headlen)
	if err != nil {
		return fmt.Errorf("conver headlen err:%+v, packlen:%s", err, string(headlen))
	}

	bodyLenNum := packLenNum - headLenNum
	if bodyLenNum <= 0 || bodyLenNum < 10 {
		return fmt.Errorf("invalid body len, bodyLenNum:%d", err, bodyLenNum)
	}

	//
	version := d.buff[6:8]
	d.currproto.version[0] = version[0]
	d.currproto.version[1] = version[1]

	//
	opr := d.buff[8:12]
	d.currproto.opr[0] = opr[0]
	d.currproto.opr[1] = opr[1]
	d.currproto.opr[2] = opr[2]
	d.currproto.opr[3] = opr[3]

	//
	seq := d.buff[12:16]
	d.currproto.seq[0] = seq[0]
	d.currproto.seq[1] = seq[1]
	d.currproto.seq[2] = seq[2]
	d.currproto.seq[3] = seq[3]

	bodylen := bodyLenNum - 10
	body := d.buff[16 : 16+bodylen]
	d.currproto.body = []byte(body)
	return nil

}

func NewImDecoder(buff string) imdecoder {
	decoder := imdecoder{buff: buff}
	return decoder
}

//  处理粘包
func getBuffFromSocket() string {
	var buff string = ""
	// ....
	return buff
}

//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	// 3字节，首位补0处理成4字节
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	//
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

func main() {

	// 从socket上获取一段流数据
	buff := getBuffFromSocket()
	decoder := NewImDecoder(buff)
	if err := decoder.parse(); err != nil {
		fmt.Println("decoder parser fail : %+v", err)
	} else {
		fmt.Println("decoder parser succ, curr buff : %+v ", decoder.currproto)
	}

	return
}
