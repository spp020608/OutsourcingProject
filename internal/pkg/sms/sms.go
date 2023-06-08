package sms

import (
	"encoding/json"
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

const (
	//阿里云平台自己的AK AccessKey
	accessKeyId     = "LTAI5t6Kzp9gKzB4GXFnAv26"
	accessKeySecret = "LkY5zvYh2AusLqg9SoqfAZe25jOyGK"
	templateCode    = "SMS_460895074"
	signName        = "YuSheng"
	smsDomainName   = "dysmsapi.aliyuncs.com"
)

func SendCode(phone string) (code string, ok bool) {
	code = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	client, err := creatClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if err != nil {
		fmt.Println(err.Error())
		return "", false
	}
	bCode, _ := json.Marshal(map[string]interface{}{
		"code": code,
	})

	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		TemplateCode:  tea.String(templateCode),
		SignName:      tea.String(signName),
		TemplateParam: tea.String(string(bCode)),
	}

	sendSmsResponse, _ := client.SendSms(sendSmsRequest)
	if *sendSmsResponse.Body.Code == "OK" {
		zap.S().Infof("发送给手机号%s的短信验证码成功【%s】", phone, code)
		return code, true
	}
	return "", false
}

func creatClient(accessKeyId *string, accessKeySecret *string) (client *dysmsapi.Client, err error) {

	config := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
	}
	config.Endpoint = tea.String(smsDomainName)
	client = &dysmsapi.Client{}
	client, err = dysmsapi.NewClient(config)
	return client, err
}
