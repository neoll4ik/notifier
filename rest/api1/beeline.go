package api1

import (
	"fmt"
	"strings"
	"net/url"
	"net/http"
	"crypto/tls"
	"io/ioutil"
	"golang.org/x/text/encoding/charmap"
	. "github.com/aavzz/notifier/setup/syslog"
	. "github.com/aavzz/notifier/setup/cfgfile"
	
)

func sendMessageBeeline(numbers string, message string) {
  
	c, err := CfgFileContent()
	if err != nil {
		SysLog.Err(err.Error())	
	} else {
		msg, err := charmap.Windows1251.NewEncoder().String(message)
		if err != nil {
			SysLog.Err(err.Error())	
		} else {
			parameters := url.Values{
				"user": {*c.Beeline.Login},
				"pass": {*c.Beeline.Password},
				"sender": {*c.Beeline.Sender},
				"action": {"post_sms"},
				"target": {numbers},
				"message": {msg},
			}

			url := "https://beeline.amega-inform.ru/sendsms/"
			req, err := http.NewRequest("POST", url, strings.NewReader(parameters.Encode())) 
			if err != nil {
				SysLog.Err(err.Error())
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=windows-1251")
	
			tr := &http.Transport{
		        	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			c := &http.Client{Transport: tr}

		  	resp, err := c.Do(req)
			if err != nil {
				SysLog.Err(err.Error())
			}
			if resp != nil {
				defer resp.Body.Close()
			}
			body, err := ioutil.ReadAll(resp.Body)
			bodyString := string(body)
			SysLog.Info(fmt.Sprintf("Post result: %v", resp.Status))
			SysLog.Info(fmt.Sprintf("Post result: %v", bodyString))
			SysLog.Info(fmt.Sprintf("Post result: %v", strings.NewReader(parameters.Encode())))

		}
	}
}
