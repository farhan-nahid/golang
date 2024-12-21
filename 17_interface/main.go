package main

import "fmt"

// PaymentGateway interface
type PaymentGateway interface {
    Charge(amount float64, currency string, description string) (string, error)
    Refund(transactionID string, amount float64) (string, error)
}

// PayPal implementation of PaymentGateway
type PayPal struct {
    APIKey string
}

func (p PayPal) Charge(amount float64, currency string, description string) (string, error) {
    // Simulate PayPal payment processing
    fmt.Printf("Processing PayPal payment: Amount=%.2f %s, Description=%s\n", amount, currency, description)
    return "paypal_txn_12345", nil
}

func (p PayPal) Refund(transactionID string, amount float64) (string, error) {
    // Simulate PayPal refund
    fmt.Printf("Processing PayPal refund: TransactionID=%s, Amount=%.2f\n", transactionID, amount)
    return "paypal_refund_12345", nil
}

// Stripe implementation of PaymentGateway
type Stripe struct {
    SecretKey string
}

func (s Stripe) Charge(amount float64, currency string, description string) (string, error) {
    // Simulate Stripe payment processing
    fmt.Printf("Processing Stripe payment: Amount=%.2f %s, Description=%s\n", amount, currency, description)
    return "stripe_txn_67890", nil
}

func (s Stripe) Refund(transactionID string, amount float64) (string, error) {
    // Simulate Stripe refund
    fmt.Printf("Processing Stripe refund: TransactionID=%s, Amount=%.2f\n", transactionID, amount)
    return "stripe_refund_67890", nil
}

// ProcessPayment function that uses the PaymentGateway interface
func ProcessPayment(pg PaymentGateway, amount float64, currency string, description string) {
    txnID, err := pg.Charge(amount, currency, description)
    if err != nil {
        fmt.Printf("Payment failed: %v\n", err)
        return
    }
    fmt.Printf("Payment successful: TransactionID=%s\n", txnID)
}

// ProcessRefund function that uses the PaymentGateway interface
func ProcessRefund(pg PaymentGateway, transactionID string, amount float64) {
    refundID, err := pg.Refund(transactionID, amount)
    if err != nil {
        fmt.Printf("Refund failed: %v\n", err)
        return
    }
    fmt.Printf("Refund successful: RefundID=%s\n", refundID)
}

func main() {
    // Initialize PayPal and Stripe instances
    paypal := PayPal{APIKey: "paypal-secret-key"}
    stripe := Stripe{SecretKey: "stripe-secret-key"}

    // Process payments and refunds using PayPal
    fmt.Println("Using PayPal:")
    ProcessPayment(paypal, 100.0, "USD", "Order #123")
    ProcessRefund(paypal, "paypal_txn_12345", 50.0)

    fmt.Println()

    // Process payments and refunds using Stripe
    fmt.Println("Using Stripe:")
    ProcessPayment(stripe, 200.0, "EUR", "Order #456")
    ProcessRefund(stripe, "stripe_txn_67890", 100.0)
}
