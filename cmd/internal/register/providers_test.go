package register_test

import (
	"testing"

	"github.com/accek/tegola/cmd/internal/register"
	"github.com/accek/tegola/dict"
)

func TestProviders(t *testing.T) {
	type tcase struct {
		config      []dict.Dict
		expectedErr error
	}

	fn := func(t *testing.T, tc tcase) {
		var err error

		// convert []dict.Dict -> []dict.Dicter
		provArr := make([]dict.Dicter, len(tc.config))
		for i := range provArr {
			provArr[i] = tc.config[i]
		}

		_, err = register.Providers(provArr)
		if tc.expectedErr != nil {
			if err.Error() != tc.expectedErr.Error() {
				t.Errorf("invalid error. expected: %v, got %v", tc.expectedErr, err.Error())
			}
			return
		}
		if err != nil {
			t.Errorf("unexpected err: %v", err)
			return
		}
	}

	tests := map[string]tcase{
		"missing name": {
			config: []dict.Dict{
				{
					"type": "postgis",
				},
			},
			expectedErr: register.ErrProviderNameMissing,
		},
		"name is not string": {
			config: []dict.Dict{
				{
					"name": 1,
				},
			},
			expectedErr: register.ErrProviderNameInvalid,
		},
		"missing type": {
			config: []dict.Dict{
				{
					"name": "test",
				},
			},
			expectedErr: register.ErrProviderTypeMissing("test"),
		},
		"invalid type": {
			config: []dict.Dict{
				{
					"name": "test",
					"type": 1,
				},
			},
			expectedErr: register.ErrProviderTypeInvalid("test"),
		},
		"already registered": {
			config: []dict.Dict{
				{
					"name": "test",
					"type": "debug",
				},
				{
					"name": "test",
					"type": "debug",
				},
			},
			expectedErr: register.ErrProviderAlreadyRegistered("test"),
		},
		"success": {
			config: []dict.Dict{
				{
					"name": "test",
					"type": "debug",
				},
			},
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) { fn(t, tc) })
	}
}
