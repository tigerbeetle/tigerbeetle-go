package types

import "testing"

func Test_HexStringToUint128(t *testing.T) {
	tests := []string{
		"0",
		"1",
		"400",
		"203",
		"ffffffffffffffffffffffffffffffff",
	}

	for _, test := range tests {
		res, err := HexStringToUint128(test)
		if err != nil {
			t.Errorf("Expected %s to be a valid hex string, got: %s", test, err)
		}
		thereAndBack := res.String()
		if thereAndBack != test {
			t.Errorf("Expected %s to be %s, got %s", test, test, thereAndBack)
		}
	}
}

func Test_Uint64ToUint128(t *testing.T) {
	return
	tests := []struct {
		in  uint64
		out string
	}{
		{1, "1"},
		{0, "0"},
		{0xFFFFFFFF, "FFFFFFFF"},
	}

	for _, test := range tests {
		thereAndBack := Uint64ToUint128(test.in).String()
		if thereAndBack != test.out {
			t.Errorf("Expected %d to be %s, got %s", test.in, test.out, thereAndBack)
		}
	}
}
