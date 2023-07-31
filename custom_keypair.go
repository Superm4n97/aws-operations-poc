package main

import (
	"fmt"
	"os"
	"os/exec"
)

func generateSSHKeyPair() error {
	cmd := exec.Command("ssh-keygen", "-t", "rsa", "-b", "2048", "-f", "id_rsa", "-N", "")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to generate SSH key pair: %v", err)
	}
	return nil
}

func readPublicKey(publicKeyPath string) string {
	// Read the contents of the SSH public key file
	// Replace with your own file reading implementation
	// Here's an example:
	contents, err := os.ReadFile(publicKeyPath)
	if err != nil {
		fmt.Println(fmt.Sprintf("Failed to read public key: %v", err))
	}
	publicKey := string(contents)

	return publicKey
}
