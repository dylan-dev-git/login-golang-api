package token

import (
	"errors"
	"net/http"
	"encoding/json"

	"github.com/dylan-dev-git/login-microservice-api/env"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	errTokenIssueFailed   					= errors.New("Token issue failed.")
	errTokenMicroserviceFailed      = errors.New("Token microservice calling failed.")
	errTokenVerifyFailed  					= errors.New("Token verification failed.")
)

type token struct {
	URL string 
	API_SERVER string
	API_KEY string
	API_VALUE string
	TOKEN string
}

func NewToken() *token {
	obj := token{}
	obj.API_KEY = "TOKENAPIKEYHERE"
	obj.API_VALUE = "TOKENAPIVALUEHERE"
	obj.API_SERVER = "http://"+env.VMIP+":8501/api/token"
	return &obj
}

func (obj *token) GetToken(userid string) (map[string]interface{}, error) {
	obj.URL = obj.API_SERVER +"/getToken/"+userid
	client := &http.Client{}
	req, _ := http.NewRequest("GET", obj.URL, nil)
	req.Header.Add(obj.API_KEY,obj.API_VALUE)
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(errTokenMicroserviceFailed.Error())
	}
	defer resp.Body.Close()

	var tokenResult bson.M
	json.NewDecoder(resp.Body).Decode(&tokenResult)
	if tokenResult["status"] != true {
		return nil, errors.New(errTokenIssueFailed.Error())
	}
	return map[string]interface{} {
		"token": tokenResult["token"].(string),
		"userid": tokenResult["userid"].(string)}, nil
}

func (obj *token) GetTokenMetaData(token string) (string, error) {
	obj.URL = obj.API_SERVER+"/getTokenMetaData"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", obj.URL, nil)
	req.Header.Add(obj.API_KEY,obj.API_VALUE)
	req.Header.Add("Authorization",token)
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New(errTokenMicroserviceFailed.Error())
	}
	defer resp.Body.Close()
	var tokenMetaData bson.M
	json.NewDecoder(resp.Body).Decode(&tokenMetaData)
	if tokenMetaData["status"] != true {
		return "", errors.New(errTokenVerifyFailed.Error())
	}
	return tokenMetaData["userid"].(string), nil
}