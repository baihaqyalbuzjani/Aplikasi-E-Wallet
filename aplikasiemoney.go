package main

import "fmt"

const MAX = 100

type Akun struct {
	id     string
	pin    int
	saldo  int
	status string
}

type Transaksi struct {
	idAkun     string
	jenis      string
	jumlah     int
	keterangan string
}

func konversiUang(s string) int {
	var hasil int

	for i := 0; i < len(s); i++ {
		if s[i] != '.' {
			hasil = hasil*10 + int(s[i]-'0')
		}
	}

	return hasil
}

func cariAkun(A *[MAX]Akun, n int, id string) int {
	for i := 0; i < n; i++ {
		if A[i].id == id {
			return i
		}
	}
	return -1
}

func registrasi(A *[MAX]Akun, n *int) {
	var id string
	var pin int

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&id)

	fmt.Print("Masukkan PIN : ")
	fmt.Scan(&pin)

	if cariAkun(A, *n, id) != -1 {
		fmt.Println("ID sudah digunakan")
		return
	}

	A[*n].id = id
	A[*n].pin = pin
	A[*n].saldo = 0
	A[*n].status = "pending"

	*n = *n + 1

	fmt.Println("Registrasi berhasil")
	fmt.Println("Menunggu persetujuan admin")
}

func admin(A *[MAX]Akun, n int) {
	var pilih int

	fmt.Println("\n===== DATA PENDING =====")

	for i := 0; i < n; i++ {

		if A[i].status == "pending" {

			fmt.Println("ID :", A[i].id)

			fmt.Print("1.Setujui 2.Tolak : ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				A[i].status = "aktif"
				fmt.Println("Akun diaktifkan")
			} else {
				A[i].status = "ditolak"
				fmt.Println("Akun ditolak")
			}
		}
	}
}

func login(A *[MAX]Akun, n int) int {
	var id string
	var pin int

	fmt.Print("ID : ")
	fmt.Scan(&id)

	fmt.Print("PIN : ")
	fmt.Scan(&pin)

	idx := cariAkun(A, n, id)

	if idx == -1 {
		fmt.Println("ID tidak ditemukan")
		return -1
	}

	if A[idx].status != "aktif" {
		fmt.Println("Akun belum aktif")
		return -1
	}

	if A[idx].pin != pin {
		fmt.Println("PIN salah")
		return -1
	}

	fmt.Println("Login berhasil")
	return idx
}

func topUp(idx int, A *[MAX]Akun, T *[MAX]Transaksi, nt *int) {
	var input string
	var jumlah int

	fmt.Print("Jumlah Top Up : ")
	fmt.Scan(&input)

	jumlah = konversiUang(input)

	if jumlah <= 0 {
		fmt.Println("Jumlah tidak valid")
		return
	}

	A[idx].saldo += jumlah

	T[*nt].idAkun = A[idx].id
	T[*nt].jenis = "TopUp"
	T[*nt].jumlah = jumlah
	T[*nt].keterangan = "-"

	*nt = *nt + 1

	fmt.Println("Top Up berhasil")
}

func transfer(idx int, A *[MAX]Akun, n int, T *[MAX]Transaksi, nt *int) {
	var idTujuan string
	var input string
	var jumlah int

	fmt.Print("ID Tujuan : ")
	fmt.Scan(&idTujuan)

	idxTujuan := cariAkun(A, n, idTujuan)

	if idxTujuan == -1 {
		fmt.Println("Akun tujuan tidak ditemukan")
		return
	}

	if idxTujuan == idx {
		fmt.Println("Tidak bisa transfer ke diri sendiri")
		return
	}

	fmt.Print("Jumlah Transfer : ")
	fmt.Scan(&input)

	jumlah = konversiUang(input)

	if jumlah <= 0 {
		fmt.Println("Jumlah tidak valid")
		return
	}

	if A[idx].saldo < jumlah {
		fmt.Println("Saldo tidak cukup")
		return
	}

	A[idx].saldo -= jumlah
	A[idxTujuan].saldo += jumlah

	T[*nt].idAkun = A[idx].id
	T[*nt].jenis = "Transfer"
	T[*nt].jumlah = jumlah
	T[*nt].keterangan = "Keluar"

	*nt = *nt + 1

	fmt.Println("Transfer berhasil")
}

func pembayaran(idx int, A *[MAX]Akun, T *[MAX]Transaksi, nt *int) {
	var pilih int
	var input string
	var jumlah int
	var layanan string

	fmt.Println("\n===== PEMBAYARAN =====")
	fmt.Println("1. Makanan")
	fmt.Println("2. Pulsa")
	fmt.Println("3. Listrik")
	fmt.Println("4. BPJS")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		layanan = "Makanan"
	} else if pilih == 2 {
		layanan = "Pulsa"
	} else if pilih == 3 {
		layanan = "Listrik"
	} else if pilih == 4 {
		layanan = "BPJS"
	} else {
		fmt.Println("Pilihan tidak valid")
		return
	}

	fmt.Print("Nominal : ")
	fmt.Scan(&input)

	jumlah = konversiUang(input)

	if jumlah <= 0 {
		fmt.Println("Nominal tidak valid")
		return
	}

	if A[idx].saldo < jumlah {
		fmt.Println("Saldo tidak cukup")
		return
	}

	A[idx].saldo -= jumlah

	T[*nt].idAkun = A[idx].id
	T[*nt].jenis = "Pembayaran " + layanan
	T[*nt].jumlah = jumlah
	T[*nt].keterangan = "Keluar"

	*nt = *nt + 1

	fmt.Println("Pembayaran berhasil")
}

