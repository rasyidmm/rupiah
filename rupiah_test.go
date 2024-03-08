package rupiah

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRupiah(t *testing.T) {
	t.Run("rupiah juta", func(t *testing.T) {
		rupiah := Rupiah(1000000)
		assert.Equal(t, "Satu Juta Rupiah", rupiah)
	})

	t.Run("rupiah puluh juta", func(t *testing.T) {
		rupiah := Rupiah(20000000)
		assert.Equal(t, "Dua Puluh Juta Rupiah", rupiah)
	})

	t.Run("rupiah puluh satu juta", func(t *testing.T) {
		rupiah := Rupiah(21000000)
		assert.Equal(t, "Dua Puluh Satu Juta Rupiah", rupiah)
	})
	t.Run("seratus dua puluh satu juta rupiaha", func(t *testing.T) {
		rupiah := Rupiah(121000000)
		assert.Equal(t, "Seratus Dua Puluh Satu Juta Rupiah", rupiah)
	})
	t.Run("satu milyar seratus dua puluh satu juta rupiaha", func(t *testing.T) {
		rupiah := Rupiah(1121000000)
		assert.Equal(t, "Satu Milyar Seratus Dua Puluh Satu Juta Rupiah", rupiah)
	})
	t.Run("tujuh belas", func(t *testing.T) {
		rupiah := Rupiah(17)
		assert.Equal(t, "Tujuh Belas Rupiah", rupiah)
	})
	t.Run("seratus tujuh belas rupiah", func(t *testing.T) {
		rupiah := Rupiah(117)
		assert.Equal(t, "Seratus Tujuh Belas Rupiah", rupiah)
	})

	t.Run("seratus tujuh belas ribu seratus tujuh belas rupiah", func(t *testing.T) {
		rupiah := Rupiah(117117)
		assert.Equal(t, "Seratus Tujuh Belas Ribu Seratus Tujuh Belas Rupiah", rupiah)
	})

	t.Run("tujuh belas ribu seratus tujuh belas rupiah", func(t *testing.T) {
		rupiah := Rupiah(17117)
		assert.Equal(t, "Tujuh Belas Ribu Seratus Tujuh Belas Rupiah", rupiah)
	})

}
