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

type GencodeVarintSignedBlock struct {
	Sig   [65]byte
	Block GencodeVarintBlock
}

func (d *GencodeVarintSignedBlock) Size() (s uint64) {

	{
		s += 65
	}
	{
		s += d.Block.Size()
	}
	return
}
func (d *GencodeVarintSignedBlock) Marshal(buf []byte) ([]byte, error) {
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

func (d *GencodeVarintSignedBlock) Unmarshal(buf []byte) (uint64, error) {
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

type GencodeVarintBlock struct {
	Head GencodeVarintBlockHeader
	Body GencodeVarintBlockBody
}

func (d *GencodeVarintBlock) Size() (s uint64) {

	{
		s += d.Head.Size()
	}
	{
		s += d.Body.Size()
	}
	return
}
func (d *GencodeVarintBlock) Marshal(buf []byte) ([]byte, error) {
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

func (d *GencodeVarintBlock) Unmarshal(buf []byte) (uint64, error) {
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

type GencodeVarintBlockHeader struct {
	Version  uint32
	Time     uint64
	BkSeq    uint64
	Fee      uint64
	PrevHash [32]byte
	BodyHash [32]byte
	UxHash   [32]byte
}

func (d *GencodeVarintBlockHeader) Size() (s uint64) {

	{

		t := d.Version
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Time
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.BkSeq
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Fee
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{
		s += 32
	}
	{
		s += 32
	}
	{
		s += 32
	}
	return
}
func (d *GencodeVarintBlockHeader) Marshal(buf []byte) ([]byte, error) {
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

		t := uint32(d.Version)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint64(d.Time)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint64(d.BkSeq)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint64(d.Fee)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{
		copy(buf[i+0:], d.PrevHash[:])
		i += 32
	}
	{
		copy(buf[i+0:], d.BodyHash[:])
		i += 32
	}
	{
		copy(buf[i+0:], d.UxHash[:])
		i += 32
	}
	return buf[:i+0], nil
}

func (d *GencodeVarintBlockHeader) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		bs := uint8(7)
		t := uint32(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint32(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Version = t

	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Time = t

	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.BkSeq = t

	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Fee = t

	}
	{
		copy(d.PrevHash[:], buf[i+0:])
		i += 32
	}
	{
		copy(d.BodyHash[:], buf[i+0:])
		i += 32
	}
	{
		copy(d.UxHash[:], buf[i+0:])
		i += 32
	}
	return i + 0, nil
}

type GencodeVarintBlockBody struct {
	Transactions []GencodeVarintTransaction
}

func (d *GencodeVarintBlockBody) Size() (s uint64) {

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
func (d *GencodeVarintBlockBody) Marshal(buf []byte) ([]byte, error) {
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

func (d *GencodeVarintBlockBody) Unmarshal(buf []byte) (uint64, error) {
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
			d.Transactions = make([]GencodeVarintTransaction, l)
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

type GencodeVarintTransaction struct {
	Length    uint32
	Type      uint8
	InnerHash [32]byte
	Sigs      [][65]byte
	In        [][32]byte
	Out       []GencodeVarintTransactionOutput
}

func (d *GencodeVarintTransaction) Size() (s uint64) {

	{

		t := d.Length
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Type
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
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
	return
}
func (d *GencodeVarintTransaction) Marshal(buf []byte) ([]byte, error) {
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

		t := uint32(d.Length)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint8(d.Type)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{
		copy(buf[i+0:], d.InnerHash[:])
		i += 32
	}
	{
		l := uint64(len(d.Sigs))

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
		for k0 := range d.Sigs {

			{
				copy(buf[i+0:], d.Sigs[k0][:])
				i += 65
			}

		}
	}
	{
		l := uint64(len(d.In))

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
		for k0 := range d.In {

			{
				copy(buf[i+0:], d.In[k0][:])
				i += 32
			}

		}
	}
	{
		l := uint64(len(d.Out))

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
		for k0 := range d.Out {

			{
				nbuf, err := d.Out[k0].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *GencodeVarintTransaction) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		bs := uint8(7)
		t := uint32(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint32(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Length = t

	}
	{

		bs := uint8(7)
		t := uint8(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint8(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Type = t

	}
	{
		copy(d.InnerHash[:], buf[i+0:])
		i += 32
	}
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
		if uint64(cap(d.Sigs)) >= l {
			d.Sigs = d.Sigs[:l]
		} else {
			d.Sigs = make([][65]byte, l)
		}
		for k0 := range d.Sigs {

			{
				copy(d.Sigs[k0][:], buf[i+0:])
				i += 65
			}

		}
	}
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
		if uint64(cap(d.In)) >= l {
			d.In = d.In[:l]
		} else {
			d.In = make([][32]byte, l)
		}
		for k0 := range d.In {

			{
				copy(d.In[k0][:], buf[i+0:])
				i += 32
			}

		}
	}
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
		if uint64(cap(d.Out)) >= l {
			d.Out = d.Out[:l]
		} else {
			d.Out = make([]GencodeVarintTransactionOutput, l)
		}
		for k0 := range d.Out {

			{
				ni, err := d.Out[k0].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

type GencodeVarintTransactionOutput struct {
	Address GencodeVarintAddress
	Coins   uint64
	Hours   uint64
}

func (d *GencodeVarintTransactionOutput) Size() (s uint64) {

	{
		s += d.Address.Size()
	}
	{

		t := d.Coins
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{

		t := d.Hours
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	return
}
func (d *GencodeVarintTransactionOutput) Marshal(buf []byte) ([]byte, error) {
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

		t := uint64(d.Coins)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{

		t := uint64(d.Hours)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	return buf[:i+0], nil
}

func (d *GencodeVarintTransactionOutput) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		ni, err := d.Address.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Coins = t

	}
	{

		bs := uint8(7)
		t := uint64(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint64(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Hours = t

	}
	return i + 0, nil
}

type GencodeVarintAddress struct {
	Version uint8
	Key     [20]byte
}

func (d *GencodeVarintAddress) Size() (s uint64) {

	{

		t := d.Version
		for t >= 0x80 {
			t >>= 7
			s++
		}
		s++

	}
	{
		s += 20
	}
	return
}
func (d *GencodeVarintAddress) Marshal(buf []byte) ([]byte, error) {
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

		t := uint8(d.Version)

		for t >= 0x80 {
			buf[i+0] = byte(t) | 0x80
			t >>= 7
			i++
		}
		buf[i+0] = byte(t)
		i++

	}
	{
		copy(buf[i+0:], d.Key[:])
		i += 20
	}
	return buf[:i+0], nil
}

func (d *GencodeVarintAddress) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		bs := uint8(7)
		t := uint8(buf[i+0] & 0x7F)
		for buf[i+0]&0x80 == 0x80 {
			i++
			t |= uint8(buf[i+0]&0x7F) << bs
			bs += 7
		}
		i++

		d.Version = t

	}
	{
		copy(d.Key[:], buf[i+0:])
		i += 20
	}
	return i + 0, nil
}
