package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"microtwo/models"
	"microtwo/utils"
)

type Customers []models.Customer

type CustomerRepository interface {
	GetCustomersAll(ctx context.Context) (*Customers, error)
	GetCustomerById(ctx context.Context, id string) (*models.Customer, error)
	SaveCustomer(ctx context.Context, customer models.Customer) error
	UpdateCustomer(ctx context.Context, customer models.Customer) error
	DeleteCustomer(ctx context.Context, ID uuid.UUID) error
	SearchCustomer(ctx context.Context, specification CustomerSpecification) (*Customers, error)
}

type CustomerSpecification interface{}

type CustomerRepositoryImpl struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) (*CustomerRepositoryImpl, error) {
	return &CustomerRepositoryImpl{db}, nil
}

func (r *CustomerRepositoryImpl) GetCustomersAll(ctx context.Context) (*Customers, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, fullname, email, phone FROM customers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var customers Customers

	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.Fullname, &customer.Email, &customer.Phone)
		if err != nil {
			return nil, err
		}

		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &customers, nil

}

func (r *CustomerRepositoryImpl) GetCustomerById(ctx context.Context, id string) (*models.Customer, error) {
	var customer = models.Customer{}

	query := r.db.QueryRowContext(ctx, "SELECT id, fullname, email, phone FROM customers WHERE id=?", id)
	err := query.Scan(&customer.ID, &customer.Fullname, &customer.Email, &customer.Phone)

	if err == sql.ErrNoRows {
		return nil, utils.ErrCustomerNotFound
	}

	if err != nil {
		return nil, err
	}

	return &customer, nil

}

func (r *CustomerRepositoryImpl) SaveCustomer(ctx context.Context, customer models.Customer) error {
	result, err := r.db.ExecContext(ctx, "INSERT INTO customers (id, fullname, email, phone) VALUES ($1, $2, $3, $4)", customer.ID, customer.Fullname, customer.Email, customer.Phone)

	if err != nil {
		return err
	}

	// lastId, err := result.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return utils.ErrCustomerNotSaved
	}

	return nil
}

func (r *CustomerRepositoryImpl) UpdateCustomer(ctx context.Context, customer models.Customer) error {
	result, err := r.db.ExecContext(ctx, "UPDATE customers SET fullname = $1, email = $2, phone = $3 WHERE id = $4", customer.Fullname, customer.Email, customer.Phone, customer.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return utils.ErrCustomerNotFound
	}

	return nil
}

func (r *CustomerRepositoryImpl) DeleteCustomer(ctx context.Context, ID uuid.UUID) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM customers WHERE id = $1", ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepositoryImpl) SearchCustomer(ctx context.Context, specification CustomerSpecification) (*Customers, error) {
	// query = "SELECT id, fullname, email, phone FROM customers"

	// for
	// rows, err := r.db.QueryContext(ctx, "SELECT id, fullname, email, phone FROM customers WHERE ")
	return nil, nil
}
