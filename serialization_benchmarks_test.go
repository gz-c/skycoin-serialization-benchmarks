package serializebench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/google/go-cmp/cmp"
	"github.com/niubaoshu/gotiny"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
	"github.com/skycoin/skycoin/src/coin"
)

//go:generate skyencoder -struct SignedBlock -no-test -package serializebench -output-path . github.com/skycoin/skycoin/src/coin

var validate = os.Getenv("VALIDATE") != ""

func getBlock() coin.SignedBlock {
	return coin.SignedBlock{
		Sig: cipher.MustSigFromHex("8cf145e9ef4a4a5254bc57798a7a61dfed238768f94edc5635175c6b91bccd8ec1555da603c5e31b018e135b82b1525be8a92973c468a74b5b40b8da189cb465eb"),
		Block: coin.Block{
			Head: coin.BlockHeader{
				Version:  1,
				Time:     1538036613,
				BkSeq:    9999999999,
				Fee:      1234123412341234,
				PrevHash: cipher.MustSHA256FromHex("59cb7d0e2ce8a03d1054afcc28a22fe864a8813460d241db38c59d10e7c29132"),
				BodyHash: cipher.MustSHA256FromHex("6d421469409591f0c3112884c8cf10f8bca5d8ab87c9c30dea2ea73b6751bbf9"),
				UxHash:   cipher.MustSHA256FromHex("6ea6a972cf06d25908b29953aeddb68c3b6f3a9903e8f964dc89b0abc0645dea"),
			},
			Body: coin.BlockBody{
				Transactions: coin.Transactions{
					{
						Length:    43214321,
						Type:      1,
						InnerHash: cipher.MustSHA256FromHex("cbedf8ef0bda91afc6a180eea0dddf8e3a986b6b6f87f70e8bffc63c6fbaa4e6"),
						Sigs: []cipher.Sig{
							cipher.MustSigFromHex("1cfd7a4db3a52a85d2a86708695112b6520acc8dc83c86e8da67915199fdf04964c168543598ab07c2b99c292899890891950364c2bf66f1aaa6d6a66a5c9a73ff"),
							cipher.MustSigFromHex("442167c6b3d13957bc32f83182c7f4fda0bb6bde893a41a6a04cdd8eecee0048d03a57eb2af04ea6050e1f418769c94c7f12fad9287dc650e6b307fdfce6b42a59"),
							cipher.MustSigFromHex("528392b4574173f4a3e024af876a0bf38bbd065ab49ac6a73b5d873adbb1d732716bb00f6e577ce3a6bc528241508a4bcdf951b0365a747ad225ef489aa5096d00"),
						},
						In: []cipher.SHA256{
							cipher.MustSHA256FromHex("536f0a1a915fadfa3a2720a0615641827ff67394d2b2149d6db63b8c619e14af"),
							cipher.MustSHA256FromHex("64ba5f01f90f97f84999f13aeaa75fed8d5b3e4a3a4a093dedf4795969e8bd27"),
							cipher.MustSHA256FromHex("810ff6dc7f524b2185a794e2ebf2a48aafe910cad4d8ce6399bea4238754b129"),
						},
						Out: []coin.TransactionOutput{
							{
								Address: cipher.MustDecodeBase58Address("23FF4fshzD8tZk2d88P22WATfzUpNQF1x85"),
								Coins:   987987987,
								Hours:   789789789,
							},
							{
								Address: cipher.MustDecodeBase58Address("29V2iRpZAqHiFZHHRqaZLArZZuTcZM5owqT"),
								Coins:   123123,
								Hours:   321321,
							},
							{
								Address: cipher.MustDecodeBase58Address("URcu42V6HCWBgQjgATSWj4uS3zW4CMRXbA"),
								Coins:   20000000,
								Hours:   46,
							},
						},
					},
					{
						Length:    98769876,
						Type:      0,
						InnerHash: cipher.MustSHA256FromHex("46856af925fde9a1652d39eea479dd92589a741451a0228402e399fae02f8f3d"),
						Sigs: []cipher.Sig{
							cipher.MustSigFromHex("92e289792200518df9a82cf9dddd1f334bf0d47fb0ed4ff70c25403f39577af5ab24ef2d02a11cf6b76e6bd017457ad60d6ca85c0567c21f5c62599c93ee98e18c"),
							cipher.MustSigFromHex("e995da86ed87640ecb44e624074ba606b781aa0cbeb24e8c27ff30becf7181175479c0d74d93fe1e8692bba628b5cf532ca80fed4135148d84e6ecc2a762a10b19"),
							cipher.MustSigFromHex("898f844f34173cd950375173255a753ea4913ba1770dbb89e74e442d8f70416877887c145efba21388eafbabcedfb5b13f01d2928922c6693fbe528c6496988601"),
						},
						In: []cipher.SHA256{
							cipher.MustSHA256FromHex("69b14a7ee184f24b95659d6887101ef7c921fa7977d95c73fbc0c4d0d22671bc"),
							cipher.MustSHA256FromHex("3a050b4ec33ec9ad2c789f24655ab1c8f7691d3a1c3d0e05cc14b022b4c360ea"),
							cipher.MustSHA256FromHex("c38cc8779a8954387a1cab782d01752abceee9b5a19bd0d56ea99055cd26464c"),
						},
						Out: []coin.TransactionOutput{
							{
								Address: cipher.MustDecodeBase58Address("XvvjeyGcTBVXDXmfJoTUseFiqHvm12C6oQ"),
								Coins:   15,
								Hours:   1237882,
							},
							{
								Address: cipher.MustDecodeBase58Address("fQXVLq9fbCC9XVxDKLAGDLXmQYjPy59j22"),
								Coins:   2102123,
								Hours:   1003,
							},
							{
								Address: cipher.MustDecodeBase58Address("24N6EW7mdDrZprUcYiCLK1GdW5K2mqy3oHp"),
								Coins:   1103000000,
								Hours:   12,
							},
						},
					},
					{
						Length:    1234,
						Type:      1,
						InnerHash: cipher.MustSHA256FromHex("dc88976edfd765531e26e218df0dbd657a6f7a6d9ebcb7479183f7bf040b21ed"),
						Sigs: []cipher.Sig{
							cipher.MustSigFromHex("7bc26cfb30c896118ceccccfeebe1981e0440a9b95832f5eb3c0b8aba076e99412915833435b6c6d3b3d39682b12fa1c2847fa29e4d6b9f4a775155f4471890201"),
							cipher.MustSigFromHex("71ff8f01e11b335ed10ca73051f514e3a3a3c46ab9d566fa274eaf9f916c4cb45eef4f05f13146837332a756f2096db9ba95291e9593c08d25ab9235266e744301"),
							cipher.MustSigFromHex("bafb6ff2a56ac731e3dde94f4c766d19df6f6d4059eac1ad3307e2d63815739e6a477d3376993bdc0a4e350ae2af0cd0d00393969c2f6bb790340e121194784000"),
						},
						In: []cipher.SHA256{
							cipher.MustSHA256FromHex("3b94215fe694f3adf613b70271c783fa4165770b64a848022a815e115659b4e7"),
							cipher.MustSHA256FromHex("a5600d0df2ffbd7234d3b94782c7aa61bba90a8edff77acbd4a3256ba751fb53"),
							cipher.MustSHA256FromHex("54bb8e61be2ae1cd143bc1e6a1ac342fbfb6c7698247eea1bff902de69fb0d1c"),
						},
						Out: []coin.TransactionOutput{
							{
								Address: cipher.MustDecodeBase58Address("URcu42V6HCWBgQjgATSWj4uS3zW4CMRXbA"),
								Coins:   108300100,
								Hours:   3499134,
							},
							{
								Address: cipher.MustDecodeBase58Address("2FkiBTThuf63qaaNB47nFmc433pzgMR2hb1"),
								Coins:   1000000000000,
								Hours:   123123995,
							},
							{
								Address: cipher.MustDecodeBase58Address("NmKtHJGGeytMFx7LLgioiYcifd743427MR"),
								Coins:   4500,
								Hours:   33342,
							},
						},
					},
				},
			},
		},
	}
}

