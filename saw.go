package gosaw

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDafdir(db *mongo.Database, keterangan string, kehadiran string) (InsertedID interface{}) {
	var dafdir Dafdir
	dafdir.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	dafdir.Keterangan = keterangan
	dafdir.Kehadiran = kehadiran
	return InsertOneDoc(db, "dafdir", dafdir)
}

func InsertNilai(db *mongo.Database, matpel string, kkm string, grade string) (InsertedID interface{}) {
	var nilai Nilai
	nilai.Matpel = matpel
	nilai.Kkm = kkm
	nilai.Grade = grade
	return InsertOneDoc(db, "nilai", nilai)
}

func InsertDafpel(db *mongo.Database, keterangan string, poin string) (InsertedID interface{}) {
	var dafpel Dafpel
	dafpel.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	dafpel.Keterangan = keterangan
	dafpel.Poin = poin
	return InsertOneDoc(db, "dafpel", dafpel)
}

func InsertPembayaran(db *mongo.Database, keterangan string, status string) (InsertedID interface{}) {
	var pembayaran Pembayaran
	pembayaran.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	pembayaran.Keterangan = keterangan
	pembayaran.Status = status
	return InsertOneDoc(db, "pembayaran", pembayaran)
}
