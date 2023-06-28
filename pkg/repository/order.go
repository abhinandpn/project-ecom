package repository

import (
	"time"

	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type OrderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &OrderDatabase{DB: db}
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

func (o *OrderDatabase) FindUserOrderById(uid uint) (domain.UserOrder, error) {

	var body domain.UserOrder
	query := `select * from user_orders where user_id = $1;`
	err := o.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (o *OrderDatabase) AddOrderInfo(uid, aid uint, cpid string, price float64, status string) (uint, error) {

	var body domain.OrderInfo
	CurrentTime := time.Now()

	query := `insert into order_infos (order_id,
					order_time,
					address_id,
					coupon_code,
					total_price,
					order_status)values ($1,$2,$3,$4,$5,$6)returning id;`
	err := o.DB.Raw(query, uid, CurrentTime, aid, cpid, price, status).Scan(&body).Error

	if err != nil {
		return body.Id, err
	}
	return body.Id, nil
}

func (o *OrderDatabase) FindAllOrderInfoByUid(uid uint) (domain.OrderInfo, error) {

	var body domain.OrderInfo
	query := `select * from order_infos where user_id =$1`
	err := o.DB.Raw(query, uid).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (o *OrderDatabase) AddOrderItem(oid, pfid, qty uint) error {

	var body domain.OrderItem
	query := `insert into order_items (user_order_id,product_info_id,quantity)values ($1,$2,$3);`
	err := o.DB.Raw(query, oid, pfid, qty).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}