func TestMarshaledBlockLen(t *testing.T) {
	log.SetFlags(log.LstdFlags)

	block := getBlock()

	skyBytes := encoder.Serialize(block)
	fmt.Printf("sky:\t\t\t\t %d bytes\n", len(skyBytes))

	skyEncN := EncodeSizeSignedBlock(&block)
	fmt.Printf("skyenc:\t\t\t\t %d bytes\n", skyEncN)

	var xdrBuf bytes.Buffer
	if _, err := xdr.Marshal(&xdrBuf, block); err != nil {
		t.Fatal(err)
	}
	fmt.Printf("xdr2:\t\t\t\t %d bytes\n", xdrBuf.Len())

	jsonBytes, err := json.Marshal(block)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("json:\t\t\t\t %d bytes\n", len(jsonBytes))

	gotinyEnc := gotiny.NewEncoder(coin.SignedBlock{})
	gotinyBytes := gotinyEnc.Encode(block)
	fmt.Printf("tiny:\t\t\t\t %d bytes\n", len(gotinyBytes))

	colferBlock := blockToColfer(block)
	colferBytes, err := colferBlock.MarshalBinary()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("colf:\t\t\t\t %d bytes\n", len(colferBytes))

	gencodeBlock := blockToGencode(block)
	gencodeBytes, err := gencodeBlock.Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("genc:\t\t\t\t %d bytes\n", len(gencodeBytes))

	gencodeVarintBlock := blockToGencodeVarint(block)
	gencodeVarintBytes, err := gencodeVarintBlock.Marshal(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("gencvar:\t\t\t %d bytes\n", len(gencodeVarintBytes))
}

/* sky

- Reference serializer
*/

func BenchmarkMarshalBlockBySky(b *testing.B) {
	block := getBlock()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		encoder.Serialize(block)
	}
}

