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

func riwayat(idx int, A *[MAX]Akun, T *[MAX]Transaksi, nt int) {
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

func menuUser(idx int, A *[MAX]Akun, n int, T *[MAX]Transaksi, nt *int) {
	var pilih int

	for {

		fmt.Println("\n===== MENU USER =====")
		fmt.Println("1. Transfer")
		fmt.Println("2. Top Up")
		fmt.Println("3. Riwayat")
		fmt.Println("4. Logout")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {

			transfer(idx, A, n, T, nt)

		} else if pilih == 2 {

			topUp(idx, A, T, nt)

		} else if pilih == 3 {

			riwayat(idx, A, T, *nt)

		} else if pilih == 4 {

			fmt.Println("Logout...")
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

			admin(&A, n)

		} else if pilih == 4 {

			fmt.Println("Program selesai")
			return

		} else {

			fmt.Println("Pilihan tidak valid")
		}
	}
}
