package sms

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"driver/conf"
	"driver/util/common"
)

type Shumi struct {
}

func (s *Shumi) Send(phoneNumber string, message string) error {
	if conf.SMS_SWITCH {
		message = s.appendSignature(message)
		err := s.request(phoneNumber, message)
		return err
	} else {
		fmt.Println(phoneNumber + " " + message + " SMS_SWITCH off, not send")
		return nil
	}
}

func (s *Shumi) request(phoneNumber string, message string) error {
	options := s.getDefaults()
	options["mobile"] = phoneNumber
	options["content"] = base64.StdEncoding.EncodeToString([]byte(message))

	urlValues := url.Values{}
	for key, value := range options {
		urlValues.Set(key, value)
	}

	resp, err := http.PostForm("http://api.shumi365.com:8090/sms/send.do", urlValues)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(phoneNumber + " " + message + " res:" + string(body))

	return err
}

func (s *Shumi) getDefaults() map[string]string {
	timespan := time.Now().Format("20060102150405")
	secret := conf.SHUMI_SECRET
	password := strings.ToUpper(common.MD5(secret + timespan))
	return map[string]string{
		"userid":   conf.SHUMI_USER_ID,
		"timespan": timespan,
		"pwd":      password,
		"msgfmt":   "UTF8",
	}
}

func (s *Shumi) appendSignature(content string) string {
	return content + conf.SHUMI_SIGNATURE
}

func (s *Shumi) insertSignature(content string) string {
	return conf.SHUMI_SIGNATURE + content
}