func BenchmarkUnmarshalBlockBySky(b *testing.B) {
	raw := encoder.Serialize(getBlock())
	block := getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result coin.SignedBlock
		if err := encoder.DeserializeRaw(raw, &result); err != nil {
			b.Fatal(err)
		}

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("sky unmarshal result differs")
			}
		}
	}
}

/* skyencoder

- Code generator for the reference Skycoin encoder
*/

func BenchmarkMarshalBlockBySkyencoder(b *testing.B) {
	block := getBlock()
	n := EncodeSizeSignedBlock(&block)
	buf := make([]byte, n)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EncodeSignedBlock(buf, &block)
	}
}

func BenchmarkMarshalBlockBySkyencoderWithAlloc(b *testing.B) {
	block := getBlock()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		n := EncodeSizeSignedBlock(&block)
		buf := make([]byte, n)
		EncodeSignedBlock(buf, &block)
	}
}

func BenchmarkUnmarshalBlockBySkyencoder(b *testing.B) {
	block := getBlock()
	n := EncodeSizeSignedBlock(&block)
	raw := make([]byte, n)
	err := EncodeSignedBlock(raw, &block)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result coin.SignedBlock
		if x, err := DecodeSignedBlock(raw, &result); err != nil {
			b.Fatal(err)
		} else if x != len(raw) {
			b.Fatal("skyencoder: DecodeSignedBlock bytes remain")
		}

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("skyencoder unmarshal result differs")
			}
		}
	}
}

/* XDR2

- Pads everything to 4 bytes
- Big-endian
- Otherwise similar to the skycoin serializer
*/

