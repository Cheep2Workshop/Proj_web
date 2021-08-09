package test

import (
	"log"
	"testing"

	utils "github.com/Cheep2Workshop/proj-web/utils/encryption"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EncryptionSuite struct {
	suite.Suite
	params *utils.EncodeParams
}

func (t *EncryptionSuite) SetupSuite() {
	t.params = utils.DefaultParams
}

const (
	Salt = "9vxvwau32kxcuYE8iLojAg"
)

func (t *EncryptionSuite) TestMatch() {
	paramHash, keyHash, err := utils.GenerateFromPassword("Password123", Salt, t.params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(paramHash)
	log.Println(keyHash)

	var match bool
	match, err = utils.ComparePasswordAndHash("Password", keyHash, paramHash)
	require.False(t.T(), match)

	match, err = utils.ComparePasswordAndHash("Password123", keyHash, paramHash)
	require.True(t.T(), match)
}

func TestEncryption(t *testing.T) {
	log.Println("Run suite")
	suite.Run(t, new(EncryptionSuite))
}
