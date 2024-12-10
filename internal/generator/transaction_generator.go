package generator

import (
	"fmt"
	"time"
	"github.com/1lker/sd-gen-o2/internal/types"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
)

// Transaction status constants
const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusFailed    = "failed"
	StatusCanceled  = "canceled"
)

// Transaction type constants
const (
	TypePayment     = "payment"
	TypeTransfer    = "transfer"
	TypeRefund      = "refund"
	TypeWithdrawal  = "withdrawal"
	TypeDeposit     = "deposit"
)

// GenerateTransaction generates a single transaction
func (g *Generator) GenerateTransaction(req types.TransactionRequest) types.Transaction {
	// Set default amount range if not specified
	minAmount := req.MinAmount
	maxAmount := req.MaxAmount
	if minAmount == 0 && maxAmount == 0 {
		minAmount = 1.0
		maxAmount = 1000.0
	}

	// Generate amount
	amount := g.faker.Float64Range(minAmount, maxAmount)

	// Set default currency if not specified
	currency := req.Currency
	if currency == "" {
		currencies := []string{"USD", "EUR", "GBP", "JPY"}
		currency = currencies[g.faker.IntRange(0, len(currencies)-1)]
	}

	// Generate timestamp within range or default to recent
	var timestamp time.Time
	if req.MinTimestamp != "" && req.MaxTimestamp != "" {
		minTime, _ := time.Parse(time.RFC3339, req.MinTimestamp)
		maxTime, _ := time.Parse(time.RFC3339, req.MaxTimestamp)
		timestamp = g.faker.DateRange(minTime, maxTime)
	} else {
		// Default to last 30 days
		now := time.Now()
		timestamp = g.faker.DateRange(now.AddDate(0, 0, -30), now)
	}

	// Generate status
	status := req.Status
	if status == "" {
		statuses := []string{StatusPending, StatusCompleted, StatusFailed, StatusCanceled}
		status = statuses[g.faker.IntRange(0, len(statuses)-1)]
	}

	// Generate type
	txType := req.Type
	if txType == "" {
		types := []string{TypePayment, TypeTransfer, TypeRefund, TypeWithdrawal, TypeDeposit}
		txType = types[g.faker.IntRange(0, len(types)-1)]
	}

	// Generate accounts if not specified
	fromAccount := req.FromAccount
	toAccount := req.ToAccount
	if fromAccount == "" {
		fromAccount = fmt.Sprintf("ACC%010d", g.faker.IntRange(1000000000, 9999999999))
	}
	if toAccount == "" {
		toAccount = fmt.Sprintf("ACC%010d", g.faker.IntRange(1000000000, 9999999999))
	}

	// Generate metadata based on transaction type
	metadata := generateMetadata(g.faker, txType)

	// Generate tags based on type and status
	tags := generateTags(txType, status)

	now := time.Now()

	return types.Transaction{
		ID:                 fmt.Sprintf("TXN%s", g.faker.UUID()),
		From:               fromAccount,
		To:                 toAccount,
		Amount:            amount,
		Currency:          currency,
		Timestamp:         timestamp,
		Status:            status,
		Type:              txType,
		Description:       generateDescription(txType, fromAccount, toAccount),
		Metadata:          metadata,
		ExchangeRate:      generateExchangeRate(g.faker, currency),
		ParentTransactionID: generateParentTxID(g.faker, txType),
		ExternalReferenceID: fmt.Sprintf("EXT%s", g.faker.UUID()),
		ErrorReason:        generateErrorReason(status),
		Tags:               tags,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

// Helper functions

func generateMetadata(f *gofakeit.Faker, txType string) map[string]string {
	metadata := make(map[string]string)
	
	switch txType {
	case TypePayment:
		metadata["merchant_id"] = fmt.Sprintf("MER%d", f.IntRange(1000, 9999))
		metadata["payment_method"] = f.RandomString([]string{"card", "bank_transfer", "wallet"})
	case TypeTransfer:
		metadata["bank_code"] = fmt.Sprintf("BANK%d", f.IntRange(100, 999))
		metadata["transfer_type"] = f.RandomString([]string{"domestic", "international", "internal"})
	case TypeRefund:
		metadata["original_transaction_id"] = fmt.Sprintf("TXN%s", f.UUID())
		metadata["refund_reason"] = f.RandomString([]string{"customer_request", "dispute", "error"})
	}

	return metadata
}

func generateTags(txType, status string) []string {
	tags := []string{txType, status}
	
	if status == StatusFailed {
		tags = append(tags, "requires_attention")
	}
	if txType == TypeRefund {
		tags = append(tags, "customer_service")
	}

	return tags
}

func generateDescription(txType, from, to string) string {
	switch txType {
	case TypePayment:
		return fmt.Sprintf("Payment from %s to %s", from, to)
	case TypeTransfer:
		return fmt.Sprintf("Transfer from %s to %s", from, to)
	case TypeRefund:
		return fmt.Sprintf("Refund to %s", to)
	case TypeWithdrawal:
		return fmt.Sprintf("Withdrawal from %s", from)
	case TypeDeposit:
		return fmt.Sprintf("Deposit to %s", to)
	default:
		return fmt.Sprintf("Transaction from %s to %s", from, to)
	}
}

func generateExchangeRate(f *gofakeit.Faker, currency string) float64 {
	if currency == "USD" {
		return 1.0
	}
	return f.Float64Range(0.5, 2.0)
}

func generateParentTxID(f *gofakeit.Faker, txType string) string {
	if txType == TypeRefund {
		return fmt.Sprintf("TXN%s", f.UUID())
	}
	return ""
}

func generateErrorReason(status string) string {
	if status != StatusFailed {
		return ""
	}
	
	reasons := []string{
		"insufficient_funds",
		"invalid_account",
		"network_error",
		"timeout",
		"security_check_failed",
	}
	
	return reasons[rand.Intn(len(reasons))]
}