package saw

import (
	"fmt"
	"testing"
)

func TestInsertDafdir(t *testing.T) {
	keterangan := "Masuk"
	kehadiran := "Hadir"
	hasil := InsertDafdir(keterangan, kehadiran)
	fmt.Println(hasil)
}

func TestInsertNilai(t *testing.T) {
	matpel := "Matematika"
	kkm := "70"
	nilai := "90"
	hasil := InsertNilai(matpel, kkm, nilai)
	fmt.Println(hasil)
}

func TestInsertDafpel(t *testing.T) {
	keterangan := "Terlambat"
	poin := "-2"
	hasil := InsertDafpel(keterangan, poin)
	fmt.Println(hasil)
}

func TestInsertPembayaran(t *testing.T) {
	keterangan := "Pembayaran UTS"
	status := "Terbayar"
	hasil := InsertPembayaran(keterangan, status)
	fmt.Println(hasil)
}
