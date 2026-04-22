package model

type OnboardingRequest struct {

	UserID string `json:"user_id"`

	VPA string `json:"vpa"`

	KYCLevel string `json:"kyc_level"`

	Balance float64 `json:"balance"`
}

type OnboardingEvent struct {

	EventID string `json:"event_id"`

	UserID string `json:"user_id"`

	VPA string `json:"vpa"`

	KYCLevel string `json:"kyc_level"`

	Balance float64 `json:"balance"`

	Status string `json:"status"`

	Timestamp string `json:"timestamp"`
}
