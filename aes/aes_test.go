package aes

import (
	"crypto/rand"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	tempDir, _ := ioutil.TempDir("", "encription_tests")

	keyText := make([]byte, 32) // Rand 32 bytes
	rand.Read(keyText)
	InitializeBlock(keyText)

	input1 := []byte("ooooooooooooooooooooo")
	input2 := []byte("4534862878942768574983y")
	input3 := make([]byte, 1024*1024*8)
	rand.Read(input3)

	// Write some files
	ioutil.WriteFile(tempDir+"/test/dat1", input1, 0777)
	ioutil.WriteFile(tempDir+"/test/dat2", input2, 0777)
	ioutil.WriteFile(tempDir+"/test/dat3", input3, 0777)

	streamEncrypter(tempDir + "/test/dat1")
	streamEncrypter(tempDir + "/test/dat2")
	streamEncrypter(tempDir + "/test/dat3")

	// Decrypt

	streamDecrypter(tempDir + "/test/ZGF0MQ==.enc")
	streamDecrypter(tempDir + "/test/ZGF0Mg==.enc")
	streamDecrypter(tempDir + "/test/ZGF0Mw==.enc")

	f1, _ := ioutil.ReadFile(tempDir + "/test/dat1")
	f2, _ := ioutil.ReadFile(tempDir + "/test/dat2")
	f3, _ := ioutil.ReadFile(tempDir + "/test/dat3")
	assert.Equal(t, input1, f1, "should be equal")
	assert.Equal(t, input2, f2, "should be equal")
	assert.Equal(t, input3, f3, "should be equal")
}
