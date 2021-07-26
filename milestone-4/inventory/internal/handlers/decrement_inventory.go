package handlers

import (
	log "github.com/sirupsen/logrus"
	"github.com/titanlien/asynchronous-events/milestone-4/models"
)

// DecrementInventory will decrement the inventory of produts by the specific quantity in the order
func DecrementInventory(order models.Order) error {
	log.WithField("order.id", order.ID).
		Info("attempting to decrement inventory from order")

	// We are not actually connecting to an inventory system, so just log it for now
	for _, p := range order.Products {
		log.WithField("order.id", order.ID).
			WithField("product.code", p.ProductCode).
			WithField("product.quantity", p.Quantity).
			Info("decrementing inventory for product")
	}

	return nil
}
