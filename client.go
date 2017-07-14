package pomidori

import (
	"bytes"
	"fmt"
	"github.com/segmentio/go-prompt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Token string
}

func NewClient() *Client {
	return &Client{
		Token: "",
	}
}

func NewRegisteredClient(token string) *Client {
	return &Client{
		Token: token,
	}
}

func (c *Client) CreateTask(title, taskTime string) error {
	client := &http.Client{}
	task := NewTask(title, taskTime)
	json, _ := task.ToJson()
	fmt.Println(string(json))
	req, err := http.NewRequest("POST", "http://localhost:4040/api/task", bytes.NewBuffer(json))

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	fmt.Println(resp)

	responseBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(responseBody))

	defer resp.Body.Close()

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Register() (string, error) {
	email := prompt.String("Type in your email")
	password := prompt.PasswordMasked("Type in your password")
	passwordConfirmation := prompt.PasswordMasked("Type in your password again")

	if password != passwordConfirmation {
		return "", fmt.Errorf("Your passwrods did not match")
	}

	loginDetails := NewLoginDetails(email, password)
	json, err := loginDetails.ToJson()

	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:4040/api/auth/register", "application/json", bytes.NewBuffer(json))

	if err != nil {
		return "", fmt.Errorf("The server responded with HTTP code: %s", resp.Status)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	user, err := NewUser(responseBody)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return user.Token, nil
}
