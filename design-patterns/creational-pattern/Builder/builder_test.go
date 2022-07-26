package builder

import (
	"fmt"
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	var builder = newNotificationBuilder()
	builder.SetTitle("New Notification")
	builder.SetIcon("icon.png")
	builder.SetSubTitle("Subtitle")
	builder.SetImage("image.jpg")
	builder.SetPriority(4)
	builder.SetMessage("Basic notification")
	builder.SetType("Alert")

	notif, err := builder.Build()
	if err != nil {
		fmt.Println("Error creating the notification,", err)
		return
	}
	fmt.Printf("Notification: %+v", notif)
}
