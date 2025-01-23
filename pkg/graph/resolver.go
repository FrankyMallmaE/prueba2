package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/levelstudio/payroll-4ta-crud/pkg/db"
)

type Resolver struct {
	ProductsRepo db.ProductsRepo
}
