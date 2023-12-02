package main
import "fmt"
var(
	productList = ProductList{
			{
				Name: "Non",
				Price: 4_000,
				Quantity: 10,
			},
			{
				Name: "Cola",
				Price: 13_000,
				Quantity: 0,
			},
			{
				Name: "Nestle",
				Price: 3_000,
				Quantity: 15,
			},

		}
)
type Product struct{
	Name string
	Price int 
	Quantity int 
}
type ProductSellRequest struct{
	Name string
	Quantity int
}
func (p Product) NewProduct(name string,price,quantity int) Product{
	return Product{
		Name: name,
		Price: price,
		Quantity: quantity,
	}
}

type ProductList []Product

func getProductInfo() ProductSellRequest{
	var (
		productName string 
		quantity int
	)

	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product quantity: ")
	fmt.Scan(&quantity)

	return ProductSellRequest{
		Name: productName,
		Quantity:  quantity,
	}
}