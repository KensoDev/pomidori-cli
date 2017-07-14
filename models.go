package pomidori

import "encoding/json"

type LoginDetails struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewLoginDetails(email, password string) *LoginDetails {
	return &LoginDetails{
		Email:    email,
		Password: password,
	}
}

func (loginDetails *LoginDetails) ToJson() ([]byte, error) {
	return json.Marshal(loginDetails)
}

type User struct {
	Token string `json:"token"`
}

func NewUser(jsonBody []byte) (*User, error) {
	user := &User{}
	err := json.Unmarshal(jsonBody, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

type Task struct {
	Title    string `json:"title"`
	TaskTime string `json:"taskTime"`
}

func NewTask(title, taskTime string) *Task {
	return &Task{
		Title:    title,
		TaskTime: taskTime,
	}
}

func (task *Task) ToJson() ([]byte, error) {
	return json.Marshal(task)
}
