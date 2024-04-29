package builder

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	var builder = newNotificationBuilder()
	builder.
		SetTitle("New Notification").
		SetIcon("icon.png").
		SetSubTitle("Subtitle").
		SetImage("image.jpg").
		SetPriority(4).
		SetMessage("Basic notification").
		SetType("Alert")

	notif, err := builder.Build()
	if err != nil {
		fmt.Println("Error creating the notification,", err)
		return
	}
	fmt.Printf("Notification: %+v", notif)
}
