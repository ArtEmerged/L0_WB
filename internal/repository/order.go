package repository

import (
	"fmt"
	"wblzero/internal/models"

	"github.com/jmoiron/sqlx"
)

type OrderRepo struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Add(order *models.Order) error {
	tx, err := r.db.Beginx()
	if err != nil {
		tx.Rollback()
		return err
	}
	query := fmt.Sprintf(
		"INSERT INTO %s (order_uid, track_number, entry, locale, internal_signature, "+
			"customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		ordersTable,
	)
	_, err = tx.Exec(
		query, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = fmt.Sprintf(
		"INSERT INTO %s (order_uid, name, phone, zip, city, addres, region, email) "+
			"VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		deliveryTable,
	)
	_, err = tx.Exec(
		query, order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.ZIP,
		order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(
		"INSERT INTO %s (order_uid, transaction, request_id, currency, provider, amount, payment_dt,"+
			" bank, delivery_cost, goods_total, custom_fee) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		paymentsTable,
	)
	_, err = tx.Exec(query, order.OrderUID, order.Payment.Transaction, order.Payment.RequestID,
		order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt,
		order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = fmt.Sprintf(
		"INSERT INTO %s (order_uid, chrt_id, track_number, price, rid, name, sale,"+
			" size, total_price, nm_id, brand, status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		itemsTable,
	)
	for _, item := range order.Items {
		_, err = tx.Exec(query, order.OrderUID, item.ChrtID, item.TrackNumber, item.Price, item.Rid,
			item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()

}

func (r *OrderRepo) Get(orderUID string) (*models.Order, error) {
	fmt.Println("i here DB")
	order := new(models.Order)
	query := fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", ordersTable)

	err := r.db.Get(order, query, orderUID)
	if err != nil {
		return nil, err
	}

	delivery := models.Delivery{}
	query = fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", deliveryTable)

	err = r.db.Get(&delivery, query, orderUID)
	if err != nil {
		return nil, err
	}

	payment := models.Payment{}
	query = fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", paymentsTable)

	err = r.db.Get(&payment, query, orderUID)
	if err != nil {
		return nil, err
	}

	items := []models.Item{}
	query = fmt.Sprintf("SELECT * FROM %s WHERE order_uid=$1", itemsTable)
	rows, err := r.db.Queryx(query, orderUID)
	if err != nil {
		return nil, err
	}
	item := models.Item{}
	for rows.Next() {
		err = rows.StructScan(&item)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	order.Delivery = delivery
	order.Payment = payment
	order.Items = items

	return order, nil
}

func (r *OrderRepo) GetCache(sizeCache int) ([]string, error) {
	ordersId := []string{}
	query := fmt.Sprintf("SELECT order_uid FROM %s ORDER BY date_created DESC LIMIT $1", ordersTable)
	rows, err := r.db.DB.Query(query, sizeCache)
	if err != nil {
		return nil, err
	}
	var orderUID string
	for rows.Next() {
		err = rows.Scan(&orderUID)
		if err != nil {
			return nil, err
		}
		ordersId = append(ordersId, orderUID)
	}
	if rows.Err() != nil {
		return nil, err
	}
	return ordersId, nil
}
