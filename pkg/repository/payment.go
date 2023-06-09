package repository

import (
	"github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	"gorm.io/gorm"
)

type PaymentDatabase struct {
	DB *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) interfaces.PaymentRepository {
	return &PaymentDatabase{DB: db}
}
func (p *PaymentDatabase) FindPaymentMethodById(id uint) (domain.PaymentMethod, error) {

	var body domain.PaymentMethod
	query := `select * from payment_methods where id =$1`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *PaymentDatabase) AddPaymentMethod(method string, status bool) error {

	var body domain.PaymentMethod
	query := `insert into payment_methods (method,status)values ($1,$2);`
	err := p.DB.Raw(query, method, status).Scan(&body).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PaymentDatabase) ListPaymentMethods() ([]domain.PaymentMethod, error) {

	var body []domain.PaymentMethod
	query := `select * from payment_methods;`
	err := p.DB.Raw(query).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *PaymentDatabase) GetPaymentMethodByName(name string) (domain.PaymentMethod, error) {

	var body domain.PaymentMethod
	query := `select * from payment_methods where method = $1;`
	err := p.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *PaymentDatabase) DeletePaymentMethod(id uint) error {

	query := `delete from payment_methods where id = $1;`
	err := p.DB.Exec(query, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PaymentDatabase) GetPaymentMethodById(id uint) (domain.PaymentMethod, error) {

	var body domain.PaymentMethod
	query := `select * from payment_methods where id = $1;`
	err := p.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}
