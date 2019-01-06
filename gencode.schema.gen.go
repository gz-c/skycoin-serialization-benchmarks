package serializebench

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

type GencodeSignedBlock struct {
	Sig   [65]byte
	Block GencodeBlock
}

func (d *GencodeSignedBlock) Size() (s uint64) {

	{
		s += 65
	}
	{
		s += d.Block.Size()
	}
	return
}
func (d *GencodeSignedBlock) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		copy(buf[i+0:], d.Sig[:])
		i += 65
	}
	{
		nbuf, err := d.Block.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *GencodeSignedBlock) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		copy(d.Sig[:], buf[i+0:])
		i += 65
	}
	{
		ni, err := d.Block.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type GencodeBlock struct {
	Head GencodeBlockHeader
	Body GencodeBlockBody
}

func (d *GencodeBlock) Size() (s uint64) {

	{
		s += d.Head.Size()
	}
	{
		s += d.Body.Size()
	}
	return
}
func (d *GencodeBlock) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		nbuf, err := d.Head.Marshal(buf[0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		nbuf, err := d.Body.Marshal(buf[i+0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	return buf[:i+0], nil
}

func (d *GencodeBlock) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		ni, err := d.Head.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		ni, err := d.Body.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	return i + 0, nil
}

type GencodeBlockHeader struct {
	Version  uint32
	Time     uint64
	BkSeq    uint64
	Fee      uint64
	PrevHash [32]byte
	BodyHash [32]byte
	UxHash   [32]byte
}

func (d *GencodeBlockHeader) Size() (s uint64) {

	{
		s += 32
	}
	{
		s += 32
	}
	{
		s += 32
	}
	s += 28
	return
}
func (d *GencodeBlockHeader) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Version >> 0)

		buf[1+0] = byte(d.Version >> 8)

		buf[2+0] = byte(d.Version >> 16)

		buf[3+0] = byte(d.Version >> 24)

	}
	{

		buf[0+4] = byte(d.Time >> 0)

		buf[1+4] = byte(d.Time >> 8)

		buf[2+4] = byte(d.Time >> 16)

		buf[3+4] = byte(d.Time >> 24)

		buf[4+4] = byte(d.Time >> 32)

		buf[5+4] = byte(d.Time >> 40)

		buf[6+4] = byte(d.Time >> 48)

		buf[7+4] = byte(d.Time >> 56)

	}
	{

		buf[0+12] = byte(d.BkSeq >> 0)

		buf[1+12] = byte(d.BkSeq >> 8)

		buf[2+12] = byte(d.BkSeq >> 16)

		buf[3+12] = byte(d.BkSeq >> 24)

		buf[4+12] = byte(d.BkSeq >> 32)

		buf[5+12] = byte(d.BkSeq >> 40)

		buf[6+12] = byte(d.BkSeq >> 48)

		buf[7+12] = byte(d.BkSeq >> 56)

	}
	{

		buf[0+20] = byte(d.Fee >> 0)

		buf[1+20] = byte(d.Fee >> 8)

		buf[2+20] = byte(d.Fee >> 16)

		buf[3+20] = byte(d.Fee >> 24)

		buf[4+20] = byte(d.Fee >> 32)

		buf[5+20] = byte(d.Fee >> 40)

		buf[6+20] = byte(d.Fee >> 48)

		buf[7+20] = byte(d.Fee >> 56)

	}
	{
		copy(buf[i+28:], d.PrevHash[:])
		i += 32
	}
	{
		copy(buf[i+28:], d.BodyHash[:])
		i += 32
	}
	{
		copy(buf[i+28:], d.UxHash[:])
		i += 32
	}
	return buf[:i+28], nil
}

