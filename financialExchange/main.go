package main

import (
	"financialExchange/enum"
	"financialExchange/modifiers"
	storage "financialExchange/store"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/google/uuid"
)

func main() {
	options := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	logger := slog.New(handler)

	// currencies := storage.NewStorage[string](5)
	// currencies.Add("BTC-RUB")
	// currencies.Add("ETH-RUB")
	// currencies.Add("USDT-RUB")
	// currencies.Add("JPY-RUB")
	// currencies.Add("CNY-RUB")

	// logger.Info("Работа с валютами:",
	// 	slog.Int("количество уникальных комбинаций", currencies.Length()),
	// )

	// pair := "USDT-RUB"
	// logger.Info("Проверка наличия валюты в хранилище:",
	// 	slog.Bool("наличие USDT-RUB", currencies.Has(pair)),
	// )

	// if currencies.Has(pair) {
	// 	currencies.Remove(pair)
	// 	logger.Info("USDT-RUB удалена из хранилища")
	// }

	// logger.Info("Работа с валютами:",
	// 	slog.Int("количество уникальных комбинаций", currencies.Length()),
	// )



	
	// traderStore := storage.NewStore[uuid.UUID, float64](10)
	// traderID := uuid.New()

	// traderStore.Put(traderID, 50000.0004)
	// logger.Info("Новый трейдер добавлен в хранилище:",
	//  slog.String("UUID", traderID.String()),
	// slog.Float64("Начальный баланc", 50000.0004))


	// if balance, ok := traderStore.Get(traderID); ok {
	// 	logger.Info("поиск в хранилище:",
	// 	slog.String("найден трейдер с UUID", traderID.String()),
	// 	slog.Float64("его текущий баланс", balance))
	// }else{	
	// 	logger.Error("элемент не найден в хранилище:",
	// 	slog.String("проверьте UUID", traderID.String()))
	// }

	// orderBook := storage.NewIndexedStorage[uuid.UUID, enum.Order](100)

	orderBook := storage.NewStore[uuid.UUID, enum.Order](10)

//	tradeStore := storage.NewStore[uuid.UUID, enum.Trade](100)

	uuidOrder1 := uuid.New()

	orderBuy := enum.Order{
		ID:       uuidOrder1,
		TraderID: uuid.New(),
		Side:     enum.SideBuy,
		Type:     enum.OrderTypeLimit,
		Symbo:    "BTC-RUB",
		Price:    10000.0,
		Quantity: 2.00,
		Filled:   0.00,
		Status:   enum.OrderStatusNew,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	orderSell1 := enum.Order{
		ID:       uuid.New(),
		TraderID: uuid.New(),
		Side:     enum.SideSell,
		Type:     enum.OrderTypeLimit,
		Symbo:    "BTC-RUB",
		Price:    9800.0,
		Quantity: 2.00,
		Filled:   0.00,
		Status:   enum.OrderStatusNew,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	orderSell2 := enum.Order{
		ID:       uuid.New(),
		TraderID: uuid.New(),
		Side:     enum.SideSell,
		Type:     enum.OrderTypeLimit,
		Symbo:    "BTC-RUB",
		Price:    9900.0,
		Quantity: 0.80,
		Filled:   0.00,
		Status:   enum.OrderStatusNew,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	orderSell3 := enum.Order{
		ID:       uuid.New(),
		TraderID: uuid.New(),
		Side:     enum.SideSell,
		Type:     enum.OrderTypeLimit,
		Symbo:    "BTC-RUB",
		Price:    10000.0,
		Quantity: 0.90,
		Filled:   0.00,
		Status:   enum.OrderStatusNew,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	orderBook.Put(uuidOrder1, orderBuy)

	orderBook.Put(orderBuy.ID, orderBuy)
	orderBook.Put(orderSell1.ID, orderSell1)
	orderBook.Put(orderSell2.ID, orderSell2)
	orderBook.Put(orderSell3.ID, orderSell3)

	updatedOrder, ok := orderBook.Update(uuidOrder1, orderSell1.Quantity, modifiers.ModifierForOrderBuy(orderSell1.Quantity))

	if ok {
		logger.Info("Статус сделки на покупку изменён",
			slog.String("Type", updatedOrder.Type.String()), 
			slog.String("Status: ", updatedOrder.Status.String()), 
			slog. Float64("Value:", updatedOrder.Price), 
			slog. Float64("Quantity", updatedOrder. Quantity), 
			slog.String("Symbo", updatedOrder. Symbo),
			slog.String("TraderID", updatedOrder.TraderID.String()),
			slog.String("Date", updatedOrder.UpdatedAt.String()))
	}			

// 	for _, sellOrder := range orderBook.Ordered() {

// 		if sellOrder.Side != enum.SideSell || !sellOrder.IsActive() {
// 			continue
// 		}

// 		if sellOrder.Symbo != orderBuy.Symbo {
// 			continue
// 		}

// 		if orderBuy.Type == enum.OrderTypeLimit && sellOrder.Price > orderBuy.Price {
// 			continue
// 		}

// 		remainingBuy := orderBuy.Remaining()

// 		if remainingBuy <= 0 {
// 			break
// 		}

// 		remainingSell := sellOrder.Remaining()

// 		tradeQuantity := remainingBuy

// 		if remainingSell < tradeQuantity {
// 			tradeQuantity = remainingSell
// 		}

// 		orderBuy.Filled += tradeQuantity
// 		sellOrder.Filled += tradeQuantity

// 		if orderBuy.Remaining() <= 0 {
// 			orderBuy.Status = enum.OrderStatusFilled
// 		} else {
// 			orderBuy.Status = enum.OrderStatusPartiallyFilled
// 		}

// 		if sellOrder.Remaining() <= 0 {
// 			sellOrder.Status = enum.OrderStatusFilled
// 		} else {
// 			sellOrder.Status = enum.OrderStatusPartiallyFilled
// 		}

// 		orderBuy.UpdatedAt = time.Now().UTC()
// 		sellOrder.UpdatedAt = time.Now().UTC()

// 		orderBook.Put(orderBuy.ID, orderBuy)
// 		orderBook.Put(sellOrder.ID, sellOrder)

// 		trade := enum.Trade{
// 			ID:         uuid.New(),
// 			Symbol:     orderBuy.Symbo,
// 			BuyOrderID: orderBuy.ID,
// 			SellOrderID: sellOrder.ID,
// 			BuyerID:    orderBuy.TraderID,
// 			SellerID:   sellOrder.TraderID,
// 			Price:      sellOrder.Price,
// 			Quantity:   tradeQuantity,
// 			ExecutedAt: time.Now().UTC(),
// 		}

// 		tradeStore.Put(trade.ID, trade)

// 		logger.Info("Совершена сделка:",
// 			slog.String("TradeID", trade.ID.String()),
// 			slog.Float64("Price", trade.Price),
// 			slog.Float64("Quantity", trade.Quantity),
// 		)

// 		logger.Info("Обновление заявки на покупку:",
// 			slog.Float64("Исполнено", orderBuy.Filled),
// 			slog.Float64("Осталось купить", orderBuy.Remaining()),
// 			slog.String("Статус", orderBuy.Status.String()),
// 		)

// 		logger.Info("Обновление заявки на продажу:",
// 			slog.String("UUIDOrder", sellOrder.ID.String()),
// 			slog.Float64("Исполнено", sellOrder.Filled),
// 			slog.Float64("Осталось продать", sellOrder.Remaining()),
// 			slog.String("Статус", sellOrder.Status.String()),
// )
// 	}

// 	logger.Info("Итоговое состояние заявки на покупку:",
// 		slog.String("UUIDOrder", orderBuy.ID.String()),
// 		slog.Float64("Исходное количество", orderBuy.Quantity),
// 		slog.Float64("Исполнено", orderBuy.Filled),
// 		slog.Float64("Осталось купить", orderBuy.Remaining()),
// 		slog.String("Статус", orderBuy.Status.String()),
// 	)

	
	fmt.Scanln()
}