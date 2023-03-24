package saw

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDafdir(keterangan string, kehadiran string) (InsertedID interface{}) {
	var dafdir Dafdir
	dafdir.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	dafdir.Keterangan = keterangan
	dafdir.Kehadiran = kehadiran
	return InsertOneDoc("dbmonitor", "kehadiran", dafdir)
}

func InsertNilai(matpel string, kkm string, grade string) (InsertedID interface{}) {
	var nilai Nilai
	nilai.Matpel = matpel
	nilai.Kkm = kkm
	nilai.Grade = grade
	return InsertOneDoc("dbmonitor", "nilai", nilai)
}

func InsertDafpel(keterangan string, poin string) (InsertedID interface{}) {
	var dafpel Dafpel
	dafpel.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	dafpel.Keterangan = keterangan
	dafpel.Poin = poin
	return InsertOneDoc("dbmonitor", "dafpel", dafpel)
}

func InsertPembayaran(keterangan string, status string) (InsertedID interface{}) {
	var pembayaran Pembayaran
	pembayaran.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	pembayaran.Keterangan = keterangan
	pembayaran.Status = status
	return InsertOneDoc("dbmonitor", "pembayaran", pembayaran)
}
