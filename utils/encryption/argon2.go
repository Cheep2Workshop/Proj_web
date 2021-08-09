package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

// ref : https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go

type EncodeParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
	DefaultParams          = &EncodeParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
)

func main() {

	// Pass the plaintext password and parameters to our generateFromPassword
	// helper function.
	paramHash, keyHash, err := GenerateFromPassword("password123", "", DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(paramHash)
	fmt.Println(keyHash)

	pwd := "password"
	match, err := ComparePasswordAndHash(pwd, keyHash, paramHash)
	if err != nil {
		log.Printf("Match:%v, Err:%s\n", match, err.Error())
	}
	fmt.Printf("Pwd:%s, Match:%v\n", pwd, match)

	pwd = "password123"
	match, err = ComparePasswordAndHash(pwd, keyHash, paramHash)
	if err != nil {
		log.Printf("Match:%v, Err:%s\n", match, err.Error())
	}
	fmt.Printf("Pwd:%s, Match:%v\n", pwd, match)
}

func (p *EncodeParams) ToHash(salt string) string {
	return fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,l=%d,p=%d$%s", argon2.Version, p.memory, p.iterations, p.keyLength, p.parallelism, salt)
}

func GenerateFromPassword(password string, salt string, p *EncodeParams) (paramHash string, keyHash string, err error) {
	var saltBytes []byte
	if len(salt) <= 0 {
		// Generate a cryptographically secure random salt.
		saltBytes, err = GenerateRandomBytes(p.saltLength)
		if err != nil {
			return "", "", err
		}
		salt = base64.RawStdEncoding.EncodeToString(saltBytes)
	} else {
		// validate input of salt
		saltBytes, err = base64.RawStdEncoding.Strict().DecodeString(salt)
		if err != nil {
			return "", "", err
		}
		if len(saltBytes) != int(p.saltLength) {
			return "", "", errors.New("Length of salt bytes not equal to params.saltLength")
		}
	}

	// Pass the plaintext password, salt and parameters to the argon2.IDKey
	// function. This will generate a hash of the password using the Argon2id
	// variant.
	hash := argon2.IDKey([]byte(password), saltBytes, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Base64 encode the salt and hashed password.
	keyHash = base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	paramHash = p.ToHash(salt)

	return paramHash, keyHash, nil
}

func GenerateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func ComparePasswordAndHash(password, keyHash, paramHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, err := DecodeHash(paramHash)
	if err != nil {
		return false, err
	}
	// decode hashed key
	hash, err := base64.RawStdEncoding.DecodeString(keyHash)
	if err != nil {
		return false, err
	}
	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, errors.New("Invalid password.")
}

func DecodeHash(paramHash string) (p *EncodeParams, salt []byte, err error) {
	vals := strings.Split(paramHash, "$")
	if len(vals) != 5 {
		return nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, ErrIncompatibleVersion
	}

	p = &EncodeParams{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,l=%d,p=%d", &p.memory, &p.iterations, &p.keyLength, &p.parallelism)
	if err != nil {
		return nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	// hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	// if err != nil {
	// 	return nil, nil, nil, err
	// }
	// p.keyLength = uint32(len(hash))

	return p, salt, nil
}
