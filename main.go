package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Customer struct {
	ID           int           `json:"ID"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Certificates []Certificate `json:"certificates"`
}

type Certificate struct {
	ID         int    `json:"ID"`
	CustomerID int    `json:"customerid"`
	Key        string `json:"key"`
	Body       string `json:"body"`
}

// Customer handlers
func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Cloudflare API. v0.1")
}

func CreateCustomerHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Println("Customer creation")
	decoder := json.NewDecoder(r.Body)
	var c Customer
	if err := decoder.Decode(&c); err != nil {
		panic(err)
	}
	fmt.Println("Customer", c)
}

func ListCustomersHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func DeleteCustomerHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func ListCustomerCertificate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

// Certificate handlers

func CreateCertificateHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

//factory to build a Customer (return a Json version of the customer)
func NewCustomer(id int, name, email string, certificates []Certificate) ([]byte, error) {
	customer := Customer{id, name, email, certificates}
	c, err := json.Marshal(customer)
	if err != nil {
		fmt.Println("Error Marshalling a customer", err)
		return nil, err
	}
	return c, nil
}

// Factory to build a certificate (return a Json version of the certificate)
func NewCertificate(id, customerID int, key, body string) ([]byte, error) {

	return nil, nil
}

func main() {
	r := httprouter.New()
	r.GET("/", HomeHandler)

	// Customer routes
	r.GET("/customers", ListCustomersHandler)
	r.POST("/customers", CreateCustomerHandler)
	r.DELETE("/customers/:id", DeleteCustomerHandler)
	r.GET("/customers/:id/certificates", ListCustomerCertificate)

	// Certificates routes
	r.POST("/certificates", CreateCertificateHandler)

	// r.GET("/posts/:id", PostShowHandler)
	// r.PUT("/posts/:id", PostUpdateHandler)
	// r.GET("/posts/:id/edit", PostEditHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
