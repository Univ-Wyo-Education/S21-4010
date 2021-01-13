package lib

import (
	"encoding/hex"
	"fmt"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	// "github.com/pschlump/godebug"
)

type SignatureType string

func ValidSignature(sig SignatureType, message string, from addr.AddressType) bool {
	// TODO - xyzyz - we will validate signatures in the "Wallet" homework.  AS-05.
	return true
}

// IsValidAddress returns true if the addresss is a validate looking address.
func IsValidAddress(addrStr string) bool {
	if len(addrStr) != 40+2 { // must be 2 chars for 0x, then 20 bytes of public key.
		// fmt.Printf("Len=%d AT: %s\n", len(addrStr), godebug.LF())
		return false
	}
	if addrStr[0:2] != "0x" { // must stat with 0x
		// fmt.Printf("AT: %s\n", godebug.LF())
		return false
	}
	// We could check if the address is EIP-55 encoded also!
	addrStr = addrStr[2:] // Keep all the stuff after the 0x
	_, err := hex.DecodeString(addrStr)
	if err != nil {
		// fmt.Printf("AT: %s\n", godebug.LF())
		return false
	}
	return true
}

func ConvAddrStrToAddressType(addrStr string) (addr addr.AddressType, err error) {
	if !IsValidAddress(addrStr) {
		// fmt.Printf("ERR 1 Input: %s AT: %s\n", addrStr, godebug.LF())
		return []byte{}, fmt.Errorf("invalid address, not 0xHEX...\n")
	}
	tmp, err := hex.DecodeString(addrStr[2:])
	if err != nil {
		// fmt.Printf("ERR 2 AT: %s\n", godebug.LF())
		return []byte{}, err
	}
	addr = tmp
	// fmt.Printf("SUCCESS tmp=%x addr=%s AT: %s\n", tmp, addr, godebug.LF())
	return
}

// RunSmartContract -- Implemented in Assignment 6
func RunSmartContract() {
	// TODO - we will implement this in asssignment 6 - AS-06.
}
