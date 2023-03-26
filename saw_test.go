package gosaw

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dbmonitor",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertDafdir(t *testing.T) {
	keterangan := "Sakit"
	kehadiran := "Tidak Hadir"
	hasil := InsertDafdir(MongoConn, keterangan, kehadiran)
	fmt.Println(hasil)
}

func TestInsertNilai(t *testing.T) {
	matpel := "Biologi"
	kkm := "70"
	nilai := "80"
	hasil := InsertNilai(MongoConn, matpel, kkm, nilai)
	fmt.Println(hasil)
}

func TestInsertDafpel(t *testing.T) {
	keterangan := "Atribut Tidak Lengkap"
	poin := "-1"
	hasil := InsertDafpel(MongoConn, keterangan, poin)
	fmt.Println(hasil)
}

func TestInsertPembayaran(t *testing.T) {
	keterangan := "Pembayaran SPP"
	status := "Belum bayar"
	hasil := InsertPembayaran(MongoConn, keterangan, status)
	fmt.Println(hasil)
}
