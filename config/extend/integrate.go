package config

//IntegrateConfig : config serve receive data send from this
var IntegrateConfig integrateConfig

type integrateConfig struct {
	Transaction string
	Customer    string
}

func (itg *integrateConfig) check() {
	if itg.Transaction == "" {
		itg.Transaction = "http://localhost:3456/integrate/transaction"
	}

	if itg.Customer == "" {
		itg.Customer = "http://localhost:3456/integrate/customer"
	}
}
