package modifiers

import(
	"financialExchange/enum"
)
// функция для обновления статуса заявки на покупку
func ModifierForOrderBuy(sellerQuantity float64) func (enum.Order) enum.Order {
	return func(order enum.Order) enum.Order {
		// обновляет количество исполненной заявки
		order.Filled += sellerQuantity

		// проверяет выполнена ли заявка полностью или частично
		if order.Filled >= order.Quantity {
			// если выполнена полностью обновляем статус заявки
			order.Status = enum.OrderStatusFilled
		} else {
			// если выполнена частично обновляем статус заявки
			order.Status = enum.OrderStatusPartiallyFilled
		}
		return order
	}
}