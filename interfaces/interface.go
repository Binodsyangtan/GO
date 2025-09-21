package main

import "fmt"

//interface
type paymenter interface{
	pay(amount float32)
}

type payment struct {
	// gateway esewa  //problem solve using interface
	gateway paymenter
}

func (p payment) makePayment(amout float32) {
	// rezorpayPaymentGw := rezorpay{}
	// rezorpayPaymentGw.pay(amout)
	// esewaPaymentGw := esewa{}
	// esewaPaymentGw.pay(amout)
	p.gateway.pay(amout)

}

type rezorpay struct {
}

func (r rezorpay) pay(amount float32) {

	//login to make payment
	fmt.Println("making payment using razorpay", amount)

}

type esewa struct{}

//pay(amount float32)same as interface so we donot have to write implement word like other lang 
func (e esewa) pay(amount float32) {
	//logic to make payment
	fmt.Println("making payment using esewa", amount)

}

func main() {
	esewaPaymentGw := esewa{}
	newPayment := payment{
		gateway: esewaPaymentGw,
	}
	newPayment.makePayment(100)

}
