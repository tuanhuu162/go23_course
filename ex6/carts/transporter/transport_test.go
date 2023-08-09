package transporter

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/kataras/iris/v12"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/tuanhuu162/go23_course/ex6/carts/mock"
	"github.com/tuanhuu162/go23_course/ex6/models"
)

func TestCartTransport_AddProductToCart(t *testing.T) {

}

func TestCartTransport_DeleteProductFromCart(t *testing.T) {
}
