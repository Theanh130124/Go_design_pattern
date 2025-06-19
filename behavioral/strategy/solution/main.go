package solution

import "fmt"

type Notifier interface { //interface là truyền gì cũng đc
	Send(message string) // hàm send dùng chung
}

//Tạo thêm 1 structure mới và func tương ứng cho structure

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) { //Chỉ cần struct dùng đủ phương thức interface thì nó sẽ tự implement
	fmt.Printf("Sending message: %s (Sender: Email)", message)
}

type SMSNotifier struct{}

func (SMSNotifier) Send(message string) {
	fmt.Printf("Sending message: %s (Sender: SMS)", message)
}

type NotificationService struct {
	notifier Notifier //chứa 1 struct khác cũng được vì Notifer(là interface truyen gi cũng đc)
}

func (s NotificationService) SendNotification(message string) {
	s.notifier.Send(message) //dùng là method Send
}

func main() {
	s := NotificationService{
		notifier: EmailNotifier{}, //truyền notifier là một EmailNotifier
	}

	s.SendNotification("Hello world")
}
