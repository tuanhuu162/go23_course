package storage

// import (
// 	"github.com/tuanhuu162/go23_course/ex6/app/models"
// )

// func (c models.Cart) AddItemToCart(product_id string, quantity int64) {
// 	existingItem, err := c.Items[product_id]
// 	if err != nil {
// 		c.Items[product_id] = models.Item{ProductID: product_id, Quantity: quantity}
// 	} else {
// 		existingItem.Quantity += quantity
// 		c.Items[product_id] = existingItem
// 	}
// 	c.Total += quantity
// }

// func (c models.Cart) DeleteItemFromCart(id string) {
// 	delete(c.Items, id)
// }

// func (c models.Cart) Checkout() (models.Cart, int64) {
// 	return c, c.Total
// }