func BenchmarkMarshalBlockByXDR2(b *testing.B) {
	block := getBlock()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var w bytes.Buffer
		if _, err := xdr.Marshal(&w, block); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByXDR2(b *testing.B) {
	var w bytes.Buffer
	if _, err := xdr.Marshal(&w, getBlock()); err != nil {
		b.Fatal(err)
	}
	byt := w.Bytes()
	block := getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result coin.SignedBlock
		if _, err := xdr.Unmarshal(bytes.NewBuffer(byt), &result); err != nil {
			b.Fatal(err)
		}

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("xdr2 unmarshal result differs")
			}
		}
	}
}

/* JSON

- Included for reference
*/

func BenchmarkMarshalBlockByJSON(b *testing.B) {
	block := getBlock()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := json.Marshal(block); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByJSON(b *testing.B) {
	raw, err := json.Marshal(getBlock())
	if err != nil {
		b.Fatal(err)
	}
	block := getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result coin.SignedBlock
		if err := json.Unmarshal(raw, &result); err != nil {
			b.Fatal(err)
		}

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("JSON unmarshal result differs")
			}
		}
	}
}

/* gotiny

- Uses unsafe.Pointer for encoding
*/

func BenchmarkMarshalBlockByGotiny(b *testing.B) {
	enc := gotiny.NewEncoder(coin.SignedBlock{})
	block := getBlock()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		enc.Encode(block)
	}
}

func BenchmarkUnmarshalBlockByGotiny(b *testing.B) {
	enc := gotiny.NewEncoder(coin.SignedBlock{})
	dec := gotiny.NewDecoder(coin.SignedBlock{})
	raw := enc.Encode(getBlock())
	block := getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var result coin.SignedBlock
		if n := dec.Decode(raw, &result); n != len(raw) {
			b.Fatal("Gotiny did not decode the whole buffer")
		}

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("gotiny unmarshal result differs")
			}
		}
	}
}

/* colfer

- Does not support fixed size arrays yet (would be more optimal)
*/

func copyBytes(b []byte) []byte {
	x := make([]byte, len(b))
	copy(x[:], b[:])
	return x
}

func blockToColfer(block coin.SignedBlock) *ColferSignedBlock {
	transactions := make([]*ColferTransaction, len(block.Block.Body.Transactions))
	for i := range block.Block.Body.Transactions {
		txn := block.Block.Body.Transactions[i]

		sigs := make([][]byte, len(txn.Sigs))
		for j, s := range txn.Sigs {
			sigs[j] = copyBytes(s[:])
		}

		in := make([][]byte, len(txn.In))
		for j, h := range txn.In {
			in[j] = copyBytes(h[:])
		}

		out := make([]*ColferTransactionOutput, len(txn.Out))
		for j, o := range txn.Out {
			out[j] = &ColferTransactionOutput{
				Address: &ColferAddress{
					Version: o.Address.Version,
					Key:     copyBytes(o.Address.Key[:]),
				},
				Coins: o.Coins,
				Hours: o.Hours,
			}
		}

		transactions[i] = &ColferTransaction{
			Length:    txn.Length,
			Type:      txn.Type,
			InnerHash: copyBytes(txn.InnerHash[:]),
			Sigs:      sigs,
			In:        in,
			Out:       out,
		}
	}

	return &ColferSignedBlock{
		Sig: block.Sig[:],
		Block: &ColferBlock{
			Head: &ColferBlockHeader{
				Version:  block.Block.Head.Version,
				Time:     block.Block.Head.Time,
				BkSeq:    block.Block.Head.BkSeq,
				Fee:      block.Block.Head.Fee,
				PrevHash: copyBytes(block.Block.Head.PrevHash[:]),
				BodyHash: copyBytes(block.Block.Head.BodyHash[:]),
				UxHash:   copyBytes(block.Block.Head.UxHash[:]),
			},
			Body: &ColferBlockBody{
				Transactions: transactions,
			},
		},
	}
}

