package test

import (
	"log"
	"testing"

	"github.com/Cheep2Workshop/proj-web/models"
	"github.com/Cheep2Workshop/proj-web/models/repo"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"
)

var testconfig = repo.DbConfig{
	Account:  "root",
	Password: "QMKAJNNjNK9vBO88",
	Ip:       "127.0.0.1",
	Port:     "3306",
	DbName:   "mysql",
}

type OrmSuite struct {
	suite.Suite
	client *repo.DbClient
}

// initial user cases
var usercases = []models.User{
	{Name: "Admin", Email: "Admin@gmail.com", Password: "Admin123456", Admin: true},
	{Name: "Bob", Email: "Bob@gmail.com", Password: "Bob123456", Admin: false},
	{Name: "Charly", Email: "Charly@gmail.com", Password: "Charly123456", Admin: false},
}

func (t *OrmSuite) SetupSuite() {
	client, err := testconfig.ConnectDb(false)
	if err != nil {
		log.Fatal(err)
	}
	// init client with tables and users
	client.Init()
	for _, user := range usercases {
		client.Signup(user)
	}
	t.client = client

	client.SaveLoginLog("Admin@gmail.com")
	client.SaveLoginLog("Bob@gmail.com")
}

func (t *OrmSuite) TearDownSuite() {
	// drop all tables
	log.Println("Tear down suite")
	t.client.DropAllTables()
}

func (t *OrmSuite) TestSignup() {
	log.Println("Test signup")
	var user models.User
	// test duplicate user signup
	user = usercases[0]
	result, err := t.client.Signup(user)
	require.False(t.T(), result)
	require.Error(t.T(), err)
	// test new user signup
	user = models.User{Name: "ABC", Email: "ABC@gmail.com", Password: "123456", Admin: false}
	result, err = t.client.Signup(user)
	require.True(t.T(), result)
	require.NoError(t.T(), err)
}

func (t *OrmSuite) TestLogin() {
	log.Println("Test login")
	user := usercases[0]
	req := repo.LoginReq{
		Email:    user.Email,
		Password: user.Password,
	}
	_, err := t.client.Login(req)
	require.NoError(t.T(), err)
	_, err = t.client.Login(repo.LoginReq{
		Email:    user.Email,
		Password: "",
	})
	require.Error(t.T(), err)
}

type SetUserInfo struct {
	Name     string
	Password string
}

func (t *OrmSuite) TestSetUser() {
	var err error
	var req repo.SetUserReq
	// valid cases
	infos := []SetUserInfo{
		{Name: "", Password: "BillyYO"},
		{Name: "Billy", Password: ""},
		{Name: "Bob", Password: "Boby123456"},
	}
	user := &usercases[1]
	for _, info := range infos {
		req = repo.SetUserReq{
			Email:           user.Email,
			TargetEmail:     user.Email,
			ChangedName:     info.Name,
			ChangedPassword: info.Password,
		}
		err = t.client.SetUser(req)
		// change password
		if len(info.Password) > 0 {
			user.Password = info.Password
		}
		require.NoError(t.T(), err)
	}
	// not matched email or password
	badUser := models.User{
		Name:     "Billy",
		Email:    "Billy@gmail.com",
		Password: "",
		Admin:    false,
	}
	req = repo.SetUserReq{
		Email:           badUser.Email,
		TargetEmail:     badUser.Email,
		ChangedName:     "Billy",
		ChangedPassword: "",
	}
	err = t.client.SetUser(req)
	require.Error(t.T(), err)
	// both changed name/email are empty
	req = repo.SetUserReq{
		Email:           user.Email,
		TargetEmail:     user.Email,
		ChangedName:     "",
		ChangedPassword: "",
	}
	err = t.client.SetUser(req)
	require.Error(t.T(), err)
}

func (t *OrmSuite) TestSaveLoginLog() {
	require.NoError(t.T(), t.client.SaveLoginLog("Bob@gmail.com"))
}

func (t *OrmSuite) TestGetLoginLog() {
	results, err := t.client.GetLoginLogs("Bob@gmail.com")
	for _, result := range results {
		log.Println(result)
	}
	require.NoError(t.T(), err)
	require.GreaterOrEqual(t.T(), len(results), 1)
}

func (t *OrmSuite) TestDeleteUser() {
	charly := usercases[2]
	req := repo.DeleteUserReq{
		Email:       usercases[0].Email,
		DeleteEmail: charly.Email,
	}
	err := t.client.DeleteUser(req)
	require.NoError(t.T(), err)
}

func (t *OrmSuite) TestCheckAuth() {
	user := usercases[0]
	ok, err := t.client.CheckAuth(user.Name, user.Email)
	require.NoError(t.T(), err)
	require.True(t.T(), ok)
}

// start suite test
func TestOrm(t *testing.T) {
	log.Println("Run suite")
	suite.Run(t, new(OrmSuite))
}
