package whatsapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var baseOTPRequstMessage = `JANGAN BERI KODE ini ke siapa pun, TERMASUK SALESMAN & PIC SIDOMUNCUL. WASPADA PENIPUAN! MASUK KE AKUN dengan kode verifikasi (OTP) #OTP#. OTP bisa lewat WA/SMS`

type Client struct {
	apiurl string
	apikey string
}

func NewWooWAClient(ApiUrl, Apikey string) *Client {
	return &Client{
		apiurl: ApiUrl,
		apikey: Apikey,
	}
}

func (cl Client) SendWA(phoneNo, txtMessages string) (err error) {
	type confi struct {
		Key   string `json:"key"`
		Phone string `json:"phone_no"`
		Pesan string `json:"message"`
	}

	Waconf := new(confi)
	Waconf.Key = cl.apikey
	Waconf.Phone = phoneNo
	Waconf.Pesan = strings.ReplaceAll(baseOTPRequstMessage, "#OTP#", txtMessages)
	jsonReq, err := json.Marshal(Waconf)
	// fmt.Println("api key ", cl.apikey)
	client := &http.Client{}
	req, err := http.NewRequest("POST", cl.apiurl, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic dXNtYW5ydWJpYW50b3JvcW9kcnFvZHJiZWV3b293YToyNjM3NmVkeXV3OWUwcmkzNDl1ZA==")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject http.Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
	return nil
}

func (cl Client) SendTransactionWA(phoneNo, txtMessages string) (err error) {
	type confi struct {
		Key   string `json:"key"`
		Phone string `json:"phone_no"`
		Pesan string `json:"message"`
	}

	Waconf := new(confi)
	Waconf.Key = cl.apikey
	Waconf.Phone = phoneNo
	Waconf.Pesan = txtMessages
	jsonReq, err := json.Marshal(Waconf)

	client := &http.Client{}
	req, err := http.NewRequest("POST", cl.apiurl, bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic dXNtYW5ydWJpYW50b3JvcW9kcnFvZHJiZWV3b293YToyNjM3NmVkeXV3OWUwcmkzNDl1ZA==")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject http.Response
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Printf("API Response as struct %+v\n", responseObject)
	return nil
}
