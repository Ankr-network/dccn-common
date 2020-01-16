package cert_test

import (
	"testing"
	"fmt"

	certmanager "github.com/Ankr-network/dccn-common/cert"
)

const        ECDSA_CLIENT_CERT = `
-----BEGIN CERTIFICATE-----
MIICJzCCAc6gAwIBAgIUQNK8zuB47TrjMK/9apa4+ODmGP8wCgYIKoZIzj0EAwIw
dDELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEUMBIGA1UE
CRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MQ4wDAYDVQQKEwVIVUJDQTEV
MBMGA1UEAxMMbXlodWItY2EuY29tMB4XDTE5MDUxMjAxNDY1NVoXDTI5MDUxMjAx
NDY1NVowfTELMAkGA1UEBhMCVVMxCzAJBgNVBAgTAkNBMQswCQYDVQQHEwJTRjEU
MBIGA1UECRMLTUlTU0lPTiBTVC4xDjAMBgNVBBETBTk0MTA1MRMwEQYDVQQKEwpE
YXRhQ2VudGVyMRkwFwYDVQQDExBteWRhdGFjZW50ZXIuY29tMFkwEwYHKoZIzj0C
AQYIKoZIzj0DAQcDQgAEM49mdr428vS5+uHc0wjJBqyQ5n8d0QLra97C40uaEw94
l6RWjMOGbQfHGg6YbZzQ6Zc0qIxf7xu+RX//sTmqCaM1MDMwDgYDVR0PAQH/BAQD
AgeAMBMGA1UdJQQMMAoGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwCgYIKoZIzj0E
AwIDRwAwRAIgUxRoNWAjjyvTmnzU8c8s02g0wZURKGo76kh9LNVXcp4CIBAvaZ5u
Y88YwWeiSVJNBDC6MIcgPLAM4YuLvNjP6M6W
-----END CERTIFICATE-----`

func TestGenerateRsaKeys(t *testing.T) {
	t.Log("Testing GenerateCA")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")

	if err == nil {
		t.Logf("\nCA Cert: %s\n CA Public Key: %s\n", caCert, caKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateEcdsaKeys(t *testing.T) {
	t.Log("Testing Generate Ecdsa CA")

	caCert, caKey, err := certmanager.GenerateEcdsaCA("myhub-ca.com")

	if err == nil {
		t.Logf("\nCA Cert: %s\n CA Public Key: %s\n", caCert, caKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateRsaServerCert(t *testing.T) {
	t.Log("Testing GenerateServerCert")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}

	cert, privateKey, err := certmanager.GenerateServerCert("myhub.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", cert, privateKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateEcdsaServerCert(t *testing.T) {
	t.Log("Testing GenerateEcdsaServerCert")

	caCert, caKey, err := certmanager.GenerateEcdsaCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}

	cert, privateKey, err := certmanager.GenerateEcdsaServerCert("myhub.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", cert, privateKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateRsaClientCert(t *testing.T) {
	t.Log("Testing GenerateClientCert")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}

	cert, privateKey, err := certmanager.GenerateClientCert("mydatacenter.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", cert, privateKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateEcdsaClientCert(t *testing.T) {
	t.Log("Testing GenerateEcdsaClientCert")

	caCert, caKey, err := certmanager.GenerateEcdsaCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}

	cert, privateKey, err := certmanager.GenerateClientCert("mydatacenter.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", cert, privateKey)
	} else {
		t.Error(err)
	}
}

func TestGenerateRsaServerAndClientCert(t *testing.T) {
	t.Log("Testing GenerateClientCert")

	caCert, caKey, err := certmanager.GenerateCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("CA CERT:")
	//fmt.Println(caCert)
	//fmt.Println("CA KEY:")
	//fmt.Println(caKey)

	scert, sprivateKey, err := certmanager.GenerateServerCert("myhub.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", scert, sprivateKey)
	} else {
		t.Error(err)
	}
	//fmt.Println("SERVER CERT:")
	//fmt.Println(scert)
	//fmt.Println("SERVER KEY:")
	//fmt.Println(sprivateKey)

	ccert, cprivateKey, err := certmanager.GenerateClientCert("mydatacenter.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", ccert, cprivateKey)
	} else {
		t.Error(err)
	}
	//fmt.Println("CLIENT CERT:")
	//fmt.Println(ccert)
	//fmt.Println("CLIENT KEY:")
	//fmt.Println(cprivateKey)
}

func TestGenerateEcdsaServerAndClientCert(t *testing.T) {
	t.Log("Testing GenerateClientCert")

	caCert, caKey, err := certmanager.GenerateEcdsaCA("myhub-ca.com")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("ECDSA CA CERT:")
	fmt.Println(caCert)
	fmt.Println("ECDSA CA KEY:")
	fmt.Println(caKey)

	scert, sprivateKey, err := certmanager.GenerateEcdsaServerCert("myhub.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", scert, sprivateKey)
	} else {
		t.Error(err)
	}
	fmt.Println("ECDSA SERVER CERT:")
	fmt.Println(scert)
	fmt.Println("ECDSA SERVER KEY:")
	fmt.Println(sprivateKey)

	ccert, cprivateKey, err := certmanager.GenerateEcdsaClientCert("mydatacenter.com", caCert, caKey)

	if err == nil {
		t.Logf("\ncert: %s\n private key: %s\n", ccert, cprivateKey)
	} else {
		t.Error(err)
	}
	fmt.Println("ECDSA CLIENT CERT:")
	fmt.Println(ccert)
	fmt.Println("ECDSA CLIENT KEY:")
	fmt.Println(cprivateKey)
}

func TestGenerateEcdsaSelfsignCert(t *testing.T) {
	t.Log("Testing GenerateSelfsignCert")
	scert, sprivateKey, err := certmanager.GenerateEcdsaSelfsignCert("ankr-metering")

        if err == nil {
                t.Logf("\ncert: %s\n private key: %s\n", scert, sprivateKey)
        } else {
                t.Error(err)
        }
        fmt.Println("ECDSA Metering CERT:")
        fmt.Println(scert)
        fmt.Println("ECDSA Metering KEY:")
        fmt.Println(sprivateKey)
}

func TestIsCommonNameMatch(t *testing.T) {
	t.Log("Testing IsCommonNameMatch")

	result := certmanager.IsCommomNameMatch("mydatacenter.com", ECDSA_CLIENT_CERT)

	if !result {
                t.Error(nil)
	}
}