func colferToBlock(b *ColferSignedBlock) coin.SignedBlock {
	transactions := make([]coin.Transaction, len(b.Block.Body.Transactions))
	for i := range b.Block.Body.Transactions {
		txn := b.Block.Body.Transactions[i]

		sigs := make([]cipher.Sig, len(txn.Sigs))
		for j, s := range txn.Sigs {
			sigs[j] = cipher.MustNewSig(s)
		}

		in := make([]cipher.SHA256, len(txn.In))
		for j, h := range txn.In {
			in[j] = cipher.MustSHA256FromBytes(h)
		}

		out := make([]coin.TransactionOutput, len(txn.Out))
		for j, o := range txn.Out {
			out[j] = coin.TransactionOutput{
				Address: cipher.Address{
					Version: byte(o.Address.Version),
					Key:     cipher.MustRipemd160FromBytes(o.Address.Key),
				},
				Coins: o.Coins,
				Hours: o.Hours,
			}
		}

		transactions[i] = coin.Transaction{
			Length:    txn.Length,
			Type:      txn.Type,
			InnerHash: cipher.MustSHA256FromBytes(txn.InnerHash),
			Sigs:      sigs,
			In:        in,
			Out:       out,
		}
	}

	return coin.SignedBlock{
		Sig: cipher.MustNewSig(b.Sig),
		Block: coin.Block{
			Head: coin.BlockHeader{
				Version:  b.Block.Head.Version,
				Time:     b.Block.Head.Time,
				BkSeq:    b.Block.Head.BkSeq,
				Fee:      b.Block.Head.Fee,
				PrevHash: cipher.MustSHA256FromBytes(b.Block.Head.PrevHash),
				BodyHash: cipher.MustSHA256FromBytes(b.Block.Head.BodyHash),
				UxHash:   cipher.MustSHA256FromBytes(b.Block.Head.UxHash),
			},
			Body: coin.BlockBody{
				Transactions: transactions,
			},
		},
	}
}

func BenchmarkMarshalBlockByColfer(b *testing.B) {
	block := getBlock()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		colferBlock := blockToColfer(block)
		if _, err := colferBlock.MarshalBinary(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByColfer(b *testing.B) {
	block := getBlock()
	colferBlock := blockToColfer(block)
	raw, err := colferBlock.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}
	block = getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var colferResult ColferSignedBlock
		if err := colferResult.UnmarshalBinary(raw); err != nil {
			b.Fatal(err)
		}

		result := colferToBlock(&colferResult)

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("colfer unmarshal result differs")
			}
		}
	}
}

