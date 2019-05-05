package sign

import (
	"crypto"
	"crypto/x509"
	"crypto/rsa"
	"fmt"
	"errors"
	"crypto/sha256"
	"crypto/rand"
	"encoding/pem"
)


/*
use pem format string RSA private_key to sign the input
*/
func RsaSign(privpem, input string) (string, error) {
	sum := sha256.Sum256([]byte(input))
	priv, err := parseRsaPrivateKeyFromPemStr(privpem)
	if (err != nil) {
		fmt.Println(err)
		return "", nil
	}

	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	signature, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, sum[:32], &opts)

	if err != nil {
		fmt.Println(err)
		return  "", err
	}

	return string(signature), nil

}

/*
use pem format string RSA public_key to verify the input's signature.
*/
func RsaVerify(pubpem, input, signature string) bool{
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	rsaPublicKey, err := parseRsaPublicKeyFromPemStr(pubpem)
	if (err != nil) {
		fmt.Println(err)
		return false
	}

	sum := sha256.Sum256([]byte(input))
	err = rsa.VerifyPSS(rsaPublicKey, crypto.SHA256, sum[:32], []byte(signature), &opts)

	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

//=================================================================================
/* helper functions */

func parseRsaPrivateKeyFromPemStr(privPEM string) (*rsa.PrivateKey, error) {
    block, _ := pem.Decode([]byte(privPEM))
    if block == nil {
            return nil, errors.New("failed to parse PEM block containing the privkey")
    }

    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
            return nil, err
    }

    return priv, nil
}

func parseRsaPublicKeyFromPemStr(pubPEM string) (*rsa.PublicKey, error) {
        block, _ := pem.Decode([]byte(pubPEM))
        if block == nil {
                fmt.Println("failed to parse certificate PEM")
		return nil, errors.New("failed to parse PEM block containing the cert")
        }
        cert, err := x509.ParseCertificate(block.Bytes)
        if err != nil {
                fmt.Println("failed to parse certificate: ", err.Error())
		return nil, err
        }

	pub := cert.PublicKey.(*rsa.PublicKey)

	return pub, nil
}
