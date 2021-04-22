package model

import (
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	*UserInput `bson:",inline"`
	ID         primitive.ObjectID `json:"id" xml:"id" bson:"_id,omitempty"`
}

type UserInput struct {
	Name  string `json:"name" xml:"name" bson:"name" validate:"required,alpha"`
	Email string `json:"email" xml:"email" bson:"email" validate:"required,email"`
}

type PagedUser struct {
	Data     []User                         `json:"data" xml:"data"`
	PageInfo mongopagination.PaginationData `json:"pageInfo" xml:"pageInfo"`
}
