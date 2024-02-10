package util

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/rulanugrh/venus/internal/entity/web"
)

type SuiteInterface interface {
	Result(res *web.Success, http *http.Response)
	Login(path string, body *bytes.Buffer, res *web.Success) (*web.Success, *http.Response, error)
	Post(path string, body *bytes.Buffer, res *web.Success, token string) (*web.Success, *http.Response, error)
	Get(path string, res *web.Success, token string) (*web.Success, *http.Response, error)
	Put(path string, body *bytes.Buffer, res *web.Success, token string) (*web.Success, *http.Response, error)
	Delete(path string, res *web.Success, token string) (*http.Response, error)

}

type Suites struct {
	client *http.Client
}

func NewSuiteUtils(client *http.Client) SuiteInterface {
	return &Suites{
		client: client,
	}
}

func (s *Suites) sendRequest(method string, path string, body *bytes.Buffer, token string) (*http.Response, error) {
	bdy := bytes.NewBuffer(nil)
	if body != nil {
		bdy = body
	}

	req, err := http.NewRequest(method, path, bdy)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := s.client.Do(req)
	return resp, err
}

func (s *Suites) loginRequest(method string, path string, body *bytes.Buffer) (*http.Response, error) {
	bdy := bytes.NewBuffer(nil)
	if body != nil {
		bdy = body
	}

	req, err := http.NewRequest(method, path, bdy)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	return resp, err
}
func (s *Suites) Result(res *web.Success, http *http.Response) {
	decoder := json.NewDecoder(http.Body)
	err := decoder.Decode(&res)
	if err != nil {
		panic(err)
	}
}

func (s *Suites) Login(path string, body *bytes.Buffer, res *web.Success) (*web.Success, *http.Response, error) {
	resp, err := s.loginRequest("POST", path, body)
	if err != nil {
		return nil, nil, err
	}

	s.Result(res, resp)
	return res, resp, nil
}

func (s *Suites) Post(path string, body *bytes.Buffer, res *web.Success, token string) (*web.Success, *http.Response, error) {
	resp, err := s.sendRequest("POST", path, body, token)
	if err != nil {
		return nil, nil, err
	}

	s.Result(res, resp)
	return res, resp, nil
}

func (s *Suites) Get(path string, res *web.Success, token string) (*web.Success, *http.Response, error) {
	resp, err := s.sendRequest("GET", path, nil, token)
	if err != nil {
		return nil, nil, err
	}

	s.Result(res, resp)
	return res, resp, nil
}

func (s *Suites) Put(path string, body *bytes.Buffer, res *web.Success, token string) (*web.Success, *http.Response, error) {
	resp, err := s.sendRequest("PUT", path, body, token)
	if err != nil {
		return nil, nil, err
	}

	s.Result(res, resp)
	return res, resp, nil
}

func (s *Suites) Delete(path string, res *web.Success, token string) (*http.Response, error) {
	resp, err := s.sendRequest("DELETE", path, nil, token)
	if err != nil {
		return nil, err
	}

	s.Result(res, resp)
	return resp, nil
}
