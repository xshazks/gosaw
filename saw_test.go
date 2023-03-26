package saw

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dbmonitor",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertDafdir(t *testing.T) {
	keterangan := "Masuk"
	kehadiran := "Hadir"
	hasil := InsertDafdir(MongoConn, keterangan, kehadiran)
	fmt.Println(hasil)
}

func TestInsertNilai(t *testing.T) {
	matpel := "Matematika"
	kkm := "70"
	nilai := "90"
	hasil := InsertNilai(MongoConn, matpel, kkm, nilai)
	fmt.Println(hasil)
}

func TestInsertDafpel(t *testing.T) {
	keterangan := "Terlambat"
	poin := "-2"
	hasil := InsertDafpel(MongoConn, keterangan, poin)
	fmt.Println(hasil)
}

func TestInsertPembayaran(t *testing.T) {
	keterangan := "Pembayaran UTS"
	status := "Terbayar"
	hasil := InsertPembayaran(MongoConn, keterangan, status)
	fmt.Println(hasil)
}
