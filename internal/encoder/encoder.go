package encoder

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd []byte) []byte {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return nil
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return hash
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}

// func Secret() ([]byte, error) {

// 	// generates a 16-byte slice where the key will be stored
// 	key := make([]byte, 16)

// 	// fill the key with random values ​​to encrypt differently
// 	if _, err := rand.Read(key); err != nil {
// 		return nil, err
// 	}

// 	return key, nil
// }

// func createHash(key string) string {
// 	hasher := md5.New()
// 	hasher.Write([]byte(key))
// 	return hex.EncodeToString(hasher.Sum(nil))
// }

// func Encrypt(data []byte, passphrase string) string {
// 	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonce := make([]byte, gcm.NonceSize())
// 	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
// 		panic(err.Error())
// 	}
// 	ciphertext := gcm.Seal(nonce, nonce, data, nil)
// 	return base64.URLEncoding.EncodeToString(ciphertext)
// }

// func Decrypt(data []byte, passphrase string) []byte {
// 	stringData, err := base64.URLEncoding.DecodeString(string(data))

// 	dataToByte := []byte(stringData)
// 	key := []byte(createHash(passphrase))
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	nonceSize := gcm.NonceSize()
// 	nonce, ciphertext := dataToByte[:nonceSize], dataToByte[nonceSize:]
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return nil
// 	}
// 	return plaintext
// }
