package keystore

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"encoding/json"
	"time"

	"golang.org/x/crypto/pbkdf2"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/sha3"
	wallet "github.com/Ankr-network/dccn-common/wallet"
)

const (
	keyHeaderKDF = "scrypt"

	// StandardScryptN is the N parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptN = 1 << 18

	// StandardScryptP is the P parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptP = 1

	// LightScryptN is the N parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptN = 1 << 12

	// LightScryptP is the P parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptP = 6

	scryptR     = 8
	scryptDKLen = 32

	keyJSONVersion = 3
)

type CryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherparamsJSON       `json:"cipherparams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfparams"`
	MAC          string                 `json:"mac"`
}

type cipherparamsJSON struct {
	IV string `json:"iv"`
}

type EncryptedKeyJSONV3 struct {
	Name           string     `json:"name,omitempty"`
	Address        string     `json:"address"`
	PublicKey      string     `json:"publickey"`
	Crypto         CryptoJSON `json:"crypto"`
	KeyJSONVersion int        `json:"version"`
}

func EncryptDataV3(data, auth []byte, scryptN, scryptP int) (CryptoJSON, error) {

	salt := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	derivedKey, err := scrypt.Key(auth, salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return CryptoJSON{}, err
	}
	encryptKey := derivedKey[:16]

	iv := make([]byte, aes.BlockSize) // 16
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	cipherText, err := aesCTRXOR(encryptKey, data, iv)
	if err != nil {
		return CryptoJSON{}, err
	}
	mac := Keccak256(derivedKey[16:32], cipherText)

	scryptParamsJSON := make(map[string]interface{}, 5)
	scryptParamsJSON["n"] = scryptN
	scryptParamsJSON["r"] = scryptR
	scryptParamsJSON["p"] = scryptP
	scryptParamsJSON["dklen"] = scryptDKLen
	scryptParamsJSON["salt"] = hex.EncodeToString(salt)
	cipherParamsJSON := cipherparamsJSON{
		IV: hex.EncodeToString(iv),
	}

	cryptoStruct := CryptoJSON{
		Cipher:       "aes-128-ctr",
		CipherText:   hex.EncodeToString(cipherText),
		CipherParams: cipherParamsJSON,
		KDF:          keyHeaderKDF,
		KDFParams:    scryptParamsJSON,
		MAC:          hex.EncodeToString(mac),
	}
	return cryptoStruct, nil
}

func Keccak256(data ...[]byte) []byte {
	d := sha3.NewLegacyKeccak256()
	for _, b := range data {
		d.Write(b)
	}
	return d.Sum(nil)
}

func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
	// AES-128 is selected due to size of encryptKey.
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesBlock, iv)
	outText := make([]byte, len(inText))
	stream.XORKeyStream(outText, inText)
	return outText, err
}

func DecryptDataV3(cryptoJson CryptoJSON, auth string) ([]byte, error) {
	if cryptoJson.Cipher != "aes-128-ctr" {
		return nil, fmt.Errorf("Cipher not supported: %v", cryptoJson.Cipher)
	}
	mac, err := hex.DecodeString(cryptoJson.MAC)
	if err != nil {
		fmt.Printf("DecodeString MAC error: %s", err)
		return nil, err
	}

	iv, err := hex.DecodeString(cryptoJson.CipherParams.IV)
	if err != nil {
		fmt.Printf("DecodeString CipherParams error: %s", err)
		return nil, err
	}

	cipherText, err := hex.DecodeString(cryptoJson.CipherText)
	if err != nil {
		fmt.Printf("DecodeString CipherText error: %s", err)
		return nil, err
	}

	derivedKey, err := getKDFKey(cryptoJson, auth)
	if err != nil {
		fmt.Printf("getKDFKey error: %s", err)
		return nil, err
	}

	calculatedMAC := Keccak256(derivedKey[16:32], cipherText)
	if !bytes.Equal(calculatedMAC, mac) {
		return nil, errors.New("could not decrypt key with given passphrase")
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		fmt.Printf("aesCTRXOR error: %s", err)
		return nil, err
	}
	return plainText, err
}

