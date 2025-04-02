package pkg

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"reflect"

	"golang.org/x/crypto/argon2"
)

type Salt struct {
	salt    []byte
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func CoffeeSalt() *Salt {
	salt, _ := genByte(32)
	return &Salt{
		salt:    salt,
		time:    3,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}
}

func genByte(keylen uint32) ([]byte, error) {
	bytes := make([]byte, keylen)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func Hash(password string, coffeeSalt *Salt) (string, error) {
    var salt []byte
    if coffeeSalt.salt != nil {
        salt = coffeeSalt.salt
    } else {
        var err error
        salt, err = genByte(coffeeSalt.keyLen)
        if err != nil {
            return "", fmt.Errorf("error generating salt: %v", err)
        }
    }

    hashedByte := argon2.IDKey([]byte(password), salt,
        coffeeSalt.time, coffeeSalt.memory, coffeeSalt.threads, coffeeSalt.keyLen)

    hashedWithSalt := append(hashedByte, salt...)
    hashedString := base64.StdEncoding.EncodeToString(hashedWithSalt)

    return hashedString, nil
}
func verifyHashed(password, hashedPassword string) (bool, error) {
    decodedHashed, err := base64.StdEncoding.DecodeString(hashedPassword)
    if err != nil {
        return false, fmt.Errorf("error decoding hashed password: %v", err)
    }

    coffeSalt := CoffeeSalt()
    salt := decodedHashed[len(decodedHashed)-int(coffeSalt.keyLen):]
    coffeSalt.salt = salt

    hashedFromInput, err := Hash(password, coffeSalt)
    if err != nil {
        return false, fmt.Errorf("error generating hash from input: %v", err)
    }

    decodedHashedFromInput, err := base64.StdEncoding.DecodeString(hashedFromInput)
    if err != nil {
        return false, fmt.Errorf("error decoding hash from input: %v", err)
    }

    return reflect.DeepEqual(decodedHashed, decodedHashedFromInput), nil
}
