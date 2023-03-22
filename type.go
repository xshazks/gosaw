package saw

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dafdir struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime   primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Keterangan string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
	Kehadiran  string             `bson:"kehadiran,omitempty" json:"kehadiran,omitempty"`
}

type Nilai struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Matpel string             `bson:"matpel,omitempty" json:"matpel,omitempty"`
	Kkm    string             `bson:"kkm,omitempty" json:"kkm,omitempty"`
	Grade  string             `bson:"grade,omitempty" json:"grade,omitempty"`
}

type Dafpel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime   primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Keterangan string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
	Poin       string             `bson:"poin,omitempty" json:"poin,omitempty"`
}

type Pembayaran struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Datetime   primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Keterangan string             `bson:"keterangan,omitempty" json:"keterangan,omitempty"`
	Status     string             `bson:"status,omitempty" json:"status,omitempty"`
}
