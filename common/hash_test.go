// Copyright © 2019-2020 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package common

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSHA512_256(t *testing.T) {
	tests := []struct {
		name     string
		srcBytes [][]byte
		want     []byte
	}{
		{
			name:     "one",
			srcBytes: [][]byte{[]byte("one")},
			want: []byte{155, 131, 203, 249, 98, 196, 229, 70, 2, 28, 211, 87, 227, 190, 237, 33, 234, 222,
				237, 187, 64, 76, 128, 71, 25, 78, 154, 136, 22, 45, 51, 41,
			},
		},
		{
			name:     "two",
			srcBytes: [][]byte{[]byte("hello"), []byte("world")},
			want: []byte{
				123, 126, 124, 145, 206, 51, 245, 169, 8, 47, 212, 46, 66, 170, 66, 11, 82, 160,
				117, 28, 8, 114, 142, 122, 134, 191, 158, 155, 65, 179, 239, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := SHA512_256(tt.srcBytes...)
			if len(hash) != 32 {
				t.Fatal("SHA512_256 returned unexpected length")
			}
			for idx, actual := range hash {
				expected := tt.want[idx]
				if actual != expected {
					t.Fatal("SHA512_256 returned unexpected value")
				}
			}
		})
	}
}

func TestSHA512_256i(t *testing.T) {
	tests := []struct {
		name     string
		srcBytes []*big.Int
		want     []byte
	}{
		{
			name:     "one",
			srcBytes: []*big.Int{big.NewInt(12345678)},
			want: []byte{
				67, 219, 167, 235, 231, 133, 107, 20, 13, 26, 137, 209, 227, 44, 166, 243, 178, 187, 225,
				8, 188, 216, 190, 110, 158, 214, 125, 4, 251, 94, 93, 188,
			},
		},
		{
			name:     "two",
			srcBytes: []*big.Int{big.NewInt(12345678), big.NewInt(34567890)},
			want: []byte{
				204, 108, 54, 96, 23, 83, 16, 141, 6, 196, 205, 169, 56, 190, 16, 86, 190, 140, 255, 179, 57, 7, 138,
				28, 226, 9, 15, 169, 24, 135, 190, 32,
			},
		},
		{
			name:     "three",
			srcBytes: []*big.Int{big.NewInt(-1), big.NewInt(0), big.NewInt(-500)},
			want: []byte{
				149, 83, 31, 152, 21, 77, 151, 186, 148, 244, 25, 56, 98, 166, 118, 245, 200, 210, 53, 19, 195, 30,
				234, 21, 199, 87, 18, 155, 79, 33, 210, 65,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SHA512_256i(tt.srcBytes...)
			hash := res.Bytes()
			if len(hash) != 32 {
				t.Fatal("SHA512_256 returned unexpected length")
			}
			print(fmt.Sprintf("hash(%s): [", tt.name))
			for idx, actual := range hash {
				print(fmt.Sprintf("%d, ", actual))
				expected := tt.want[idx]
				if actual != expected {
					t.Fatal("SHA512_256 returned unexpected value")
				}
			}
			println("]")
		})
	}
}

func TestSHA512_256i_TAGGED(t *testing.T) {
	tests := []struct {
		name     string
		tag      []byte
		srcBytes []*big.Int
		want     []byte
	}{
		{
			name:     "one",
			tag:      []byte("tag-a"),
			srcBytes: []*big.Int{big.NewInt(12345678)},
			want: []byte{
				62, 229, 129, 172, 169, 125, 219, 131, 105, 95, 195,
				233, 170, 196, 197, 213, 236, 163, 114, 155, 156,
				196, 165, 198, 67, 235, 246, 30, 140, 248, 88, 95,
			},
		},
		{
			name:     "two",
			tag:      []byte("tag-b"),
			srcBytes: []*big.Int{big.NewInt(12345678), big.NewInt(34567890)},
			want: []byte{
				27, 182, 159, 219, 92, 228, 45, 221, 84, 231, 52, 154,
				154, 33, 20, 84, 83, 190, 12, 89, 205, 95, 64, 217, 176,
				132, 5, 157, 75, 168, 73, 38,
			},
		},
		{
			name:     "three",
			tag:      []byte("tag-c"),
			srcBytes: []*big.Int{big.NewInt(12345678), big.NewInt(0), big.NewInt(34567890)},
			want: []byte{
				213, 63, 56, 234, 196, 75, 69, 74, 68, 71, 105, 213, 75,
				149, 4, 237, 211, 185, 20, 151, 149, 84, 187, 218, 108,
				208, 171, 58, 202, 185, 168, 189,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := SHA512_256i_TAGGED(tt.tag, tt.srcBytes...)
			hash := res.Bytes()
			if len(hash) != 32 {
				t.Fatal("SHA512_256 returned unexpected length")
			}
			print(fmt.Sprintf("hash(%s): [", tt.name))
			for idx, actual := range hash {
				print(fmt.Sprintf("%d, ", actual))
				expected := tt.want[idx]
				if actual != expected {
					t.Fatal("SHA512_256 returned unexpected value")
				}
			}
			println("]")
		})
	}
}