/*
input: keystore_name(can be empty), priv_key, password
output: keystore json string, ankr token address, error
*/
func PrivKeyToKeyStoreJsonString(name, priv_key, password string) (json_key string, address string, err error) {
        pub, err := wallet.GetPublicKeyByPrivateKey(priv_key)
        if err != nil {
                return "", "", err
        }

	address, err = wallet.GetAddressByPublicKey(pub)
        if err != nil {
                return "", "", err
        }

        cryptoStruct, err := EncryptDataV3([]byte(priv_key),
                []byte(password), StandardScryptN, StandardScryptP)

        if err == nil {
                //fmt.Printf("\ncryptoStruct: %v\n ", cryptoStruct)
        } else {
                return "", "", err
        }

	encryptedKeyJSONV3 := EncryptedKeyJSONV3{
			name,
			address,
			pub,
			cryptoStruct,
			keyJSONVersion,
		}

	jsonKey, err := json.Marshal(encryptedKeyJSONV3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err.Error())
		return "", "", err
	}

	return string(jsonKey), address, nil
}

/*
input: keystore json string, password
output: priv key, error
*/
func KeyStoreJsonStringToPrivKey(json_string, password string) (priv_key string, err error) {
	var key EncryptedKeyJSONV3
	err = json.Unmarshal([]byte(json_string), &key)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err.Error())
		return "", err
	}


        priv, err := DecryptDataV3(key.Crypto, password)
        if err == nil {
                //fmt.Printf("\nresult: %s\n", string(priv))
        } else {
                return "", err
        }

	return string(priv), nil
}

/*
input: keystore json string, old password, new password
output: new keystore json string, ankr token address, error
*/
func ResetPasswordKeyStoreJsonString(json_string, old_password, new_password string) (json_key, address string, err error) {
	var key EncryptedKeyJSONV3
        err = json.Unmarshal([]byte(json_string), &key)
        if err != nil {
                fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err.Error())
                return "", "", err
        }


        priv, err := DecryptDataV3(key.Crypto, old_password)
        if err == nil {
                //fmt.Printf("\nresult: %s\n", string(priv))
        } else {
                return "", "", err
        }

	jsonkey, address, err := PrivKeyToKeyStoreJsonString(key.Name,
                string(priv), new_password)
        if err == nil {
                //fmt.Printf(jsonkey)
        } else {
                return "", "", err
        }

	return jsonkey, address, nil
}

func getKDFKey(cryptoJSON CryptoJSON, auth string) ([]byte, error) {
	authArray := []byte(auth)
	salt, err := hex.DecodeString(cryptoJSON.KDFParams["salt"].(string))
	if err != nil {
		return nil, err
	}
	dkLen := ensureInt(cryptoJSON.KDFParams["dklen"])

	if cryptoJSON.KDF == keyHeaderKDF {
		n := ensureInt(cryptoJSON.KDFParams["n"])
		r := ensureInt(cryptoJSON.KDFParams["r"])
		p := ensureInt(cryptoJSON.KDFParams["p"])
		return scrypt.Key(authArray, salt, n, r, p, dkLen)

	} else if cryptoJSON.KDF == "pbkdf2" {
		c := ensureInt(cryptoJSON.KDFParams["c"])
		prf := cryptoJSON.KDFParams["prf"].(string)
		if prf != "hmac-sha256" {
			return nil, fmt.Errorf("Unsupported PBKDF2 PRF: %s", prf)
		}
		key := pbkdf2.Key(authArray, salt, c, dkLen, sha256.New)
		return key, nil
	}

	return nil, fmt.Errorf("Unsupported KDF: %s", cryptoJSON.KDF)
}

func ensureInt(x interface{}) int {
	res, ok := x.(int)
	if !ok {
		res = int(x.(float64))
	}
	return res
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}
