package app

import (
	"net/http"

	"github.com/banking/service"
	"github.com/gin-gonic/gin"
)

// type Customer struct {
// 	Name    string `json:"name"`
// 	City    string `json:"city"`
// 	Pincode int    `json:"pincode"`
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, world!")
// }

type CustomerHandlers struct {
	Service service.CustomerService
}

func (h *CustomerHandlers) getAllCustomers(c *gin.Context) {

	customers, err := h.Service.GetAllCustomer(c.Query("status"))
	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Message,
		})
		return
	}
	// Mux
	// if r.Header.Get("Content-Type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }

	if c.GetHeader("Content-Type") == "application/xml" {
		c.XML(http.StatusOK, customers)
	} else {
		c.JSON(http.StatusOK, customers)
	}

}

func (h *CustomerHandlers) getCustomerById(c *gin.Context) {

	customer, err := h.Service.GetCustomerById(c.Param("id"))
	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Message,
		})
		return
	}
	c.JSON(http.StatusFound, gin.H{
		"message": "Customer found",
		"data":    customer,
	})
}
