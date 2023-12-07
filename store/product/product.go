package product
import "fmt"
var(
	List = ProductList{
			{
				Name: "Non",
				Price: 4_000,
				Originalprice: 2_000,
				Quantity: 10,
			},
			{
				Name: "Cola",
				Price: 13_000,
				Originalprice: 6_500,
				Quantity: 0,
			},
			{
				Name: "Nestle",
				Price: 3_000,
				Originalprice: 1_500,
				Quantity: 15,
			},

		}
)
type Product struct{
	Name string
	Price int 
	Originalprice int
	Quantity int 
}
type ProductSellRequest struct{
	Name string
	Quantity int
	//Basket []Basket
}
func NewProduct(name string,price,quantity int) Product{
	return Product{
		Name: name,
		Price: price,
		Quantity: quantity,
	}
}

type ProductList []Product
func (p *ProductList) AddProduct(product Product){
	*p = append(*p, product)
}
func (p *ProductList) RemoveProduct(index int){
	*p = append((*p)[:index],(*p)[index+1:]...)
}
func GetProductInfo() ProductSellRequest{
	var (
		productName string 
		quantity int = 0
	)

	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product quantity: ")
	fmt.Scan(&quantity)

	fmt.Println()

	return ProductSellRequest{
		Name: productName,
		Quantity:  quantity,
	}
}