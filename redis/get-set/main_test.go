package main

import (
	"go-examples/common/config"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	key   string
	value string
}{
	{
		key:   "world",
		value: "Hello world",
	},
	{
		key:   "123",
		value: "Hello 123",
	},
}

func TestNewCache(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	cfg := &config.Redis{
		Server: s.Addr(),
	}
	cache := NewCache(cfg)
	assert.NotNil(t, cache)
}

func TestCache_SetValue(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	cfg := &config.Redis{
		Server: s.Addr(),
	}
	cache := NewCache(cfg)

	for _, tt := range tests {
		err := cache.SetValue(tt.key, tt.value)
		assert.NoError(t, err)
	}
}

func TestCache_GetValue(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	cfg := &config.Redis{
		Server: s.Addr(),
	}
	cache := NewCache(cfg)

	for _, tt := range tests {
		err := s.Set(tt.key, tt.value)
		assert.NoError(t, err)

		val, err := cache.GetValue(tt.key)
		assert.NoError(t, err)
		assert.Equal(t, tt.value, val)
	}
}
