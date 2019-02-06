# skycoin-serialization-benchmarks
Benchmarks of serializers for core Skycoin types

## How to run

```sh
go test -v -bench='.*' ./
```

## Serializers benchmarked

* Skycoin encoder (reflect-based) https://godoc.org/github.com/skycoin/skycoin/src/cipher/encoder
* Skyencoder (generated code) https://github.com/skycoin/skyencoder
* Gencode (with and without varints) https://github.com/andyleap/gencode
* Gotiny https://github.com/niubaoshu/gotiny
* XDR2 https://github.com/davecgh/go-xdr/tree/master/xdr2
* Colfer https://github.com/pascaldekloe/colfer
* JSON https://golang.org/pkg/encoding/json/

Serialization and deserialization is benchmarked for each of the above, using a `coin.SignedBlock` from the [Skycoin](github.com/skycoin/skycoin) source.
This block has 3 transactions, each with 3 inputs and 3 outputs.

For serializers that rely upon code generation, the conversion between the generated struct and `coin.SignedBlock`
is included in the benchmarked code, since this conversion is necessary in many cases.
Benchmarks with `NoTransform` in the name skip the transformation step between `coin.SignedBlock` and the generated struct or vice-versa.

Flatbuffers and protobuf are not benchmarked due to their internal complexity.
However, [gogoprotobuf](https://github.com/gogo/protobuf) should be tested, because it claims to offer a mechanism
to generate code that matches existing structs and would not require a copy step between the generated code's struct
and the program's struct.

## Serializer notes

### Skycoin encoder

This is based upon a simple binary encoding format. It is similar to XDR except that it is little-endian and does not pad to
4-byte boundaries. It uses reflect-based encoding at runtime.

### Skyencoder

This is the same as the reflect-based Skycoin encoder but generates the code to eliminate reflection and minimize memory allocations.

### Gencode

This uses code generation from a custom schema definition. The generated code is relatively simply but a bug was encountered
when generating this code. The generated code would not compile due to a minor error and the generated code has been repaired manually.
Uses varints, which are optional except for variable-length array prefixes, which are always varints.

The "gencode with varints" benchmark uses varints in the schema file.

### Gotiny

This uses reflect-based encoding but not at runtime. It reflects an object once during initialization to build a tree that is used
for subsequent encoding.

Glancing at the code, `unsafe.Pointer` is used, the reason unclear. The code would need auditing before use.

Documentation is in Chinese so it is difficult to understand in detail. However, it appears to use varints.

### Colfer

This uses code generation from a custom schema definition. It does not have support for fixed-size arrays, so all hash and
signature objects must be variable length, which is less efficient.  It uses varints.

The source code for colfer is fairly readable and could be used as a model to build a code generator for the Skycoin encoder.

### JSON

This is only included as a reference point. It is not suitable for encoding `coin.SignedBlock`.

## Results

MBP Mid 2015 Base Model

```sh
$ go test . -bench '.*' -v
```

```
=== RUN   TestMarshaledBlockLen
sky:				 1546 bytes
skyenc:				 1546 bytes
xdr2:				 1612 bytes
json:				 5673 bytes
tiny:				 1423 bytes
colf:				 1535 bytes
genc:				 1516 bytes
gencvar:			 1421 bytes
--- PASS: TestMarshaledBlockLen (0.00s)
goos: darwin
goarch: amd64
pkg: github.com/gz-c/skycoin-serialization-benchmarks
BenchmarkMarshalBlockBySky-8                          	  100000	     20653 ns/op	    4304 B/op	     186 allocs/op
BenchmarkUnmarshalBlockBySky-8                        	  100000	     14662 ns/op	    5672 B/op	     131 allocs/op
BenchmarkMarshalBlockBySkyencoder-8                   	 5000000	       420 ns/op	       0 B/op	       0 allocs/op
BenchmarkMarshalBlockBySkyencoderWithAlloc-8          	 2000000	       701 ns/op	    1792 B/op	       1 allocs/op
BenchmarkUnmarshalBlockBySkyencoder-8                 	 1000000	      1039 ns/op	    1648 B/op	      10 allocs/op
BenchmarkMarshalBlockByXDR2-8                         	  100000	     16703 ns/op	    9264 B/op	     177 allocs/op
BenchmarkUnmarshalBlockByXDR2-8                       	  100000	     18757 ns/op	    5936 B/op	     212 allocs/op
BenchmarkMarshalBlockByJSON-8                         	   20000	     61305 ns/op	    6375 B/op	       2 allocs/op
BenchmarkUnmarshalBlockByJSON-8                       	    5000	    286629 ns/op	    6976 B/op	    1305 allocs/op
BenchmarkMarshalBlockByGotiny-8                       	  200000	      6197 ns/op	     224 B/op	       1 allocs/op
BenchmarkUnmarshalBlockByGotiny-8                     	  200000	      7334 ns/op	    2192 B/op	      21 allocs/op
BenchmarkMarshalBlockByColfer-8                       	  300000	      4314 ns/op	    4960 B/op	      70 allocs/op
BenchmarkUnmarshalBlockByColfer-8                     	  300000	      4934 ns/op	    4832 B/op	      70 allocs/op
BenchmarkMarshalBlockByColferNoTransform-8            	 1000000	      1044 ns/op	    1536 B/op	       1 allocs/op
BenchmarkUnmarshalBlockByColferNoTransform-8          	  500000	      3682 ns/op	    3184 B/op	      60 allocs/op
BenchmarkMarshalBlockByGencode-8                      	  500000	      2438 ns/op	    3408 B/op	      12 allocs/op
BenchmarkUnmarshalBlockByGencode-8                    	  500000	      2469 ns/op	    3296 B/op	      20 allocs/op
BenchmarkMarshalBlockByGencodeNoTransform-8           	 2000000	       807 ns/op	    1536 B/op	       1 allocs/op
BenchmarkUnmarshalBlockByGencodeNoTransform-8         	 2000000	      1004 ns/op	    1648 B/op	      10 allocs/op
BenchmarkMarshalBlockByGencodeVarint-8                	  500000	      2836 ns/op	    3408 B/op	      12 allocs/op
BenchmarkUnmarshalBlockByGencodeVarint-8              	  500000	      2587 ns/op	    3296 B/op	      20 allocs/op
BenchmarkMarshalBlockByGencodeVarintNoTransform-8     	 1000000	      1405 ns/op	    1536 B/op	       1 allocs/op
BenchmarkUnmarshalBlockByGencodeVarintNoTransform-8   	 1000000	      1181 ns/op	    1648 B/op	      10 allocs/op
PASS
ok  	github.com/gz-c/skycoin-serialization-benchmarks	38.815s
```

## Results interpretation

Varint encoding provides at best an 8% improvement in size.
A `coin.SignedBlock` is mostly hashes or signatures which do not compress well in general,
and do not benefit from varint encoding, limiting the potential size gains.

Gencode without varints (except the mandatory varints in variable-length array prefixes) without the transformation step
is the fastest. This is also the most similar to what a code generator for the Skycoin encoder would be.
So, a code generator for the Skycoin encoder could expect similar performance,
which is ~23x faster serialization and ~15x faster deserialization, as well as a massive reduction in memory allocations.