func (d *GencodeBlockHeader) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Version = 0 | (uint32(buf[i+0+0]) << 0) | (uint32(buf[i+1+0]) << 8) | (uint32(buf[i+2+0]) << 16) | (uint32(buf[i+3+0]) << 24)

	}
	{

		d.Time = 0 | (uint64(buf[i+0+4]) << 0) | (uint64(buf[i+1+4]) << 8) | (uint64(buf[i+2+4]) << 16) | (uint64(buf[i+3+4]) << 24) | (uint64(buf[i+4+4]) << 32) | (uint64(buf[i+5+4]) << 40) | (uint64(buf[i+6+4]) << 48) | (uint64(buf[i+7+4]) << 56)

	}
	{

		d.BkSeq = 0 | (uint64(buf[i+0+12]) << 0) | (uint64(buf[i+1+12]) << 8) | (uint64(buf[i+2+12]) << 16) | (uint64(buf[i+3+12]) << 24) | (uint64(buf[i+4+12]) << 32) | (uint64(buf[i+5+12]) << 40) | (uint64(buf[i+6+12]) << 48) | (uint64(buf[i+7+12]) << 56)

	}
	{

		d.Fee = 0 | (uint64(buf[i+0+20]) << 0) | (uint64(buf[i+1+20]) << 8) | (uint64(buf[i+2+20]) << 16) | (uint64(buf[i+3+20]) << 24) | (uint64(buf[i+4+20]) << 32) | (uint64(buf[i+5+20]) << 40) | (uint64(buf[i+6+20]) << 48) | (uint64(buf[i+7+20]) << 56)

	}
	{
		copy(d.PrevHash[:], buf[i+28:])
		i += 32
	}
	{
		copy(d.BodyHash[:], buf[i+28:])
		i += 32
	}
	{
		copy(d.UxHash[:], buf[i+28:])
		i += 32
	}
	return i + 28, nil
}

type GencodeBlockBody struct {
	Transactions []GencodeTransaction
}

