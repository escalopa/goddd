package mongo

import (
	"context"
	"github.com/escalopa/ddd-go/aggregate"
	"github.com/google/uuid"
	mng "go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository struct {
	db       *mng.Database
	customer *mng.Collection
}

type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewCustomerRepository(db *mng.Database) *CustomerRepository {
	return &CustomerRepository{
		db:       db,
		customer: db.Collection("customer"),
	}
}

func (r *CustomerRepository) Find(ctx context.Context, id string) (*aggregate.Customer, error) {
	customer := &mongoCustomer{}
	err := r.customer.FindOne(ctx, map[string]string{"id": id}).Decode(customer)
	if err != nil {
		return nil, err
	}
	return r.toAggregate(customer)
}

func (r *CustomerRepository) Save(ctx context.Context, customer *aggregate.Customer) error {
	mc, err := r.toMongo(customer)
	if err != nil {
		return err
	}
	_, err = r.customer.InsertOne(ctx, mc)
	return err
}

func (r *CustomerRepository) Update(ctx context.Context, customer *aggregate.Customer) error {
	mc, err := r.toMongo(customer)
	if err != nil {
		return err
	}
	_, err = r.customer.UpdateOne(ctx, map[string]string{"id": customer.GetID()}, mc)
	return err
}

func (r *CustomerRepository) toAggregate(customer *mongoCustomer) (c *aggregate.Customer, err error) {
	err = c.SetID(customer.ID.String())
	if err != nil {
		return nil, err
	}
	err = c.SetName(customer.Name)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CustomerRepository) toMongo(customer *aggregate.Customer) (c *mongoCustomer, err error) {
	id, err := uuid.Parse(customer.GetID())
	if err != nil {
		return nil, err
	}
	return &mongoCustomer{
		ID:   id,
		Name: customer.GetName(),
	}, nil
}
