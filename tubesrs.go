package main

import "fmt"

const NMAX = 100
const NKAMAR = 9

type Pasien struct {
	id       string
	nama     string
	umur     int
	penyakit string
	status   string
	kamar    string
}

type Kamar struct {
	kode   string
	terisi bool
}

type DataPasien [NMAX]Pasien
type DataKamar [NKAMAR]Kamar

var pasien DataPasien
var kamar DataKamar

// ======================================
// FUNCTION
// ======================================

func generateID(id int) string {
	if id < 10 {
		return fmt.Sprintf("P00%d", id)
	} else if id < 100 {
		return fmt.Sprintf("P0%d", id)
	}

	return fmt.Sprintf("P%d", id)
}

func cariIndexKamar(K DataKamar, kode string) int {
	var i int

	for i = 0; i < NKAMAR; i++ {
		if K[i].kode == kode {
			return i
		}
	}

	return -1
}

func sequentialSearchNama(A DataPasien, n int, nama string) int {
	var i int

	for i = 0; i < n; i++ {
		if A[i].nama == nama {
			return i
		}
	}

	return -1
}

func binarySearchID(A DataPasien, n int, id string) int {
	var left, right, mid int

	left = 0
	right = n - 1

	for left <= right {

		mid = (left + right) / 2

		if A[mid].id == id {
			return mid
		}

		if id < A[mid].id {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// ======================================
// SORTING
// ======================================

func sortID(A *DataPasien, n int) {
	var pass, i int
	var temp Pasien

	for pass = 1; pass < n; pass++ {

		temp = (*A)[pass]
		i = pass - 1

		for i >= 0 && temp.id < (*A)[i].id {

			(*A)[i+1] = (*A)[i]
			i--
		}

		(*A)[i+1] = temp
	}
}

func selectionSortPasienTertua(A *DataPasien, n int) {
	var pass, i, max int
	var temp Pasien

	for pass = 0; pass < n-1; pass++ {

		max = pass

		for i = pass + 1; i < n; i++ {

			if (*A)[i].umur > (*A)[max].umur ||
				((*A)[i].umur == (*A)[max].umur &&
					(*A)[i].id < (*A)[max].id) {

				max = i
			}
		}

		temp = (*A)[pass]
		(*A)[pass] = (*A)[max]
		(*A)[max] = temp
	}
}

func insertionSortPasienTermuda(A *DataPasien, n int) {
	var pass, i int
	var temp Pasien

	for pass = 1; pass < n; pass++ {

		temp = (*A)[pass]
		i = pass - 1

		for i >= 0 &&
			(temp.umur < (*A)[i].umur ||
				(temp.umur == (*A)[i].umur &&
					temp.id < (*A)[i].id)) {

			(*A)[i+1] = (*A)[i]
			i--
		}

		(*A)[i+1] = temp
	}
}

// ======================================
// PROCEDURE OPERASIONAL
// ======================================

func unitKamar(K *DataKamar) {
	K[0] = Kamar{"R01", false}
	K[1] = Kamar{"R02", false}
	K[2] = Kamar{"R03", false}
	K[3] = Kamar{"R04", false}
	K[4] = Kamar{"R05", false}
	K[5] = Kamar{"R06", false}
	K[6] = Kamar{"R07", false}
	K[7] = Kamar{"R08", false}
	K[8] = Kamar{"R09", false}
}

func tampilKamarKosong(K DataKamar) {
	var i int

	fmt.Println("\nKamar Kosong:")

	for i = 0; i < NKAMAR; i++ {
		if !K[i].terisi {
			fmt.Println(K[i].kode)
		}
	}
}

func tambahPasien(A *DataPasien, n *int, K *DataKamar, nextID *int) {
	var nama, penyakit, kodeKamar string
	var umur, cek, idxKamar int

	if *n >= NMAX {

		fmt.Println("Data pasien penuh")

	} else {

		fmt.Print("Nama     : ")
		fmt.Scan(&nama)

		fmt.Print("Umur     : ")
		cek, _ = fmt.Scan(&umur)

		if cek != 1 || umur < 1 || umur > 99 {

			fmt.Println("Input umur tidak valid")
			return
		}

		fmt.Print("Penyakit : ")
		fmt.Scan(&penyakit)

		tampilKamarKosong(*K)

		fmt.Print("Kamar    : ")
		fmt.Scan(&kodeKamar)

		idxKamar = cariIndexKamar(*K, kodeKamar)

		if idxKamar == -1 {

			fmt.Println("Kode kamar tidak ditemukan")

		} else if K[idxKamar].terisi {

			fmt.Println("Kamar sudah terisi")

		} else {

			(*A)[*n].id = generateID(*nextID)
			(*A)[*n].nama = nama
			(*A)[*n].umur = umur
			(*A)[*n].penyakit = penyakit
			(*A)[*n].status = "Dirawat"
			(*A)[*n].kamar = kodeKamar

			K[idxKamar].terisi = true

			fmt.Println("Data berhasil ditambahkan")
			fmt.Println("ID Pasien :", (*A)[*n].id)

			*n = *n + 1
			*nextID = *nextID + 1
		}
	}
}

func tampilPasienRekursif(A DataPasien, n int, idx int) {
	if idx < n {

		fmt.Println("\n====================")
		fmt.Println("ID       :", A[idx].id)
		fmt.Println("Nama     :", A[idx].nama)
		fmt.Println("Umur     :", A[idx].umur)
		fmt.Println("Penyakit :", A[idx].penyakit)
		fmt.Println("Status   :", A[idx].status)
		fmt.Println("Kamar    :", A[idx].kamar)

		tampilPasienRekursif(A, n, idx+1)
	}
}

func cariPasien(A *DataPasien, n int) {
	var pilihan int
	var namaCari string
	var idCari string
	var idx, i int

	if n == 0 {
		fmt.Println("Belum ada data pasien")
		return
	}

	fmt.Println("\n===== CARI PASIEN =====")
	fmt.Println("1. Berdasarkan Nama")
	fmt.Println("2. Berdasarkan ID")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {

		fmt.Print("Nama : ")
		fmt.Scan(&namaCari)

		idx = sequentialSearchNama(*A, n, namaCari)

		if idx == -1 {
			fmt.Println("Data tidak ditemukan")
			return
		}

		for i = 0; i < n; i++ {

			if (*A)[i].nama == namaCari {

				fmt.Println("\nData Ditemukan")
				fmt.Println("ID       :", (*A)[i].id)
				fmt.Println("Nama     :", (*A)[i].nama)
				fmt.Println("Umur     :", (*A)[i].umur)
				fmt.Println("Penyakit :", (*A)[i].penyakit)
				fmt.Println("Status   :", (*A)[i].status)
				fmt.Println("Kamar    :", (*A)[i].kamar)
			}
		}

	} else if pilihan == 2 {

		sortID(A, n)

		fmt.Print("ID Pasien : ")
		fmt.Scan(&idCari)

		idx = binarySearchID(*A, n, idCari)

		if idx == -1 {
			fmt.Println("Data tidak ditemukan")
			return
		}

		fmt.Println("\nData Ditemukan")
		fmt.Println("ID       :", (*A)[idx].id)
		fmt.Println("Nama     :", (*A)[idx].nama)
		fmt.Println("Umur     :", (*A)[idx].umur)
		fmt.Println("Penyakit :", (*A)[idx].penyakit)
		fmt.Println("Status   :", (*A)[idx].status)
		fmt.Println("Kamar    :", (*A)[idx].kamar)

	} else {

		fmt.Println("Pilihan tidak tersedia")
	}
}

func urutkanPasien(A *DataPasien, n int) {
	var pilihan int

	if n == 0 {

		fmt.Println("Belum ada data pasien")

	} else {

		fmt.Println("\n===== URUTKAN PASIEN =====")
		fmt.Println("1. Pasien Tertua")
		fmt.Println("2. Pasien Termuda")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {

			selectionSortPasienTertua(A, n)

			fmt.Println("\nPasien Tertua :",
				(*A)[0].id,
				(*A)[0].nama,
				(*A)[0].umur,
				(*A)[0].penyakit,
				(*A)[0].status,
				(*A)[0].kamar,
			)

		} else if pilihan == 2 {

			insertionSortPasienTermuda(A, n)

			fmt.Println("\nPasien Termuda :",
				(*A)[0].id,
				(*A)[0].nama,
				(*A)[0].umur,
				(*A)[0].penyakit,
				(*A)[0].status,
				(*A)[0].kamar,
			)

		} else {

			fmt.Println("Pilihan tidak tersedia")
			return
		}
	}
	fmt.Println("\nHasil Pengurutan:")

	tampilPasienRekursif(*A, n, 0)
}

// ======================================
// PROCEDURE LANJUTAN
// ======================================

func updateStatusPasien(A *DataPasien, n int, K *DataKamar) {
	var idCari string
	var idx int

	if n == 0 {

		fmt.Println("Belum ada data pasien")

	} else {

		sortID(A, n)

		fmt.Print("ID Pasien : ")
		fmt.Scan(&idCari)

		idx = binarySearchID(*A, n, idCari)

		if idx == -1 {

			fmt.Println("Pasien tidak ditemukan")

		} else {

			var pilihStatus int
			var statusBaru string

			fmt.Println("\n===== UPDATE STATUS PASIEN =====")
			fmt.Println("1. Dirawat")
			fmt.Println("2. Sembuh")
			fmt.Println("3. Meninggal")
			fmt.Print("Pilih Status : ")
			fmt.Scan(&pilihStatus)

			if pilihStatus == 1 {
				statusBaru = "Dirawat"
			} else if pilihStatus == 2 {
				statusBaru = "Sembuh"
			} else if pilihStatus == 3 {
				statusBaru = "Meninggal"
			}

			if (*A)[idx].status != "Dirawat" &&
				statusBaru == "Dirawat" {

				fmt.Println("Status tidak dapat dikembalikan ke Dirawat")

			} else {

				if statusBaru == "Sembuh" ||
					statusBaru == "Meninggal" {

					var idxKamar int

					idxKamar = cariIndexKamar(
						*K,
						(*A)[idx].kamar,
					)

					if idxKamar != -1 {
						K[idxKamar].terisi = false
					}

					(*A)[idx].kamar = "-"
				}

				(*A)[idx].status = statusBaru

				fmt.Println("Status berhasil diperbarui")
			}
		}
	}
}

		

func hapusPasien(A *DataPasien, n *int, K *DataKamar) {
	var idCari string
	var idx, i int
	var konfirmasi int

	if *n == 0 {

		fmt.Println("Belum ada data pasien")

	} else {

		sortID(A, *n)

		fmt.Print("ID Pasien : ")
		fmt.Scan(&idCari)

		idx = binarySearchID(*A, *n, idCari)

		if idx == -1 {

			fmt.Println("Pasien tidak ditemukan")

		} else {

			fmt.Println("1. Ya")
			fmt.Println("2. Tidak")
			fmt.Print("Yakin hapus data? : ")
			fmt.Scan(&konfirmasi)

			if konfirmasi == 1 {

				if (*A)[idx].kamar != "-" {

					var idxKamar int

					idxKamar =
						cariIndexKamar(
							*K,
							(*A)[idx].kamar,
						)

					if idxKamar != -1 {
						K[idxKamar].terisi = false
					}
				}

				for i = idx; i < *n-1; i++ {
					(*A)[i] = (*A)[i+1]
				}

				*n = *n - 1

				fmt.Println("Data berhasil dihapus")

			} else {

				fmt.Println("Penghapusan dibatalkan")
			}
		}
	}
}

func infoKamar(K DataKamar) {
	var i int

	fmt.Println("\n===== INFORMASI KAMAR =====")

	for i = 0; i < NKAMAR; i++ {

		if K[i].terisi {

			fmt.Println(K[i].kode, "- Terisi")

		} else {

			fmt.Println(K[i].kode, "- Kosong")

		}
	}
}

// ======================================
// MENU
// ======================================

func menu() {
	fmt.Println("\n===== SISTEM MANAJEMEN RUMAH SAKIT =====")
	fmt.Println("1. Tambah Pasien")
	fmt.Println("2. Tampilkan Data Pasien")
	fmt.Println("3. Cari Pasien")
	fmt.Println("4. Urutkan Data Pasien")
	fmt.Println("5. Update Status Pasien")
	fmt.Println("6. Hapus Data Pasien")
	fmt.Println("7. Ketersediaan Kamar")
	fmt.Println("8. Keluar")
	fmt.Print("Pilih : ")
}

// ======================================
// MAIN
// ======================================

func main() {
	var pilih int
	var nPasien int
	var nextID int
	var selesai bool

	unitKamar(&kamar)

	nPasien = 0
	nextID = 1
	selesai = false

	for !selesai {

		menu()
		fmt.Scan(&pilih)

		switch pilih {

		case 1:

			tambahPasien(&pasien, &nPasien, &kamar, &nextID)

		case 2:

			if nPasien == 0 {

				fmt.Println("Belum ada data pasien")

			} else {

				tampilPasienRekursif(pasien, nPasien, 0)
			
			}

		case 3:

			cariPasien(&pasien, nPasien)

		case 4:

			urutkanPasien(&pasien, nPasien)

		case 5:

			updateStatusPasien(&pasien, nPasien, &kamar)

		case 6:

			hapusPasien(&pasien, &nPasien, &kamar)

		case 7:

			infoKamar(kamar)

		case 8:

			selesai = true

		default:

			fmt.Println("Menu tidak tersedia")
		}
	}

	fmt.Println("\nTERIMA KASIH")
}