func BenchmarkMarshalBlockByColferNoTransform(b *testing.B) {
	block := getBlock()
	colferBlock := blockToColfer(block)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := colferBlock.MarshalBinary(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByColferNoTransform(b *testing.B) {
	block := getBlock()
	colferBlock := blockToColfer(block)
	raw, err := colferBlock.MarshalBinary()
	if err != nil {
		b.Fatal(err)
	}
	block = getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var colferResult ColferSignedBlock
		if err := colferResult.UnmarshalBinary(raw); err != nil {
			b.Fatal(err)
		}

		if validate {
			result := colferToBlock(&colferResult)
			if !cmp.Equal(result, block) {
				b.Fatal("colfer unmarshal result differs")
			}
		}
	}
}

/* gencode

- generated code does not compile due to "for k0 := range txn.Sigs" - a problem with fixed size byte arrays

- this library has the best performance, but I do not trust the code, given the bugs
- the code is also hard to read
*/

func copySig(s cipher.Sig) [65]byte {
	var ss [65]byte
	copy(ss[:], s[:])
	return ss
}

func copySHA256(h cipher.SHA256) [32]byte {
	var hh [32]byte
	copy(hh[:], h[:])
	return hh
}

func copyRipemd160(h cipher.Ripemd160) [20]byte {
	var hh [20]byte
	copy(hh[:], h[:])
	return hh
}

func blockToGencode(block coin.SignedBlock) *GencodeSignedBlock {
	transactions := make([]GencodeTransaction, len(block.Block.Body.Transactions))
	for i := range block.Block.Body.Transactions {
		txn := block.Block.Body.Transactions[i]

		sigs := make([][65]byte, len(txn.Sigs))
		for j, s := range txn.Sigs {
			sigs[j] = copySig(s)
		}

		in := make([][32]byte, len(txn.In))
		for j, h := range txn.In {
			in[j] = copySHA256(h)
		}

		out := make([]GencodeTransactionOutput, len(txn.Out))
		for j, o := range txn.Out {
			out[j] = GencodeTransactionOutput{
				Address: GencodeAddress{
					Version: o.Address.Version,
					Key:     copyRipemd160(o.Address.Key),
				},
				Coins: o.Coins,
				Hours: o.Hours,
			}
		}

		transactions[i] = GencodeTransaction{
			Length:    txn.Length,
			Type:      txn.Type,
			InnerHash: copySHA256(txn.InnerHash),
			Sigs:      sigs,
			In:        in,
			Out:       out,
		}
	}

	return &GencodeSignedBlock{
		Sig: copySig(block.Sig),
		Block: GencodeBlock{
			Head: GencodeBlockHeader{
				Version:  block.Block.Head.Version,
				Time:     block.Block.Head.Time,
				BkSeq:    block.Block.Head.BkSeq,
				Fee:      block.Block.Head.Fee,
				PrevHash: copySHA256(block.Block.Head.PrevHash),
				BodyHash: copySHA256(block.Block.Head.BodyHash),
				UxHash:   copySHA256(block.Block.Head.UxHash),
			},
			Body: GencodeBlockBody{
				Transactions: transactions,
			},
		},
	}
}

func gencodeToBlock(b *GencodeSignedBlock) coin.SignedBlock {
	transactions := make([]coin.Transaction, len(b.Block.Body.Transactions))
	for i := range b.Block.Body.Transactions {
		txn := b.Block.Body.Transactions[i]

		sigs := make([]cipher.Sig, len(txn.Sigs))
		for j, s := range txn.Sigs {
			sigs[j] = cipher.MustNewSig(s[:])
		}

		in := make([]cipher.SHA256, len(txn.In))
		for j, h := range txn.In {
			in[j] = cipher.MustSHA256FromBytes(h[:])
		}

		out := make([]coin.TransactionOutput, len(txn.Out))
		for j, o := range txn.Out {
			out[j] = coin.TransactionOutput{
				Address: cipher.Address{
					Version: byte(o.Address.Version),
					Key:     cipher.MustRipemd160FromBytes(o.Address.Key[:]),
				},
				Coins: o.Coins,
				Hours: o.Hours,
			}
		}

		transactions[i] = coin.Transaction{
			Length:    txn.Length,
			Type:      txn.Type,
			InnerHash: cipher.MustSHA256FromBytes(txn.InnerHash[:]),
			Sigs:      sigs,
			In:        in,
			Out:       out,
		}
	}

	return coin.SignedBlock{
		Sig: cipher.MustNewSig(b.Sig[:]),
		Block: coin.Block{
			Head: coin.BlockHeader{
				Version:  b.Block.Head.Version,
				Time:     b.Block.Head.Time,
				BkSeq:    b.Block.Head.BkSeq,
				Fee:      b.Block.Head.Fee,
				PrevHash: cipher.MustSHA256FromBytes(b.Block.Head.PrevHash[:]),
				BodyHash: cipher.MustSHA256FromBytes(b.Block.Head.BodyHash[:]),
				UxHash:   cipher.MustSHA256FromBytes(b.Block.Head.UxHash[:]),
			},
			Body: coin.BlockBody{
				Transactions: transactions,
			},
		},
	}
}

func BenchmarkMarshalBlockByGencode(b *testing.B) {
	block := getBlock()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gencodeBlock := blockToGencode(block)
		if _, err := gencodeBlock.Marshal(nil); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByGencode(b *testing.B) {
	block := getBlock()
	gencodeBlock := blockToGencode(block)
	raw, err := gencodeBlock.Marshal(nil)
	if err != nil {
		b.Fatal(err)
	}
	block = getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var gencodeResult GencodeSignedBlock
		if _, err := gencodeResult.Unmarshal(raw); err != nil {
			b.Fatal(err)
		}

		result := gencodeToBlock(&gencodeResult)

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("gencode unmarshal result differs")
			}
		}
	}
}

