package cert

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	_"math/big"
	_"os"
	"time"
	_"fmt"
	"bytes"
	"math/big"

	"crypto/tls"
)

// create a RSA CA
// input: name
// output: cert, priv_key
func GenerateCA(name string) (cert string, priv_key string, err error) {
	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	ca := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"HUBCA"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		//AuthorityKeyId:      []byte{1, 2, 3, 4, 6},
	}

	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("create ca failed", err)
		return "", "", err
	}

	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	//fmt.Println(certOut.String())

	keyOut := new(bytes.Buffer)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// create an Ecdsa CA
// input: name
// output: cert, priv_key
func GenerateEcdsaCA(name string) (cert string, priv_key string, err error) {
	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	ca := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"HUBCA"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		//AuthorityKeyId:      []byte{1, 2, 3, 4, 6},
	}

	//priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := &priv.PublicKey
	ca_b, err := x509.CreateCertificate(rand.Reader, ca, ca, pub, priv)
	if err != nil {
		log.Println("create ca failed", err)
		return "", "", err
	}

	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: ca_b})
	//fmt.Println(certOut.String())

	keyOut := new(bytes.Buffer)
	priv_marshal, _ := x509.MarshalECPrivateKey(priv)
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: priv_marshal})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// Generate a cert
// input: name, certCAPem, keyCAPem
// output: server cert, priv_key
func GenerateServerCert(name, certCAPem, keyCAPem string) (scert string, priv_key string, err error) {
	// Load CA
	catls, err := tls.X509KeyPair([]byte(certCAPem), []byte(keyCAPem))
	if err != nil {
		return "", "", err
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return "", "", err
	}

	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	cert := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"HUB"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		//SubjectKeyId: []byte{1, 2, 3, 4, 6},
		//ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		IsCA:                  false,
		BasicConstraintsValid: true,
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)

	// Public key
	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	//fmt.Println(certOut.String())

	// Private key
	keyOut := new(bytes.Buffer)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// Generate an ECDSA cert
// input: name, certCAPem, keyCAPem
// output: server cert, priv_key
func GenerateEcdsaServerCert(name, certCAPem, keyCAPem string) (scert string, priv_key string, err error) {
	// Load CA
	catls, err := tls.X509KeyPair([]byte(certCAPem), []byte(keyCAPem))
	if err != nil {
		return "", "", err
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return "", "", err
	}

	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	cert := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"HUB"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		//SubjectKeyId: []byte{1, 2, 3, 4, 6},
		//ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		IsCA:                  false,
		BasicConstraintsValid: true,
	}
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)

	// Public key
	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	//fmt.Println(certOut.String())

	// Private key
	keyOut := new(bytes.Buffer)
	priv_marshal, _ := x509.MarshalECPrivateKey(priv)
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: priv_marshal})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// Generate a client cert
// input: name, certCAPem, keyCAPem
// output: server cert, priv_key
func GenerateClientCert(name, certCAPem, keyCAPem string) (scert string, priv_key string, err error) {
	// Load CA
	catls, err := tls.X509KeyPair([]byte(certCAPem), []byte(keyCAPem))
	if err != nil {
		return "", "", err
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return "", "", err
	}

	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	cert := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"DataCenter"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		IsCA:                  false,
		BasicConstraintsValid: true,
	}
	priv, _ := rsa.GenerateKey(rand.Reader, 2048)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)

	// Public key
	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	//fmt.Println(certOut.String())

	// Private key
	keyOut := new(bytes.Buffer)
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// Generate an ecdsa client cert
// input: name, certCAPem, keyCAPem
// output: server cert, priv_key
func GenerateEcdsaClientCert(name, certCAPem, keyCAPem string) (scert string, priv_key string, err error) {
	// Load CA
	catls, err := tls.X509KeyPair([]byte(certCAPem), []byte(keyCAPem))
	if err != nil {
		return "", "", err
	}
	ca, err := x509.ParseCertificate(catls.Certificate[0])
	if err != nil {
		return "", "", err
	}

	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	cert := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"DataCenter"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		IsCA:                  false,
		BasicConstraintsValid: true,
	}
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, ca, pub, catls.PrivateKey)

	// Public key
	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	//fmt.Println(certOut.String())

	// Private key
	keyOut := new(bytes.Buffer)
	priv_marshal, _ := x509.MarshalECPrivateKey(priv)
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: priv_marshal})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// Generate an ecdsa self-signed cert
// input: name
// output: cert, priv_key
func GenerateEcdsaSelfsignCert(name string) (scert string, priv_key string, err error) {
	sn, err := generateSerialNumber()
	if err != nil {
		return "", "", err
	}

	cert := &x509.Certificate{
		SerialNumber: sn,
		Subject: pkix.Name{
			CommonName: name,
			Organization:  []string{"DataCenter"},
			Country:       []string{"US"},
			Province:      []string{"CA"},
			Locality:      []string{"SF"},
			StreetAddress: []string{"MISSION ST."},
			PostalCode:    []string{"94105"},
		},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		IsCA:                  false,
		BasicConstraintsValid: true,
	}
	priv, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	pub := &priv.PublicKey

	// Sign the certificate
	cert_b, err := x509.CreateCertificate(rand.Reader, cert, cert, pub, priv)

	// Public key
	certOut := new(bytes.Buffer)
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert_b})
	//fmt.Println(certOut.String())

	// Private key
	keyOut := new(bytes.Buffer)
	priv_marshal, _ := x509.MarshalECPrivateKey(priv)
	pem.Encode(keyOut, &pem.Block{Type: "EC PRIVATE KEY", Bytes: priv_marshal})
	//fmt.Println(keyOut.String())

	return certOut.String(), keyOut.String(), nil
}

// GenerateSerialNumber generates a serial number suitable for a certificate
func generateSerialNumber() (*big.Int, error) {
        serial, err := rand.Int(rand.Reader, (&big.Int{}).Exp(big.NewInt(2), big.NewInt(159), nil))
        if err != nil {
                return nil, err
        }
        return serial, nil
}

