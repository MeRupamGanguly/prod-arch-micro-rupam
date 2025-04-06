package transports

import (
	"encoding/json"
	"net/http"
	"net/url"
	"rupx/common"
	"rupx/user/domain"
)

type AddReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AddResp struct {
	Err error `json:"error,omitempty"`
}

func DecodeAddReq(r *http.Request, w http.ResponseWriter) (name string, email string, err error) {
	var req AddReq
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return "", "", err
	}
	return req.Name, req.Email, nil
}
func EncodeAddRes(r *http.Request, w http.ResponseWriter, err error) {
	cResp := common.CustomRespCreator(nil, err)
	json.NewEncoder(w).Encode(cResp)
}

type GetReq struct {
	ID string `query:"id"`
}

type GetResp struct {
	User domain.User `json:"user"`
	Err  error       `json:"error,omitempty"`
}

func DecodeGetReq(r *http.Request, w http.ResponseWriter) (id string, err error) {
	var req GetReq
	u, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return "", err
	}
	req.ID = u["id"][0]
	return req.ID, nil
}
func EncodeGetRes(r *http.Request, w http.ResponseWriter, user domain.User, err error) {
	getIdResp := GetResp{User: user}
	cResp := common.CustomRespCreator(getIdResp, err)
	json.NewEncoder(w).Encode(cResp)
}

type ListReq struct {
}

type Listesp struct {
	Users []domain.User `json:"users"`
	Err   error         `json:"error,omitempty"`
}

func DecodeListReq(r *http.Request, w http.ResponseWriter) (err error) {
	return
}
func EncodeListRes(r *http.Request, w http.ResponseWriter, users []domain.User, err error) {
	getIdResp := Listesp{Users: users}
	cResp := common.CustomRespCreator(getIdResp, err)
	json.NewEncoder(w).Encode(cResp)
}
