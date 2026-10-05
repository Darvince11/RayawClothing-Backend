package models

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod string
type PaymentStatus string

const (
	PaymentMethodCreditCard  PaymentMethod = "card"
	PaymentMethodMobileMoney PaymentMethod = "mobile_money"
)

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
)

type PaymentHistory struct {
	Id            int           `json:"id"`
	OrderId       uuid.UUID     `json:"order_id"`
	UserId        int           `json:"user_id"`
	Reference     string        `json:"reference"`
	Currency      string        `json:"currency"`
	PaymentMethod PaymentMethod `json:"payment_method"`
	Amount        float64       `json:"amount"`
	PaymentStatus PaymentStatus `json:"payment_status"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type UpdatePaymentHistoryRequest struct {
	Reference     *string
	Currency      *string
	PaymentMethod *PaymentMethod
	PaymentStatus *PaymentStatus
}

type PaymentInit struct {
	Email        string `json:"email"`
	Amount       int    `json:"amount"`
	Reference    string `json:"reference"`
	Callback_Url string `json:"callback_url"`
}

type PaystackInitResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		AuthorizationURL string `json:"authorization_url"`
		AccessCode       string `json:"access_code"`
		Reference        string `json:"reference"`
	} `json:"data"`
}

type PaystackLog struct {
	StartTime int64                `json:"start_time"`
	TimeSpent int                  `json:"time_spent"`
	Attempts  int                  `json:"attempts"`
	Errors    int                  `json:"errors"`
	Success   bool                 `json:"success"`
	Mobile    bool                 `json:"mobile"`
	Input     []any                `json:"input"`
	History   []PaystackLogHistory `json:"history"`
}

type PaystackLogHistory struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int    `json:"time"`
}

type PaystackAuthorization struct {
	AuthorizationCode string  `json:"authorization_code"`
	BIN               string  `json:"bin"`
	Last4             string  `json:"last4"`
	ExpMonth          string  `json:"exp_month"`
	ExpYear           string  `json:"exp_year"`
	Channel           string  `json:"channel"`
	CardType          string  `json:"card_type"`
	Bank              string  `json:"bank"`
	CountryCode       string  `json:"country_code"`
	Brand             string  `json:"brand"`
	Reusable          bool    `json:"reusable"`
	Signature         string  `json:"signature"`
	AccountName       *string `json:"account_name"`
}

type PaystackCustomer struct {
	ID                 int64   `json:"id"`
	FirstName          *string `json:"first_name"`
	LastName           *string `json:"last_name"`
	Email              string  `json:"email"`
	CustomerCode       string  `json:"customer_code"`
	Phone              *string `json:"phone"`
	Metadata           any     `json:"metadata"`
	RiskAction         string  `json:"risk_action"`
	InternationalPhone *string `json:"international_format_phone"`
}

type PaystackTransaction struct {
	ID                 int64                 `json:"id"`
	Domain             string                `json:"domain"`
	Status             string                `json:"status"`
	Reference          string                `json:"reference"`
	ReceiptNumber      *string               `json:"receipt_number"`
	Amount             int64                 `json:"amount"`
	Message            *string               `json:"message"`
	GatewayResponse    string                `json:"gateway_response"`
	PaidAt             time.Time             `json:"paid_at"`
	CreatedAt          time.Time             `json:"created_at"`
	Channel            string                `json:"channel"`
	Currency           string                `json:"currency"`
	IPAddress          string                `json:"ip_address"`
	Metadata           any                   `json:"metadata"`
	Log                PaystackLog           `json:"log"`
	Fees               int64                 `json:"fees"`
	FeesSplit          any                   `json:"fees_split"`
	Authorization      PaystackAuthorization `json:"authorization"`
	Customer           PaystackCustomer      `json:"customer"`
	Plan               any                   `json:"plan"`
	Split              map[string]any        `json:"split"`
	OrderID            any                   `json:"order_id"`
	PaidAtAlt          time.Time             `json:"paidAt"`
	CreatedAtAlt       time.Time             `json:"createdAt"`
	RequestedAmount    int64                 `json:"requested_amount"`
	POSTransactionData any                   `json:"pos_transaction_data"`
	Source             any                   `json:"source"`
	FeesBreakdown      any                   `json:"fees_breakdown"`
	Connect            any                   `json:"connect"`
	TransactionDate    time.Time             `json:"transaction_date"`
	PlanObject         map[string]any        `json:"plan_object"`
	Subaccount         map[string]any        `json:"subaccount"`
}

type PaystackVerifyResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    PaystackTransaction `json:"data"`
}
