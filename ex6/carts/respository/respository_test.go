package respository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	productRepo "github.com/iris-contrib/clean-arch/article/repository"
	"github.com/iris-contrib/clean-arch/models"
)

func TestAddProductToCart(t *Testing) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "updated_at", "created_at"}).
		AddRow(1, "abc", 100, time.Now(), time.Now())

	query := "SELECT id, name, price, created_at, updated_at FROM product WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	cart := models.Cart{
		ID: 1,
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    1,
			},
		},
		Total: 1,
	}
	payload := models.CartRequestPayload{
		ProductID: 1,
		Quantity:  1,
	}
	CartRepository := CartRepository{
		pr: productRepo.ProductRepository{
			Conn: db,
		},
	}
	err = CartRepository.AddProductToCart(payload, &cart)
	assert.NoError(t, err)
	assert.Equal(t, cart.Total, int64(200))
	assert.Equal(t, cart.Items[1].Quantity, 2)
	assert.Equal(t, cart.Items[1].ProductName, "abc")
}

func TestDeleteProductFromCart(t *Testing) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "updated_at", "created_at"}).
		AddRow(1, "abc", 100, time.Now(), time.Now())

	query := "SELECT id, name, price, created_at, updated_at FROM product WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	cart := models.Cart{
		ID: 1,
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    2,
			},
		},
		Total: 200,
	}
	payload := models.CartRequestPayload{
		ProductID: 1,
		Quantity:  1,
	}
	CartRepository := CartRepository{
		pr: productRepo.ProductRepository{
			Conn: db,
		},
	}
	err = CartRepository.DeleteProductFromCart(payload, &cart)
	assert.NoError(t, err)
	assert.Equal(t, cart.Total, int64(100))
	assert.Equal(t, cart.Items[1].Quantity, 1)
	assert.Equal(t, cart.Items[1].ProductName, "abc")
}

func TestDeleteProductFromCartWithQuantityGreaterThanInCart(t *Testing) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "updated_at", "created_at"}).
		AddRow(1, "abc", 100, time.Now(), time.Now())

	query := "SELECT id, name, price, created_at, updated_at FROM product WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	cart := models.Cart{
		ID: 1,
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    2,
			},
		},
		Total: 200,
	}
	payload := models.CartRequestPayload{
		ProductID: 1,
		Quantity:  3,
	}
	CartRepository := CartRepository{
		pr: productRepo.ProductRepository{
			Conn: db,
		},
	}
	err := CartRepository.DeleteProductFromCart(payload, &cart)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Quantity of product 1 in cart is less than 3")
}

func TestDeleteProductFromCartWithProductNotInCart(t *Testing) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "updated_at", "created_at"}).
		AddRow(1, "abc", 100, time.Now(), time.Now())

	query := "SELECT id, name, price, created_at, updated_at FROM product WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	userID := int64(1)
	prep.ExpectQuery().WithArgs(userID).WillReturnRows(rows)

	cart := models.Cart{
		ID: 1,
		Items: map[uint]models.Item{
			1: models.Item{
				ProductID:   1,
				ProductName: "abc",
				Quantity:    2,
			},
		},
		Total: 200,
	}
	payload := models.CartRequestPayload{
		ProductID: 2,
		Quantity:  3,
	}
	CartRepository := CartRepository{
		pr: productRepo.ProductRepository{
			Conn: db,
		},
	}
	err := CartRepository.DeleteProductFromCart(payload, &cart)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "Product 2 is not in cart")
}
