package mock

import (
	"fmt"

	"github.com/r0612885/PapierA4/Dgraph/Events/models"
)
// NotificationService is the mock implementation of the actual notification service
type NotificationService struct{}

// Publish is the mock implementation of the notification service
func (n *NotificationService) Publish(notification *models.Notification) {
	// This would then send the notification to the notification stream
	fmt.Printf("Sending %s:%s [%s] with payload %v\n", notification.Type, notification.Target, notification.Hash, notification.Payload)
}
