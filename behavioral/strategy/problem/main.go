package problem

import "fmt"

type NotificationService struct {
	notifierType string
}

// SendNotification tên hàm   , dL gửi (s NotificationService) , DL trả về message string
// Vấn đề đầu tiên là 1 hàm xử lý chung 1 tác vụ -> phải tăng if else ... (không nên)
func (s NotificationService) SendNotification(message string) {
	if s.notifierType == "email" {
		fmt.Printf("Sending message: %s (Sender: Email)\n", message)
	} else if s.notifierType == "sms" {
		fmt.Printf("Sending message: %s (Sender: SMS)\n", message)
	}
}

func main() {
	s := NotificationService{notifierType: "email"}
	s.SendNotification("Hello world")
}
