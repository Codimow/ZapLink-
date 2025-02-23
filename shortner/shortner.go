package shortener

import (
	"crypto/sha256"
	"fmt"
	"os"
	"github.com/itchyny/base58-go"
)

func sha256f(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	return string(hash.Sum(nil))
}

func base58encoding(byte []byte) string {
	encoding:= base58.BitcoinEncoding
	encoded, err := encoding.Encode(byte)
	if err != nil {
		fmt.Println("Error in encoding")
		os.Exit(1)
	}
	return string(encoded)
}

func generateShortURL(initalLink string, userID string) string {
	urlHash := sha256f(initalLink + userID)
	generateNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalStrign := base58Encoded([]byte(fmt.Sprintf("%d", generateNumber)))
	return finalStrign[:8]
}
