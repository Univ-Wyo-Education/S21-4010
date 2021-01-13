package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (cc *CLI) InstructorValidateSignature(addr, sig, msg string) (isValid bool, err error) {
	_, _, err = VerifySignature(addr, sig, msg) // (recoveredAddress, recoveredPublicKey string, err error) {
	if err != nil {
		return // isValid is false.
	}

	isValid = true
	return
}

// VerifySignature takes hex encoded addr, sig and msg and verifies that the signature matches with the address.
// Pulled form go-ethereum code.
func VerifySignature(addr, sig, msg string) (recoveredAddress, recoveredPublicKey string, err error) {
	message, err := hex.DecodeString(msg)
	if err != nil {
		return "", "", fmt.Errorf("unabgle to decode message (invalid hex data) Error:%s", err)
	}
	if !common.IsHexAddress(addr) {
		return "", "", fmt.Errorf("invalid address: %s", addr)
	}
	address := common.HexToAddress(addr)
	signature, err := hex.DecodeString(sig)
	if err != nil {
		return "", "", fmt.Errorf("signature is not valid hex Error:%s", err)
	}

	recoveredPubkey, err := crypto.SigToPub(signHash(message), signature)
	if err != nil || recoveredPubkey == nil {
		return "", "", fmt.Errorf("signature verification failed Error:%s", err)
	}
	recoveredPublicKey = hex.EncodeToString(crypto.FromECDSAPub(recoveredPubkey))
	rawRecoveredAddress := crypto.PubkeyToAddress(*recoveredPubkey)
	if address != rawRecoveredAddress {
		return "", "", fmt.Errorf("signature did not verify, addresses did not match")
	}
	recoveredAddress = rawRecoveredAddress.Hex()
	return
}

// signHash is a helper function that calculates a hash for the given message
// that can be safely used to calculate a signature from.
//
// The hash is calulcated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
