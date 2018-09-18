package translater

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

// Translater is a struct to translate
type Translater struct {
	sess *session.Session
}

// Init is a function to initialize aws session
func (translater *Translater) Init() error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	translater.sess = sess
	return nil
}

// TranslateString is a function to translate origin from en to ja
func (translater *Translater) TranslateString(origin string) (string, error) {
	svc := translate.New(translater.sess)
	textInput := translate.TextInput{}
	textInput.SetSourceLanguageCode("en")
	textInput.SetTargetLanguageCode("ja")
	textInput.SetText(origin)

	result, err := svc.Text(&textInput)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	translated := *result.TranslatedText
	return translated, nil
}
