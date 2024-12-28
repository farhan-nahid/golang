package main

import "fmt"

// Enum definition using iota
type OrderStatus int

const (
    Pending OrderStatus = iota // 0
    Processing                 // 1
    Shipped                    // 2
    Delivered                  // 3
    Cancelled                  // 4
)

// Convert OrderStatus to string for better readability
func (s OrderStatus) String() string {
    return [...]string{"Pending", "Processing", "Shipped", "Delivered", "Cancelled"}[s]
}

// Example usage of the enum
func main() {
    var status OrderStatus

    status = Pending
    fmt.Printf("Order is currently: %s\n", status)

    status = Shipped
    fmt.Printf("Order is currently: %s\n", status)

    status = Delivered
    fmt.Printf("Order is currently: %s\n", status)

    // Demonstrating invalid value handling
    invalidStatus := OrderStatus(10)
    fmt.Printf("Invalid status: %d\n", invalidStatus)
}
