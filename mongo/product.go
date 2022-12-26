package mongo

import (
	"context"
	"github.com/escalopa/ddd-go/aggregate"
	"github.com/google/uuid"
	mng "go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	db      *mng.Database
	product *mng.Collection
}

type mongoProduct struct {
	ID          uuid.UUID `bson:"id"`
	Name        string    `bson:"name"`
	Description string    `bson:"description"`
	Price       float64   `bson:"price"`
	Quantity    int       `bson:"quantity"`
}

func NewProductRepository(db *mng.Database) *ProductRepository {
	return &ProductRepository{
		db:      db,
		product: db.Collection("product"),
	}
}

func (r *ProductRepository) Find(ctx context.Context, id string) (*aggregate.Product, error) {
	product := &mongoProduct{}
	err := r.product.FindOne(ctx, map[string]string{"id": id}).Decode(product)
	if err != nil {
		return nil, err
	}
	return r.toAggregate(product)
}

func (r *ProductRepository) Save(ctx context.Context, product *aggregate.Product) error {
	mp, err := r.toMongo(product)
	if err != nil {
		return err
	}
	_, err = r.product.InsertOne(ctx, mp)
	return err
}

func (r *ProductRepository) Update(ctx context.Context, product *aggregate.Product) error {
	mp, err := r.toMongo(product)
	if err != nil {
		return err
	}
	_, err = r.product.UpdateOne(ctx, map[string]string{"id": product.GetID()}, mp)
	return err
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]*aggregate.Product, error) {
	cursor, err := r.product.Find(ctx, map[string]string{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var products []*aggregate.Product
	for cursor.Next(ctx) {
		product := &mongoProduct{}
		err := cursor.Decode(product)
		if err != nil {
			return nil, err
		}
		aggregateProd, err := r.toAggregate(product)
		if err != nil {
			return nil, err
		}
		products = append(products, aggregateProd)
	}
	return products, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	_, err := r.product.DeleteOne(ctx, map[string]string{"id": id})
	return err
}

func (r *ProductRepository) toAggregate(product *mongoProduct) (p *aggregate.Product, err error) {
	err = p.SetID(product.ID.String())
	if err != nil {
		return nil, err
	}
	err = p.SetName(product.Name)
	if err != nil {
		return nil, err
	}
	err = p.SetDescription(product.Description)
	if err != nil {
		return nil, err
	}
	err = p.SetPrice(product.Price)
	if err != nil {
		return nil, err
	}
	err = p.SetQuantity(product.Quantity)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) toMongo(product *aggregate.Product) (mp *mongoProduct, err error) {
	mp = &mongoProduct{}
	err = mp.ID.UnmarshalText([]byte(product.GetID()))
	if err != nil {
		return nil, err
	}
	mp.Name = product.GetItem().Name
	mp.Description = product.GetItem().Description
	mp.Price = product.GetPrice()
	mp.Quantity = product.GetQuantity()
	return mp, nil
}
