package ssh

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
	// "io/ioutil"
	// "fmt"
	"strings"
)

// GenSSHKeyPair generates and returns SSH private/public key-pair as strings and an error
func GenSSHKeyPair() (string, string, error) {
	// savePrivateFileTo := "./id_rsa_test"
	// savePublicFileTo := "./id_rsa_test.pub"
	bitSize := 2048

	initKey, err := genPrivKey(bitSize)
	if err != nil {
		return "", "", err
	}

	publicKey, err := genPubKey(&initKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	privateKey := encPrivKeyToPEM(initKey)

	return strings.TrimSpace(privateKey), strings.TrimSpace(publicKey), nil

	// err = writeKeyToFile(privateKeyBytes, savePrivateFileTo)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// err = writeKeyToFile([]byte(publicKeyBytes), savePublicFileTo)
	// if err != nil {
	// 	panic(err.Error())
	// }
}

// genPrivKey creates a RSA Private Key of specified byte size
func genPrivKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	// Validate Private Key
	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// encPrivKeyToPEM encodes Private Key from RSA to PEM format
func encPrivKeyToPEM(privateKey *rsa.PrivateKey) string {
	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	// Private key in PEM format
	privatePEM := string(pem.EncodeToMemory(&privBlock))

	return privatePEM
}

// genPubKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func genPubKey(privatekey *rsa.PublicKey) (string, error) {
	publicRsaKey, err := ssh.NewPublicKey(privatekey)
	if err != nil {
		return "", err
	}

	pubKeyBytes := string(ssh.MarshalAuthorizedKey(publicRsaKey))

	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
// func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
// 	err := ioutil.WriteFile(saveFileTo, keyBytes, 0600)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("Key saved to: %s \n", saveFileTo)
// 	return nil
// }
