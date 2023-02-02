package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Rahul", "Chennai", "400023", "15-02-1993", "1"},
		{"1002", "Ramesh", "Mumbai", "400089", "15-10-2000", "1"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
