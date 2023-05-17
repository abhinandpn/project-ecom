package twillio

import (
	"github.com/abhinandpn/project-ecom/pkg/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

func TwillioOtpSent(phoneNumber string) (string, error) {

	// Create A twillio clint

	User := config.GetCofig().ACCOUNTSID
	Pass := config.GetCofig().AUTHTOKEN
	ServiceId := config.GetCofig().SERVICEID

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: User,
		Password: Pass,
	})

	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(ServiceId, params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
func TwilioVerifyOTP(phoneNumber string, code string) error {

	//create a twilio client with twilio details
	password := config.GetCofig().AUTHTOKEN
	userName := config.GetCofig().ACCOUNTSID
	seviceSid := config.GetCofig().SERVICEID

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Password: password,
		Username: userName,
	})

	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo(phoneNumber)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(seviceSid, params)
	if err != nil {
		return err
	} else if *resp.Status == "approved" {
		return nil
	}

	return nil
}
