package addresses

import "testing"

type testCase struct {
	address          string
	expected_address string
}

func TestAddressType(t *testing.T) {

	t.Parallel() // Run Test in Parallel

	testCases := []testCase{
		{"Street A", "Street"},
		{"Avenue B", "Avenue"},
		{"Road C", "Road"},
		{"Highway D", "Highway"},
		{"Garden E", "Invalid Type!"},
		{"Avenida 1", "Invalid Type!"},
		{"street X", "Street"},
		{"", "Invalid Type!"},
	}

	for _, testCase := range testCases {
		received_address := AddressType(testCase.address)
		if received_address != testCase.expected_address {
			t.Errorf("Address type is different from expected! Expected: %s, Received: %s",
				testCase.expected_address,
				received_address,
			)
		}
	}

}

func TestAny(t *testing.T) {

	t.Parallel()

	if 1 > 2 {
		t.Errorf("Broke!")
	}
}