func BenchmarkMarshalBlockByGencodeNoTransform(b *testing.B) {
	block := getBlock()
	gencodeBlock := blockToGencode(block)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := gencodeBlock.Marshal(nil); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByGencodeNoTransform(b *testing.B) {
	block := getBlock()
	gencodeBlock := blockToGencode(block)
	raw, err := gencodeBlock.Marshal(nil)
	if err != nil {
		b.Fatal(err)
	}
	block = getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var gencodeResult GencodeSignedBlock
		if _, err := gencodeResult.Unmarshal(raw); err != nil {
			b.Fatal(err)
		}

		if validate {
			result := gencodeToBlock(&gencodeResult)
			if !cmp.Equal(result, block) {
				b.Fatal("gencode unmarshal result differs")
			}
		}
	}
}

/* gencode with varints

- gencode schema using varints for some values (should be smaller size, but slower)

- ~ 6% smaller
- ~ 6% slower (unmarshal)
*/

func blockToGencodeVarint(block coin.SignedBlock) *GencodeVarintSignedBlock {
	transactions := make([]GencodeVarintTransaction, len(block.Block.Body.Transactions))
	for i := range block.Block.Body.Transactions {
		txn := block.Block.Body.Transactions[i]

		sigs := make([][65]byte, len(txn.Sigs))
		for j, s := range txn.Sigs {
			sigs[j] = copySig(s)
		}

		in := make([][32]byte, len(txn.In))
		for j, h := range txn.In {
			in[j] = copySHA256(h)
		}

		out := make([]GencodeVarintTransactionOutput, len(txn.Out))
		for j, o := range txn.Out {
			out[j] = GencodeVarintTransactionOutput{
				Address: GencodeVarintAddress{
					Version: o.Address.Version,
					Key:     copyRipemd160(o.Address.Key),
				},
				Coins: o.Coins,
				Hours: o.Hours,
			}
		}

		transactions[i] = GencodeVarintTransaction{
			Length:    txn.Length,
			Type:      txn.Type,
			InnerHash: copySHA256(txn.InnerHash),
			Sigs:      sigs,
			In:        in,
			Out:       out,
		}
	}

	return &GencodeVarintSignedBlock{
		Sig: copySig(block.Sig),
		Block: GencodeVarintBlock{
			Head: GencodeVarintBlockHeader{
				Version:  block.Block.Head.Version,
				Time:     block.Block.Head.Time,
				BkSeq:    block.Block.Head.BkSeq,
				Fee:      block.Block.Head.Fee,
				PrevHash: copySHA256(block.Block.Head.PrevHash),
				BodyHash: copySHA256(block.Block.Head.BodyHash),
				UxHash:   copySHA256(block.Block.Head.UxHash),
			},
			Body: GencodeVarintBlockBody{
				Transactions: transactions,
			},
		},
	}
}