func (d *GencodeBlockBody) Size() (s uint64) {

	{
		l := uint64(len(d.Transactions))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Transactions {

			{
				s += d.Transactions[k0].Size()
			}

		}

	}
	return
}
func (d *GencodeBlockBody) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Transactions))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Transactions {

			{
				nbuf, err := d.Transactions[k0].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *GencodeBlockBody) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Transactions)) >= l {
			d.Transactions = d.Transactions[:l]
		} else {
			d.Transactions = make([]GencodeTransaction, l)
		}
		for k0 := range d.Transactions {

			{
				ni, err := d.Transactions[k0].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

type GencodeTransaction struct {
	Length    uint32
	Type      uint8
	InnerHash [32]byte
	Sigs      [][65]byte
	In        [][32]byte
	Out       []GencodeTransactionOutput
}

func (d *GencodeTransaction) Size() (s uint64) {

	{
		s += 32
	}
	{
		l := uint64(len(d.Sigs))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for _ = range d.Sigs {

			{
				s += 65
			}

		}

	}
	{
		l := uint64(len(d.In))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for _ = range d.In {

			{
				s += 32
			}

		}

	}
	{
		l := uint64(len(d.Out))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Out {

			{
				s += d.Out[k0].Size()
			}

		}

	}
	s += 5
	return
}
func (d *GencodeTransaction) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Length >> 0)

		buf[1+0] = byte(d.Length >> 8)

		buf[2+0] = byte(d.Length >> 16)

		buf[3+0] = byte(d.Length >> 24)

	}
	{

		buf[0+4] = byte(d.Type >> 0)

	}
	{
		copy(buf[i+5:], d.InnerHash[:])
		i += 32
	}
	{
		l := uint64(len(d.Sigs))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+5] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+5] = byte(t)
			i++

		}
		for k0 := range d.Sigs {

			{
				copy(buf[i+5:], d.Sigs[k0][:])
				i += 65
			}

		}
	}
	{
		l := uint64(len(d.In))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+5] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+5] = byte(t)
			i++

		}
		for k0 := range d.In {

			{
				copy(buf[i+5:], d.In[k0][:])
				i += 32
			}

		}
	}
	{
		l := uint64(len(d.Out))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+5] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+5] = byte(t)
			i++

		}
		for k0 := range d.Out {

			{
				nbuf, err := d.Out[k0].Marshal(buf[i+5:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+5], nil
}

func (d *GencodeTransaction) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Length = 0 | (uint32(buf[i+0+0]) << 0) | (uint32(buf[i+1+0]) << 8) | (uint32(buf[i+2+0]) << 16) | (uint32(buf[i+3+0]) << 24)

	}
	{

		d.Type = 0 | (uint8(buf[i+0+4]) << 0)

	}
	{
		copy(d.InnerHash[:], buf[i+5:])
		i += 32
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+5] & 0x7F)
			for buf[i+5]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+5]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Sigs)) >= l {
			d.Sigs = d.Sigs[:l]
		} else {
			d.Sigs = make([][65]byte, l)
		}
		for k0 := range d.Sigs {

			{
				copy(d.Sigs[k0][:], buf[i+5:])
				i += 65
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+5] & 0x7F)
			for buf[i+5]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+5]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.In)) >= l {
			d.In = d.In[:l]
		} else {
			d.In = make([][32]byte, l)
		}
		for k0 := range d.In {

			{
				copy(d.In[k0][:], buf[i+5:])
				i += 32
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+5] & 0x7F)
			for buf[i+5]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+5]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Out)) >= l {
			d.Out = d.Out[:l]
		} else {
			d.Out = make([]GencodeTransactionOutput, l)
		}
		for k0 := range d.Out {

			{
				ni, err := d.Out[k0].Unmarshal(buf[i+5:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 5, nil
}

type GencodeTransactionOutput struct {
	Address GencodeAddress
	Coins   uint64
	Hours   uint64
}

func (d *GencodeTransactionOutput) Size() (s uint64) {

	{
		s += d.Address.Size()
	}
	s += 16
	return
}
func (d *GencodeTransactionOutput) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		nbuf, err := d.Address.Marshal(buf[0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{

		buf[i+0+0] = byte(d.Coins >> 0)

		buf[i+1+0] = byte(d.Coins >> 8)

		buf[i+2+0] = byte(d.Coins >> 16)

		buf[i+3+0] = byte(d.Coins >> 24)

		buf[i+4+0] = byte(d.Coins >> 32)

		buf[i+5+0] = byte(d.Coins >> 40)

		buf[i+6+0] = byte(d.Coins >> 48)

		buf[i+7+0] = byte(d.Coins >> 56)

	}
	{

		buf[i+0+8] = byte(d.Hours >> 0)

		buf[i+1+8] = byte(d.Hours >> 8)

		buf[i+2+8] = byte(d.Hours >> 16)

		buf[i+3+8] = byte(d.Hours >> 24)

		buf[i+4+8] = byte(d.Hours >> 32)

		buf[i+5+8] = byte(d.Hours >> 40)

		buf[i+6+8] = byte(d.Hours >> 48)

		buf[i+7+8] = byte(d.Hours >> 56)

	}
	return buf[:i+16], nil
}

func (d *GencodeTransactionOutput) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		ni, err := d.Address.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{

		d.Coins = 0 | (uint64(buf[i+0+0]) << 0) | (uint64(buf[i+1+0]) << 8) | (uint64(buf[i+2+0]) << 16) | (uint64(buf[i+3+0]) << 24) | (uint64(buf[i+4+0]) << 32) | (uint64(buf[i+5+0]) << 40) | (uint64(buf[i+6+0]) << 48) | (uint64(buf[i+7+0]) << 56)

	}
	{

		d.Hours = 0 | (uint64(buf[i+0+8]) << 0) | (uint64(buf[i+1+8]) << 8) | (uint64(buf[i+2+8]) << 16) | (uint64(buf[i+3+8]) << 24) | (uint64(buf[i+4+8]) << 32) | (uint64(buf[i+5+8]) << 40) | (uint64(buf[i+6+8]) << 48) | (uint64(buf[i+7+8]) << 56)

	}
	return i + 16, nil
}

type GencodeAddress struct {
	Version uint8
	Key     [20]byte
}

func (d *GencodeAddress) Size() (s uint64) {

	{
		s += 20
	}
	s += 1
	return
}
func (d *GencodeAddress) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Version >> 0)

	}
	{
		copy(buf[i+1:], d.Key[:])
		i += 20
	}
	return buf[:i+1], nil
}

func (d *GencodeAddress) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Version = 0 | (uint8(buf[i+0+0]) << 0)

	}
	{
		copy(d.Key[:], buf[i+1:])
		i += 20
	}
	return i + 1, nil
}
