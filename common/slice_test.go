// Copyright © 2019-2020 Binance
//
// This file is part of Binance. The full Binance copyright notice, including
// terms governing use, modification, and redistribution, is contained in the
// file LICENSE at the root of the source code distribution tree.

package common

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestBigIntsToBytes(t *testing.T) {
	type args struct {
		bigInts []*big.Int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "happy path",
			args: args{
				bigInts: []*big.Int{
					big.NewInt(1),
					big.NewInt(2),
					big.NewInt(3),
				},
			},
			want: [][]byte{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "happy path with zero",
			args: args{
				bigInts: []*big.Int{
					big.NewInt(1),
					big.NewInt(2),
					big.NewInt(0),
					big.NewInt(731),
				},
			},
			want: [][]byte{
				{1},
				{2},
				{},
				{2, 219},
			},
		},
		{
			name: "happy path with nil",
			args: args{
				bigInts: []*big.Int{
					big.NewInt(1),
					nil,
					big.NewInt(3),
				},
			},
			want: [][]byte{
				{1},
				nil,
				{3},
			},
		},
		{
			name: "happy path with empty",
			args: args{
				bigInts: []*big.Int{},
			},
			want: [][]byte{},
		},
		{
			name: "happy path of nil",
			args: args{
				bigInts: nil,
			},
			want: [][]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BigIntsToBytes(tt.args.bigInts)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMultiBytesToBigInts(t *testing.T) {
	type args struct {
		bytes [][]byte
	}
	tests := []struct {
		name string
		args args
		want []*big.Int
	}{
		{
			name: "happy path",
			args: args{
				bytes: [][]byte{
					{1},
					{2},
					{3},
				},
			},
			want: []*big.Int{
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(3),
			},
		},
		{
			name: "happy path with zero",
			args: args{
				bytes: [][]byte{
					{1},
					{2},
					{},
					{2, 219},
				},
			},
			want: []*big.Int{
				big.NewInt(1),
				big.NewInt(2),
				big.NewInt(0),
				big.NewInt(731),
			},
		},
		{
			name: "happy path with nil",
			args: args{
				bytes: [][]byte{
					{1},
					nil,
					{3},
				},
			},
			want: []*big.Int{
				big.NewInt(1),
				big.NewInt(0),
				big.NewInt(3),
			},
		},
		{
			name: "happy path with empty",
			args: args{
				bytes: [][]byte{},
			},
			want: []*big.Int{},
		},
		{
			name: "happy path of nil",
			args: args{
				bytes: nil,
			},
			want: []*big.Int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MultiBytesToBigInts(tt.args.bytes)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPadToLengthBytesInPlace(t *testing.T) {
	type args struct {
		src    []byte
		length int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "happy path",
			args: args{
				src:    []byte{1, 2, 3},
				length: 5,
			},
			want: []byte{0, 0, 1, 2, 3},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{1, 2, 3},
				length: 3,
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{1, 2, 3},
				length: 2,
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{1, 2, 3},
				length: 1,
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{1, 2, 3},
				length: 0,
			},
			want: []byte{1, 2, 3},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{},
				length: 0,
			},
			want: []byte{},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{},
				length: 1,
			},
			want: []byte{0},
		},
		{
			name: "happy path with zero",
			args: args{
				src:    []byte{},
				length: 2,
			},
			want: []byte{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PadToLengthBytesInPlace(tt.args.src, tt.args.length)
			assert.Equal(t, tt.want, got)
		})
	}
}
