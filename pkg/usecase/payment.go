package usecase

import (
	"errors"

	domain "github.com/abhinandpn/project-ecom/pkg/domain"
	interfaces "github.com/abhinandpn/project-ecom/pkg/repository/interface"
	service "github.com/abhinandpn/project-ecom/pkg/usecase/interfaces"
)

type PaymentUseCase struct {
	paymentRepo interfaces.PaymentRepository
}

func NewPaymentUseCase(PaymentRepo interfaces.PaymentRepository) service.PaymentuseCase {
	return &PaymentUseCase{paymentRepo: PaymentRepo}
}

func (p *PaymentUseCase) PaymentMethods() ([]domain.PaymentMethod, error) {

	body, err := p.paymentRepo.ListPaymentMethods()
	if err != nil {
		return body, err
	}
	return body, nil
}

func (p *PaymentUseCase) AddPaymentMethod(name string, status bool) error {

	method, err := p.paymentRepo.GetPaymentMethodByName(name)
	if err != nil {
		return err
	}
	if method.Id != 0 {
		res := errors.New("payment method alredy exist")
		return res
	} else {
		err := p.paymentRepo.AddPaymentMethod(name, status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *PaymentUseCase) DeletePaymentMethod(id uint) error {

	method, err := p.paymentRepo.GetPaymentMethodById(id)
	if err != nil {
		return err
	}
	if method.Id != 0 {
		err := p.paymentRepo.DeletePaymentMethod(method.Id)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("payment method does not exist")
		return res
	}
	return nil
}

func (p *PaymentUseCase) CreatePaymentStatus(name string) error {

	body, err := p.paymentRepo.GetPaymentStatusByName(name)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		res := errors.New("payment status alredy exist")
		return res
	}
	err = p.paymentRepo.AddPaymentStatus(name)
	if err != nil {
		return err
	}
	return nil
}

func (p *PaymentUseCase) UpdatePaymentStatus(id uint, name string) (domain.PaymentStatus, error) {

	body, err := p.paymentRepo.GetPaymentStatusById(id)
	if err != nil {
		return body, err
	}
	if body.Id != 0 {
		body, err := p.paymentRepo.GetPaymentStatusByName(name)
		if err != nil {
			return body, err
		}
		if body.Id != 0 {
			res := errors.New("status alredy exist")
			return body, res
		}
		body, err = p.paymentRepo.EditPaymentStatus(id, name)
		if err != nil {
			return body, err
		}
	}
	return body, nil
}

func (p *PaymentUseCase) DeltePaymentStatus(id uint) error {

	body, err := p.paymentRepo.GetPaymentStatusById(id)
	if err != nil {
		return err
	}
	if body.Id != 0 {
		err := p.paymentRepo.DeltePaymentStatus(id)
		if err != nil {
			return err
		}
	} else {
		res := errors.New("status does not exist")
		return res
	}
	return nil
}

func (p *PaymentUseCase) FindPaymentStatusById(id uint) (domain.PaymentStatus, error) {

	body, err := p.paymentRepo.GetPaymentStatusById(id)
	if err != nil {
		return body, err
	}
	if body.Id == 0 {
		res := errors.New("payment status does not exist")
		return body, res
	}
	return body, nil
}

func (p *PaymentUseCase) GetAllPaymentStatus() ([]domain.PaymentStatus, error) {

	body, err := p.paymentRepo.GetAllPaymentStatus()
	if err != nil {
		return body, err
	}
	if body == nil {
		res := errors.New("payment status does not exist")
		return body, res
	}
	return body, nil
}
