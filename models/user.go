package models

type SubscriptionTier string

const (
    Explorer SubscriptionTier = "Explorer"
    Scholar  SubscriptionTier = "Scholar"
)

type User struct {
    Name               string           `json:"name"`
    Email              string           `json:"email"`
    SubscriptionStatus SubscriptionTier `json:"subscriptionStatus"`
}
