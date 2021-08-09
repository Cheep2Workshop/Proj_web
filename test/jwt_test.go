package test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/utils"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type JwtSuite struct {
	suite.Suite
}

var (
	issuer = "Test"
	secret = "HelloWorld!"
)

func (t *JwtSuite) TestGenerateToken() {
	user := models.User{
		Name:  "Alice",
		Email: "Alice@gmail.com",
	}

	claim := utils.GenerateClaims(user, issuer)
	token, err := utils.GenerateToken(claim, secret)
	require.NoError(t.T(), err)
	j, err := json.Marshal(claim)
	require.NoError(t.T(), err)
	// print result
	log.Println(string(j))
	log.Println(token)
}

func (t *JwtSuite) TestParseToken() {
	user := models.User{
		Name:  "Alice",
		Email: "Alice@gmail.com",
	}

	claim := utils.GenerateClaims(user, issuer)
	token, err := utils.GenerateToken(claim, secret)
	require.NoError(t.T(), err)
	claims, err := utils.ParseToken(token, secret)
	require.NoError(t.T(), err)
	require.Equal(t.T(), user.Name, claims.Name)
	require.Equal(t.T(), user.Email, claims.Email)
}

func TestJwt(t *testing.T) {
	suite.Run(t, new(JwtSuite))
}
