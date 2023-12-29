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

func TestModInt_Add(t *testing.T) {
	type args struct {
		x *big.Int
		y *big.Int
	}
	tests := []struct {
		name string
		mi   *modInt
		args args
		want *big.Int
	}{
		{
			name: "happy path",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(3),
			},
			want: big.NewInt(5),
		},
		{
			name: "happy path overflow",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(7),
				y: big.NewInt(8),
			},
			want: big.NewInt(5),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(3),
			},
			want: big.NewInt(1),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(-3),
			},
			want: big.NewInt(5),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(-3),
			},
			want: big.NewInt(9),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mi.Add(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModInt_Sub(t *testing.T) {
	type args struct {
		x *big.Int
		y *big.Int
	}
	tests := []struct {
		name string
		mi   *modInt
		args args
		want *big.Int
	}{
		{
			name: "happy path",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(3),
			},
			want: big.NewInt(9),
		},
		{
			name: "happy path overflow",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(17),
				y: big.NewInt(8),
			},
			want: big.NewInt(9),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(3),
			},
			want: big.NewInt(5),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(-3),
			},
			want: big.NewInt(1),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(-3),
			},
			want: big.NewInt(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mi.Sub(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModInt_Div(t *testing.T) {
	type args struct {
		x *big.Int
		y *big.Int
	}
	tests := []struct {
		name string
		mi   *modInt
		args args
		want *big.Int
	}{
		{
			name: "happy path",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(7),
				y: big.NewInt(2),
			},
			want: big.NewInt(3),
		},
		{
			name: "happy path",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(3),
			},
			want: big.NewInt(0),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(3),
			},
			want: big.NewInt(9),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(-3),
			},
			want: big.NewInt(1),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(-3),
			},
			want: big.NewInt(0),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(9),
				y: big.NewInt(3),
			},
			want: big.NewInt(3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mi.Div(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModInt_Exp(t *testing.T) {
	type args struct {
		x *big.Int
		y *big.Int
	}
	tests := []struct {
		name string
		mi   *modInt
		args args
		want *big.Int
	}{
		{
			name: "happy path",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(2),
				y: big.NewInt(3),
			},
			want: big.NewInt(8),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(-2),
				y: big.NewInt(3),
			},
			want: big.NewInt(2),
		},
		{
			name: "happy path with one",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(1),
				y: big.NewInt(3),
			},
			want: big.NewInt(1),
		},
		{
			name: "happy path with overflow",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				x: big.NewInt(3),
				y: big.NewInt(4),
			},
			want: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mi.Exp(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestModInt_ModInverse(t *testing.T) {
	type args struct {
		g *big.Int
	}
	tests := []struct {
		name string
		mi   *modInt
		args args
		want *big.Int
	}{
		{
			name: "happy path",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				g: big.NewInt(3),
			},
			want: big.NewInt(7),
		},
		{
			name: "happy path with negative",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				g: big.NewInt(-3),
			},
			want: big.NewInt(3),
		},
		{
			name: "happy path with one",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				g: big.NewInt(1),
			},
			want: big.NewInt(1),
		},
		{
			name: "happy path with overflow",
			mi:   ModInt(big.NewInt(10)),
			args: args{
				g: big.NewInt(11),
			},
			want: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.mi.ModInverse(tt.args.g)
			assert.Equal(t, tt.want, got)
		})
	}
}
