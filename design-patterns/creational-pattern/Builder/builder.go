package builder

import "fmt"

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) *NotificationBuilder {
	nb.Title = title
	return nb
}

func (nb *NotificationBuilder) SetSubTitle(subtitle string) *NotificationBuilder {
	nb.Subtitle = subtitle
	return nb
}

func (nb *NotificationBuilder) SetImage(image string) *NotificationBuilder {
	nb.Image = image
	return nb
}

func (nb *NotificationBuilder) SetMessage(msg string) *NotificationBuilder {
	nb.Message = msg
	return nb
}

func (nb *NotificationBuilder) SetIcon(icon string) *NotificationBuilder {
	nb.Icon = icon
	return nb
}

func (nb *NotificationBuilder) SetPriority(pri int) *NotificationBuilder {
	nb.Priority = pri
	return nb
}

func (nb *NotificationBuilder) SetType(notType string) *NotificationBuilder {
	nb.NotType = notType
	return nb
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	if nb.Icon != "" && nb.Subtitle == "" {
		return nil, fmt.Errorf("you need to specify a subtitle when using an icon")
	}
	if nb.Priority > 5 {
		return nil, fmt.Errorf("priority must be 0 to 5")
	}
	return &Notification{
		title:    nb.Title,
		subtitle: nb.Subtitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
