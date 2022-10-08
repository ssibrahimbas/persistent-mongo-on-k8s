package internal

import (
	"context"

	"github.com/ssibrahimbas/ssi-core/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"ssibrahimbas.com/persistent-mongo/car/src/entity"
)

type Repo struct {
	c   *mongo.Collection
	m   *db.MongoDB
	ctx context.Context
}

type RepoConfig struct {
	MongoDB *db.MongoDB
}

func NewRepo(cnf *RepoConfig) *Repo {
	return &Repo{
		c:   cnf.MongoDB.GetCollection("cars"),
		m:   cnf.MongoDB,
		ctx: context.TODO(),
	}
}

func (r *Repo) Create(c *entity.Car) (*entity.Car, error) {
	id, err := r.c.InsertOne(r.ctx, c)
	if err != nil {
		return nil, err
	}
	c.ID = id.InsertedID.(primitive.ObjectID).Hex()
	return c, nil
}

func (r *Repo) GetList() ([]*entity.Car, error) {
	cur, err := r.c.Find(r.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var articles []*entity.Car
	for cur.Next(r.ctx) {
		var a entity.Car
		err := cur.Decode(&a)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &a)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}
