package repository

import (
	"fmt"
	"time"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"github.com/abhinandpn/project-ecom/pkg/util/res"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB   *gorm.DB
	user interfaces.UserRepository
}

func NewOrderRepository(db *gorm.DB, UserRepo interfaces.UserRepository) interfaces.OrderRepository {
	return &OrderDatabase{DB: db,
		user: UserRepo,
	}
}

func (o *OrderDatabase) CreateUserOrder(id uint) error {

	var body domain.UserOrder
	query := `insert into user_orders (user_id)values ($1);`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) FindUserOrderByUId(uid uint) (domain.UserOrder, error) {

	var body domain.UserOrder
	query := `select * from user_orders where user_id = $1;`
	err := o.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (o *OrderDatabase) AddOrderInfo(orderId,
	AdrsId,
	CopId uint,
	price float64,
	status uint,
	payId uint) (uint, error) {

	var body domain.OrderInfo
	CurrentTime := time.Now()

	query := `insert into order_infos (order_id,
		                      order_time,
		                      address_id,
		                      coupon_code,
		                      total_price,
		                      order_status,
		                      payment_id)values ($1,$2,$3,$4,$5,$6,$7)returning id;`
	err := o.DB.Raw(query, orderId, CurrentTime, AdrsId, CopId, price, status, payId).Scan(&body).Error

	if err != nil {
		return body.Id, err
	}
	return body.Id, nil
}

func (o *OrderDatabase) FindOrderInfoByUid(uid uint) (domain.OrderInfo, error) {

	var body domain.OrderInfo
	query := `select * from order_infos where order_id = $1`
	err := o.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (o *OrderDatabase) AddOrderItem(oid, pfid, qty uint) error {

	var body domain.OrderItem
	query := `insert into order_items (order_info,product_info_id,quantity)values ($1,$2,$3);`
	err := o.DB.Raw(query, oid, pfid, qty).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) CartAllOrder(orderId, OrderinfoId uint, cart []res.CartDisplay) error {

	for _, item := range cart {
		query := `INSERT INTO order_items (order_id,order_info,
			product_info_id, quantity) VALUES ($1, $2,$3,$4);`

		err := o.DB.Exec(query, orderId, OrderinfoId, item.Id, item.Quantity).Error

		if err != nil {
			return fmt.Errorf("failed to add order item: %v", err)
		}
	}

	// var body domain.OrderItem
	// query := `insert into order_items (order_id,
	// 							order_info,
	// 							product_info_id,
	// 							quantity)values ($1,$2,$3,$4);`
	// err := o.DB.Raw(query, oid, ofid, pfid, qty).Scan(&body).Error
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (o *OrderDatabase) UpdatePaymentMethod(id, pid uint) error {

	var body domain.OrderInfo
	query := `UPDATE order_infos
				SET payment_id = $1
				WHERE id = $2;`
	err := o.DB.Raw(query, pid, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) FindOrderInfoByOrderId(id uint) (domain.OrderInfo, error) {

	var body domain.OrderInfo
	query := `select * from order_infos where order_id = $1`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (o *OrderDatabase) OrderDetail(uid uint) ([]res.ResOrder, error) {

	var body []res.ResOrder
	query := `SELECT
				oi.id,
				oi.address_id,
				oi.total_price,
				oi.order_status,
				oi.payment_id,
				oi2.product_info_id,
				oi2.quantity
			  FROM
				user_orders uo
				JOIN order_infos oi ON uo.id = oi.order_id
				JOIN order_items oi2 ON oi.order_id = oi2.order_id
			  WHERE
				uo.user_id = $1;`

	err := o.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (o *OrderDatabase) ChangeOrderStatus(status string, id uint) error {

	query := `update order_infos set order_status =$1  where order_id =$2 ;`
	err := o.DB.Exec(query, status, id).Error
	if err != nil {
		return err
	}
	return nil
}

// --------Order Status-------
func (or *OrderDatabase) AddOrderStatus(status string) error {

	var body domain.OrderStatus
	query := `insert into order_statuses (status)values ($1);`
	err := or.DB.Raw(query, status).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (or *OrderDatabase) EditOrderStatus(id uint, status string) (domain.OrderStatus, error) {

	var body domain.OrderStatus
	query := `update order_statuses set status = $1 where id = $2 ;`
	err := or.DB.Raw(query, status, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, err
}

func (or *OrderDatabase) DeleteOrderStatusById(id uint) error {

	query := `delete from order_statuses where id = $1;`
	err := or.DB.Exec(query, id).Error
	if err != nil {
		return err
	}
	return err
}

func (or *OrderDatabase) GetOrderStatusById(id uint) (domain.OrderStatus, error) {

	var body domain.OrderStatus
	query := `select * from order_statuses where id = $1;`
	err := or.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (or *OrderDatabase) GetOrderStatusByStatus(status string) (domain.OrderStatus, error) {

	var body domain.OrderStatus
	query := `select * from order_statuses where status =$1; `
	err := or.DB.Raw(query, status).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (pr *OrderDatabase) GetAllOrderStatus() ([]domain.OrderStatus, error) {

	var body []domain.OrderStatus
	query := `select * from order_statuses;`
	err := pr.DB.Raw(query).Find(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

// 01 - 09 - 2023 - Order status updation

func (o *OrderDatabase) ListALlOrderByUid(id uint) ([]domain.OrderInfo, error) {

	var body []domain.OrderInfo
	fmt.Println("-------- >", id)
	query := `SELECT *
				FROM order_infos
				WHERE order_id = $1
				ORDER BY order_time ASC;`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil

}

func (o *OrderDatabase) UpdateOrderStatusToOrdered(id uint) error {

	var body domain.OrderInfo
	query := `update order_infos set order_status = 10 where id = $1;`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) UpdateOrderStatusToDelivered(id uint) error {

	var body domain.OrderInfo
	query := `update order_infos set order_status = 11 where id = $1;`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) UpdateOrderStatusToCancelled(id uint) error {

	var body domain.OrderInfo
	query := `UPDATE order_infos
				SET order_status = 12
				WHERE id = $1;`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) UpdateOrderStatusToReturned(id uint) error {

	var body domain.OrderInfo
	query := `update order_infos set order_status = 12 where id = $1;`
	err := o.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderDatabase) ListOrderDetailByUid(uid uint) ([]res.OrderDetailByUid, error) {

	var body []res.OrderDetailByUid
	query := `SELECT 
				    oi.id AS order_info_id,
				    oi.order_id,
				    oi.order_time,
				    a.id AS address_id,
				    a.house,
				    a.phone_number,
				    a.street,
				    a.city,
				    a.district,
				    a.pincode,
				    a.landmark,
				    a.name AS address_name,
				    oi.coupon_code,
				    ci.code AS coupon_name,
				    ci.discount_price AS coupon_discount,
				    oi.total_price,
				    os.status AS order_status,
				    pd.total_price AS payment_total_price,
				    pm.method AS payment_method,
				    ps.payment_status AS payment_status
				FROM user_orders uo
				JOIN order_infos oi ON uo.id = oi.order_id
				JOIN addresses a ON oi.address_id = a.id
				JOIN order_statuses os ON oi.order_status = os.id
				JOIN payment_details pd ON oi.id = pd.order_id
				JOIN payment_methods pm ON pd.payment_method_id = pm.id
				JOIN payment_statuses ps ON pd.payment_status_id = ps.id
				LEFT JOIN coupons ci ON oi.coupon_code = ci.id
				WHERE uo.user_id = $1;`
	err := o.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}


