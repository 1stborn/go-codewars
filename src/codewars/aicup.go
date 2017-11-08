package codewars

import (
	"bufio"
	"encoding/binary"
	"errors"
	"net"
)

var ByteOrder = binary.LittleEndian
var ErrWrongType = errors.New("wrong message type")

type AiCup struct {
	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
}

func (c *AiCup) Dial(host, port string) (err error) {
	if c.conn, err = net.Dial("tcp", host+":"+port); err == nil {
		c.reader = bufio.NewReader(c.conn)
		c.writer = bufio.NewWriter(c.conn)
	}

	return
}

func (c *AiCup) writeToken(token string) {
	c.writeOpcode(AuthenticationToken)
	c.writeString(token)
	c.flush()
}

func (c *AiCup) writeProtoVersion(ver int) {
	c.writeOpcode(ProtocolVersion)
	c.writeInt(ver)
	c.flush()
}

func (c *AiCup) ReadTeamSize() int {
	c.ensureMessageType(TeamSize)
	return c.readInt()
}

func (c *AiCup) Close() error {
	return c.conn.Close()
}

func (c *AiCup) readOpcode() MessageType {
	return MessageType(c.readByte())
}

func (c *AiCup) readIntArray() []int {
	var arr []int
	if ln := c.readInt(); ln > 0 {
		for ; ln >= 0; ln-- {
			arr = append(arr, c.readInt())
		}
	}
	return arr
}

func (c *AiCup) readInt() int {
	var v int32
	if err := binary.Read(c.reader, ByteOrder, &v); err != nil {
		panic(err)
	}
	return int(v)
}

func (c *AiCup) readInt64() int64 {
	var v int64
	if err := binary.Read(c.reader, ByteOrder, &v); err != nil {
		panic(err)
	}
	return v
}

func (c *AiCup) readFloat64() float64 {
	var v float64
	if err := binary.Read(c.reader, ByteOrder, &v); err != nil {
		panic(err)
	}
	return v
}

func (c *AiCup) writeBool(b bool) {
	if b {
		c.writeByte(1)
	} else {
		c.writeByte(0)
	}
}

func (c *AiCup) readBool() bool {
	return c.readByte() != 0
}

func (c *AiCup) readByte() byte {
	b, err := c.reader.ReadByte()
	if err != nil {
		panic(err)
	}
	return b
}

func (c *AiCup) readString() string {
	return string(c.readBytes())
}

func (c *AiCup) ensureMessageType(m MessageType) {
	if b, err := c.reader.ReadByte(); err != nil || b != byte(m) {
		panic(ErrWrongType)
	}
}

func (c *AiCup) writeOpcode(m MessageType) {
	c.writeByte(byte(m))
}

func (c *AiCup) writeInt(v int) {
	if err := binary.Write(c.writer, ByteOrder, int32(v)); err != nil {
		panic(err)
	}
}

func (c *AiCup) writeFloat64(v float64) {
	if err := binary.Write(c.writer, ByteOrder, v); err != nil {
		panic(err)
	}
}

func (c *AiCup) writeInt64(v int64) {
	if err := binary.Write(c.writer, ByteOrder, v); err != nil {
		panic(err)
	}
}

func (c *AiCup) readBytes() []byte {
	l := c.readInt()
	r := make([]byte, l)
	for i := range r {
		r[i] = c.readByte()
	}
	return r
}

func (c *AiCup) writeByte(v byte) {
	if err := c.writer.WriteByte(v); err != nil {
		panic(err)
	}
}

func (c *AiCup) writeBytes(v []byte) {
	c.writeInt(len(v))
	if _, err := c.writer.Write(v); err != nil {
		panic(err)
	}
}

func (c *AiCup) writeString(v string) {
	c.writeInt(len(v))
	if _, err := c.writer.WriteString(v); err != nil {
		panic(err)
	}
}

func (c *AiCup) flush() {
	if err := c.writer.Flush(); err != nil {
		panic(err)
	}
}
