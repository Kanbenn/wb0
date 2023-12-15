package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	type args struct {
		key string
		val []byte
	}
	type res struct {
		val   []byte
		found bool
	}
	c := NewCache()
	tests := []struct {
		name string
		c    *Cache
		args args
		resp res
	}{
		{"shall pass", c,
			args{"some_order_id", []byte("value")},
			res{[]byte("value"), true}},
		{"empty key error", c,
			args{"", []byte("some value")},
			res{[]byte(nil), false}},
		{"key already exists error", c,
			args{"some_order_id", []byte("new value")},
			res{[]byte("value"), true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Add(tt.args.key, tt.args.val)
			val, found := tt.c.Get(tt.args.key)
			assert.Equal(t, tt.resp.val, val)
			assert.Equal(t, tt.resp.found, found)
		})
	}
}
