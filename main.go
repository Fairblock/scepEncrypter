package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/spf13/cobra"
)

func main() {
	var pubKeyBase64 string
	var keyshare string

	var rootCmd = &cobra.Command{
		Use:   "encrypt",
		Short: "A CLI tool to take two strings as input and output them",
		Run: func(cmd *cobra.Command, args []string) {
			// fmt.Println("Input 1:", pubKeyBase64)
			// fmt.Println("Input 2:", keyshare)

			// Decode the base64 public key
			pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKeyBase64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Load the secp256k1 public key
			pubKey, err := btcec.ParsePubKey(pubKeyBytes, btcec.S256())
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			ciphertext, err := btcec.Encrypt(pubKey, []byte(keyshare))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Encode ciphertext as hex for easy handling
			fmt.Println(hex.EncodeToString(ciphertext))
		},
	}

	// Flags for input
	rootCmd.Flags().StringVarP(&pubKeyBase64, "pubkey64", "p", "", "pubkey to encrypt keyshare")
	rootCmd.Flags().StringVarP(&keyshare, "keyshare", "k", "", "keyshare in hex")

	// Mark the flags as required
	rootCmd.MarkFlagRequired("pubkey64")
	rootCmd.MarkFlagRequired("keyshare")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
