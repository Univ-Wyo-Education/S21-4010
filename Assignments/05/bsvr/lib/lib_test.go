package lib

import (
	"fmt"
	"testing"
	// "github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	// "github.com/pschlump/godebug"
)

func TestIsValidAddress(t *testing.T) {
	tests := []struct {
		AddrStr string
		expect  bool
	}{
		{AddrStr: "0x885765a2fcfB72e68D82D656C6411b7D27BacDD7", expect: true},
		{AddrStr: "0x885765a2fcfB72e68D82D656C6411b7D27BacD7", expect: false},
		{AddrStr: "0x885765a2fcfB72e68D82D656C6411b7D27BacD700", expect: false},
		{AddrStr: "885765a2fcfB72e68D82D656C6411b7D27BacDD7", expect: false},
		{AddrStr: "885765a2fcfB72e68D82D656C6411b7D27BacDD700", expect: false},
	}

	for ii, test := range tests {
		// func IsValidAddress(addrStr string) bool {
		got := IsValidAddress(test.AddrStr)
		if got != test.expect {
			t.Errorf("Test %d, expected %v got %v for %s\n", ii, test.expect, got, test.AddrStr)
		}
	}

}

func TestConvAddrStrToAddressType(t *testing.T) {
	tests := []struct {
		AddrStr    string
		expectErr  bool
		expectAddr string
	}{
		{
			AddrStr:    "0x885765a2fcfB72e68D82D656C6411b7D27BacDD7",
			expectErr:  false,
			expectAddr: "0x885765a2fcfb72e68d82d656c6411b7d27bacdd7",
		},
		{AddrStr: "0x885765a2fcfB72e68D82D656C6411b7D27BacD7", expectErr: true},
		{AddrStr: "0x885765a2fcfB72e68D82D656C6411b7D27BacD700", expectErr: true},
		{AddrStr: "885765a2fcfB72e68D82D656C6411b7D27BacDD7", expectErr: true},
		{AddrStr: "885765a2fcfB72e68D82D656C6411b7D27BacDD700", expectErr: true},
	}

	for ii, test := range tests {
		// func ConvAddrStrToAddressType(addrStr string) (addr addr.AddressType, err error) {
		aa, err := ConvAddrStrToAddressType(test.AddrStr)
		if err != nil && test.expectErr == false {
			t.Errorf("Test %d, expected success got err %s\n", ii, err)
		} else if err == nil && test.expectErr == true {
			t.Errorf("Test %d, expected error got none, for [%s]\n", ii, test.AddrStr)
		} else if test.expectErr == false {
			s := fmt.Sprintf("%s", aa)
			// fmt.Printf("Just Before AT:%s ->%s<-\n", godebug.LF(), s)
			if s != test.expectAddr {
				t.Errorf("Test %d\n   expected ->%s<-\n   got      ->%s<-\n", ii, test.expectAddr, aa)
			}
		}
	}

}
