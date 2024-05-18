package lib

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"golang.org/x/crypto/sha3"
)

func GetAbi() (*abi.ABI, error) {
	path, _ := filepath.Abs("./abi/timekeeping.json")
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	edabi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return nil, err
	}

	return &edabi, nil
}

func Keccak256(data []byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(data)
	return hasher.Sum(nil)
}