func gencodeVarintToBlock(b *GencodeVarintSignedBlock) coin.SignedBlock {
	transactions := make([]coin.Transaction, len(b.Block.Body.Transactions))
	for i := range b.Block.Body.Transactions {
		txn := b.Block.Body.Transactions[i]

		sigs := make([]cipher.Sig, len(txn.Sigs))
		for j, s := range txn.Sigs {
			sigs[j] = cipher.MustNewSig(s[:])
		}

		in := make([]cipher.SHA256, len(txn.In))
		for j, h := range txn.In {
			in[j] = cipher.MustSHA256FromBytes(h[:])
		}

		out := make([]coin.TransactionOutput, len(txn.Out))
		for j, o := range txn.Out {
			out[j] = coin.TransactionOutput{
				Address: cipher.Address{
					Version: byte(o.Address.Version),
					Key:     cipher.MustRipemd160FromBytes(o.Address.Key[:]),
				},
				Coins: o.Coins,
				Hours: o.Hours,
			}
		}

		transactions[i] = coin.Transaction{
			Length:    txn.Length,
			Type:      txn.Type,
			InnerHash: cipher.MustSHA256FromBytes(txn.InnerHash[:]),
			Sigs:      sigs,
			In:        in,
			Out:       out,
		}
	}

	return coin.SignedBlock{
		Sig: cipher.MustNewSig(b.Sig[:]),
		Block: coin.Block{
			Head: coin.BlockHeader{
				Version:  b.Block.Head.Version,
				Time:     b.Block.Head.Time,
				BkSeq:    b.Block.Head.BkSeq,
				Fee:      b.Block.Head.Fee,
				PrevHash: cipher.MustSHA256FromBytes(b.Block.Head.PrevHash[:]),
				BodyHash: cipher.MustSHA256FromBytes(b.Block.Head.BodyHash[:]),
				UxHash:   cipher.MustSHA256FromBytes(b.Block.Head.UxHash[:]),
			},
			Body: coin.BlockBody{
				Transactions: transactions,
			},
		},
	}
}

func BenchmarkMarshalBlockByGencodeVarint(b *testing.B) {
	block := getBlock()
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gencodeBlock := blockToGencodeVarint(block)
		if _, err := gencodeBlock.Marshal(nil); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByGencodeVarint(b *testing.B) {
	block := getBlock()
	gencodeBlock := blockToGencodeVarint(block)
	raw, err := gencodeBlock.Marshal(nil)
	if err != nil {
		b.Fatal(err)
	}
	block = getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var gencodeResult GencodeVarintSignedBlock
		if _, err := gencodeResult.Unmarshal(raw); err != nil {
			b.Fatal(err)
		}

		result := gencodeVarintToBlock(&gencodeResult)

		if validate {
			if !cmp.Equal(result, block) {
				b.Fatal("gencode unmarshal result differs")
			}
		}
	}
}

func BenchmarkMarshalBlockByGencodeVarintNoTransform(b *testing.B) {
	block := getBlock()
	gencodeBlock := blockToGencodeVarint(block)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		if _, err := gencodeBlock.Marshal(nil); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalBlockByGencodeVarintNoTransform(b *testing.B) {
	block := getBlock()
	gencodeBlock := blockToGencodeVarint(block)
	raw, err := gencodeBlock.Marshal(nil)
	if err != nil {
		b.Fatal(err)
	}
	block = getBlock()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var gencodeResult GencodeVarintSignedBlock
		if _, err := gencodeResult.Unmarshal(raw); err != nil {
			b.Fatal(err)
		}

		if validate {
			result := gencodeVarintToBlock(&gencodeResult)
			if !cmp.Equal(result, block) {
				b.Fatal("gencode unmarshal result differs")
			}
		}
	}
}

/* gogoprotobuf

- gogoprotobuf has extensions which can allows us to skip the need to copy the struct
- The extensions should be tried in order to do this
- There is also a flag for deterministic serialization that we should use

- protobuf does not have fixed-size arrays, the length is always encoded as a varint
- protobuf only supports 32-bit and 64-bit ints/uints

- protobuf seems too complex to adopt
*/
