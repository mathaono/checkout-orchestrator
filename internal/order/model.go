package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Status representa os possíveis estados de um pedido
type Status string

const (
	StatusCreated   Status = "CREATED"
	StatusPaid      Status = "PAID"
	StatusShipped   Status = "SHIPPED"
	StatusCompleted Status = "COMPLETED"
	StatusFailed    Status = "FAILED"
)

// Order representa um pedido no sistema
type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Items     []Item    `json:"items"`
	Total     float64   `json:"total"`
	Status    Status    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Item representa um item do pedido
type Item struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

func NewOrder(userID string, items []Item) (*Order, error) {
	if userID == "" {
		return nil, fmt.Errorf("ID do usuário não pode ser vazio")
	}

	if len(items) == 0 {
		return nil, fmt.Errorf("o pedido deve conter pelo menos um item")
	}

	// Valida cada item
	for _, item := range items {
		if item.ProductID == "" || item.Quantity <= 0 || item.Price <= 0 {
			return nil, fmt.Errorf("cada item deve ter um ID de produto válido, quantidade maior que zero e preço maior que zero")
		}
	}

	order := &Order{
		ID:        uuid.NewString(),
		UserID:    userID,
		Items:     items,
		Total:     0.0, // Será calculado posteriormente
		Status:    StatusCreated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	order.Total = order.CalculateTotal()

	return order, nil
}

// ValidateCreate valida se um pedido está pronto para ser criado
func (o *Order) ValidateCreate() error {
	if o.ID == "" {
		return fmt.Errorf("ID do pedido não pode ser vazio")
	}
	if o.UserID == "" {
		return fmt.Errorf("ID do usuário não pode ser vazio")
	}
	if len(o.Items) == 0 {
		return fmt.Errorf("o pedido deve conter pelo menos um item")
	}
	for _, item := range o.Items {
		if item.ProductID == "" {
			return fmt.Errorf("ID do produto não pode ser vazio")
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("quantidade do item deve ser maior que zero")
		}
		if item.Price <= 0 {
			return fmt.Errorf("preço do item deve ser maior que zero")
		}
	}
	if o.Total <= 0 {
		return fmt.Errorf("total do pedido deve ser maior que zero")
	}
	if o.Status != StatusCreated {
		return fmt.Errorf("status do pedido deve ser 'CREATED'")
	}
	return nil
}

// CalculateTotal calcula o total do pedido com base nos itens
func (o *Order) CalculateTotal() float64 {
	total := 0.0
	for _, item := range o.Items {
		total += float64(item.Quantity) * item.Price
	}
	o.Total = total
	o.UpdatedAt = time.Now()

	return total
}

// ValidateUpdate valida se um pedido pode ser atualizado
func (o *Order) ValidateUpdate() error {
	if o.ID == "" {
		return fmt.Errorf("ID do pedido não pode ser vazio")
	}
	if o.UserID == "" {
		return fmt.Errorf("ID do usuário não pode ser vazio")
	}
	if len(o.Items) == 0 {
		return fmt.Errorf("o pedido deve conter pelo menos um item")
	}
	for _, item := range o.Items {
		if item.ProductID == "" {
			return fmt.Errorf("ID do produto não pode ser vazio")
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("quantidade do item deve ser maior que zero")
		}
		if item.Price <= 0 {
			return fmt.Errorf("preço do item deve ser maior que zero")
		}
	}
	if o.Total <= 0 {
		return fmt.Errorf("total do pedido deve ser maior que zero")
	}
	if o.Status == StatusCompleted || o.Status == StatusFailed {
		return fmt.Errorf("pedido não pode ser atualizado após estar completo ou falhado")
	}
	return nil
}

// ValidateStatusChange valida se a transição de status do pedido é válida
func (o *Order) ValidateStatusChange(newStatus Status) error {
	if o.ID == "" {
		return fmt.Errorf("ID do pedido não pode ser vazio")
	}
	if o.Status == newStatus {
		return fmt.Errorf("o status do pedido já é '%s'", newStatus)
	}

	switch o.Status {
	case StatusCreated:
		if newStatus != StatusPaid && newStatus != StatusFailed {
			return fmt.Errorf("transição inválida de '%s' para '%s'", o.Status, newStatus)
		}
	case StatusPaid:
		if newStatus != StatusShipped && newStatus != StatusFailed {
			return fmt.Errorf("transição inválida de '%s' para '%s'", o.Status, newStatus)
		}
	case StatusShipped:
		if newStatus != StatusCompleted && newStatus != StatusFailed {
			return fmt.Errorf("transição inválida de '%s' para '%s'", o.Status, newStatus)
		}
	case StatusCompleted, StatusFailed:
		return fmt.Errorf("pedido já está finalizado ou falhou, não pode mudar de status")
	default:
		return fmt.Errorf("status desconhecido: %s", o.Status)
	}

	o.Status = newStatus
	o.UpdatedAt = time.Now()
	return nil
}

// Validate valida se um item está correto
func (i *Item) Validate() error {
	if i.ProductID == "" {
		return fmt.Errorf("ID do produto não pode ser vazio")
	}
	if i.Quantity <= 0 {
		return fmt.Errorf("quantidade do item deve ser maior que zero")
	}
	if i.Price <= 0 {
		return fmt.Errorf("preço do item deve ser maior que zero")
	}
	return nil
}
