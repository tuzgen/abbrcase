package http

import "fmt"

func findOrderByOrderID(orderId string) struct{} {
	return struct{}{}
}

func createHttpServer() {
	fmt.Println("zort")
}

func createHTTPServer() {
	fmt.Println("zort")
}

func main() {
	orderID := "asd"
	orderId := "Ads"

	createHTTPServer()
	createHttpServer()
	val := findOrderByOrderID(orderID)
	fmt.Println(orderID, orderId, val)

}