func sortTransaksi(T *[MAX]Transaksi, nt int) {
	var temp Transaksi

	for i := 0; i < nt-1; i++ {

		max := i

		for j := i + 1; j < nt; j++ {

			if T[j].jumlah > T[max].jumlah {
				max = j
			}
		}

		temp = T[i]
		T[i] = T[max]
		T[max] = temp
	}
}

func riwayat(idx int, A *[MAX]Akun, T *[MAX]Transaksi, nt int) {

	var pilih int

	fmt.Println("1. Normal")
	fmt.Println("2. Terurut")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 2 {
		sortTransaksi(T, nt)
	}

	fmt.Println("\n===== RIWAYAT =====")

	for i := 0; i < nt; i++ {

		if T[i].idAkun == A[idx].id {

			fmt.Println(
				T[i].jenis,
				T[i].jumlah,
				T[i].keterangan,
			)
		}
	}

	fmt.Println("Saldo :", A[idx].saldo)
}

func insertionSortAkun(A *[MAX]Akun, n int) {
	var temp Akun
	var i, j int

	for i = 1; i < n; i++ {

		temp = A[i]
		j = i - 1

		for j >= 0 && A[j].saldo < temp.saldo {

			A[j+1] = A[j]
			j--

		}

		A[j+1] = temp
	}
}

func cetakAkun(A *[MAX]Akun, n int) {

	insertionSortAkun(A, n)

	fmt.Println("\n===== DAFTAR AKUN =====")

	for i := 0; i < n; i++ {

		fmt.Println(
			"Ranking :", i+1,
			"| ID :", A[i].id,
			"| Saldo :", A[i].saldo,
			"| Status :", A[i].status,
		)
	}
}

func sortID(A *[MAX]Akun, n int) {
	var temp Akun
	var i, j int

	for i = 1; i < n; i++ {

		temp = A[i]
		j = i - 1

		for j >= 0 && A[j].id > temp.id {

			A[j+1] = A[j]
			j--

		}

		A[j+1] = temp
	}
}

func binarySearchAkun(A *[MAX]Akun, n int, cari string) int {

	left := 0
	right := n - 1

	for left <= right {

		mid := (left + right) / 2

		if A[mid].id == cari {

			return mid

		} else if cari < A[mid].id {

			right = mid - 1

		} else {

			left = mid + 1

		}
	}

	return -1
}

func cariAkunBinary(A *[MAX]Akun, n int) {

	var id string
	var idx int

	sortID(A, n)

	fmt.Print("Masukkan ID : ")
	fmt.Scan(&id)

	idx = binarySearchAkun(A, n, id)

	if idx != -1 {

		fmt.Println("\n===== DATA AKUN =====")
		fmt.Println("ID     :", A[idx].id)
		fmt.Println("Saldo  :", A[idx].saldo)
		fmt.Println("Status :", A[idx].status)

	} else {

		fmt.Println("Akun tidak ditemukan")
	}
}

func menuAdmin(A *[MAX]Akun, n int) {

	var pilih int

	for {

		fmt.Println("\n===== MENU ADMIN =====")
		fmt.Println("1. Approve / Tolak Akun")
		fmt.Println("2. Cetak Daftar Akun")
		fmt.Println("3. Cari Akun")
		fmt.Println("4. Kembali")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			admin(A, n)

		} else if pilih == 2 {

			cetakAkun(A, n)

		} else if pilih == 3 {

			cariAkunBinary(A, n)

		} else if pilih == 4 {

			return

		} else {

			fmt.Println("Pilihan tidak valid")
		}
	}
}

func menuUser(idx int, A *[MAX]Akun, n int, T *[MAX]Transaksi, nt *int) {
	var pilih int

	for {

		fmt.Println("1. Transfer")
		fmt.Println("2. Top Up")
		fmt.Println("3. Pembayaran")
		fmt.Println("4. Riwayat")
		fmt.Println("5. Logout")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			transfer(idx, A, n, T, nt)

		} else if pilih == 2 {

			topUp(idx, A, T, nt)

		} else if pilih == 3 {

			pembayaran(idx, A, T, nt)

		} else if pilih == 4 {

			riwayat(idx, A, T, *nt)
			
		} else if pilih == 5 {

			return

		} else {

			fmt.Println("Pilihan tidak valid")
		}
	}
}

func main() {
	var A [MAX]Akun
	var T [MAX]Transaksi

	var n int
	var nt int

	var pilih int

	for {

		fmt.Println("\n===== E-MONEY =====")
		fmt.Println("1. Registrasi")
		fmt.Println("2. Login")
		fmt.Println("3. Admin")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			registrasi(&A, &n)

		} else if pilih == 2 {

			idx := login(&A, n)

			if idx != -1 {
				menuUser(idx, &A, n, &T, &nt)
			}

		} else if pilih == 3 {

			menuAdmin(&A, n)

		} else if pilih == 4 {

			fmt.Println("Program selesai")
			return

		} else {

			fmt.Println("Pilihan tidak valid")
		}
	}
}
