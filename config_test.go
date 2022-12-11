package nssession

import (
	"errors"
	"testing"

	"github.com/no-src/nssession/store"
	"github.com/no-src/nssession/store/memory"
)

func TestInitDefaultConfig(t *testing.T) {
	testCases := []struct {
		name   string
		config *Config
		expect error
	}{
		{"nil config", nil, errNilConfig},
		{"nil store", &Config{}, errNilStore},
		{"use default value", &Config{Store: store.NewStore(memory.Driver)}, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := InitDefaultConfig(tc.config)
			if !errors.Is(err, tc.expect) {
				t.Errorf("expect to get error %v, but get %v", tc.expect, err)
			}
		})
	}
}
