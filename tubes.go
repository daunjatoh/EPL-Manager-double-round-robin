package main

import (
	"fmt"
)

const NMAX = 100000

type dataklub struct {
	nama                                          string
	nPertandingan, nKemenangan, nSeri, nKekalahan int
	nGolMasuk, nGolKemasukan, nSelisihGol, nPoint int
}
type tabInt [NMAX]dataklub

func inputNamaClub(A *tabInt, n int) {

	var h, i, indeks int

	if n > NMAX {
		n = NMAX
	}
	fmt.Println("Masukkan", n, "nama club yang bertanding pada musim ini")
	for i = 0; i < n; i++ {
		fmt.Print(i+1, ". ")
		fmt.Scan(&A[i].nama)
		for len(A[i].nama) != 3 {
			fmt.Println("----------------------------------------------------------------------")
			fmt.Println("!!! nama club tidak valid !!!")
			fmt.Println("Silahkan masukkan kembali nama club yang valid (3 karakter): ")
			fmt.Println("----------------------------------------------------------------------")
			fmt.Println()
			fmt.Println("Masukkan", n, "nama club yang bertanding pada musim ini")
			fmt.Print(i+1, ". ")
			fmt.Scan(&A[i].nama)
		}
	}
	fmt.Println()
	fmt.Println("Berikut adalah nama club yang bertanding pada musim ini: ")

	for h = 0; h < n; h++ {

		fmt.Print(h+1, ". ")
		fmt.Println(A[h].nama)

	}
	fmt.Println()
	inputTiappPekan(A, n, indeks)
}

func inputTiappPekan(A *tabInt, n, indeks int) {
	var h, j, k, golmasuk, golkemasukan int
	var namaClubDicari, namaClubDicari2 string

	var tempindeks1, tempindeks2 int

	for j = 0; j < ((n - 1) * 2); j++ {
		fmt.Println("Pekan ke", j+1, ": ")

		for k = 0; k < n; k++ {
			for {
				fmt.Println("Masukkan nama club yang bertanding pada pertandingan ke", k+1, ": ")
				fmt.Scan(&namaClubDicari)
				tempindeks1 = findData(*A, n, namaClubDicari)

				fmt.Println("Masukkan nama club yang menjadi lawan dari club", namaClubDicari, "pada pertandingan ke", k+1, ": ")
				fmt.Scan(&namaClubDicari2)
				tempindeks2 = findData2(*A, n, namaClubDicari2)

				if tempindeks1 > -1 && tempindeks2 > -1 {
					break
				}

				if tempindeks1 == -1 || tempindeks2 == -1 {
					fmt.Println("Nama club yang anda inputkan tidak ditemukan di dalam daftar nama club yang bermain musim ini")
					fmt.Println("Silahkan input nama club yang terdapat dalam daftar nama club yang bermain musim ini")
					fmt.Println()
					fmt.Println("Berikut adalah nama club yang bertanding pada musim ini: ")
					for h = 0; h < n; h++ {

						fmt.Print(h+1, ". ")
						fmt.Println(A[h].nama)

					}
				}
			}
			if tempindeks1 > -1 && tempindeks2 > -1 {
				A[tempindeks1].nPertandingan = A[tempindeks1].nPertandingan + 1

				A[tempindeks2].nPertandingan = A[tempindeks2].nPertandingan + 1
				fmt.Println()
				fmt.Println("Silahkan masukkan data score hasil pertandingan club pekan ke", j+1, "pertandingan ke", k+1, ": ")
				fmt.Println()
				fmt.Println(A[tempindeks1].nama, "VS", A[tempindeks2].nama)
				fmt.Println()

				fmt.Println("Masukkan jumlah gol masuk: ")
				fmt.Println("Data jumlah gol berikut adalah yang dialami oleh club", A[tempindeks1].nama)
				fmt.Scan(&golmasuk)

				fmt.Println()
				fmt.Println("Masukkan jumlah gol kemasukan: ")
				fmt.Println("Data jumlah gol berikut adalah yang dialami oleh club", A[tempindeks1].nama)
				fmt.Scan(&golkemasukan)

				fmt.Println()

				if golmasuk > golkemasukan {
					A[tempindeks1].nKemenangan = A[tempindeks1].nKemenangan + 1
					A[tempindeks2].nKekalahan = A[tempindeks2].nKekalahan + 1

				} else if golmasuk < golkemasukan {
					A[tempindeks1].nKekalahan = A[tempindeks1].nKekalahan + 1
					A[tempindeks2].nKemenangan = A[tempindeks2].nKemenangan + 1

				} else if golmasuk == golkemasukan {
					A[tempindeks1].nSeri = A[tempindeks1].nSeri + 1
					A[tempindeks2].nSeri = A[tempindeks2].nSeri + 1
				}

				A[tempindeks1].nGolMasuk = A[tempindeks1].nGolMasuk + golmasuk
				A[tempindeks1].nGolKemasukan = A[tempindeks1].nGolKemasukan + golkemasukan

				A[tempindeks2].nGolMasuk = A[tempindeks2].nGolMasuk + golkemasukan
				A[tempindeks2].nGolKemasukan = A[tempindeks2].nGolKemasukan + golmasuk

				/*mencari selisih gol*/

				A[tempindeks1].nSelisihGol = A[tempindeks1].nGolMasuk - A[tempindeks1].nGolKemasukan
				if A[tempindeks1].nSelisihGol < 0 {
					A[tempindeks1].nSelisihGol = A[tempindeks1].nSelisihGol * -1
				}

				/*mencari npoint*/

				A[tempindeks1].nPoint = (A[tempindeks1].nKemenangan * 3) + (A[tempindeks1].nSeri * 1) + (A[tempindeks1].nKekalahan * 0)

				/*mencari selisih gol*/

				A[tempindeks2].nSelisihGol = A[tempindeks2].nGolMasuk - A[tempindeks2].nGolKemasukan
				if A[tempindeks2].nSelisihGol < 0 {
					A[tempindeks2].nSelisihGol = A[tempindeks2].nSelisihGol * -1
				}

				/*mencari npoint*/

				A[tempindeks2].nPoint = (A[tempindeks2].nKemenangan * 3) + (A[tempindeks2].nSeri * 1) + (A[tempindeks2].nKekalahan * 0)
			}

			fmt.Println()
		}
	}
}

func cetakData(A *tabInt, n *int) {
	var i int
	fmt.Println("Berikut adalah data yang kami rekam sebanyak", *n, "club", ": ")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()
	fmt.Printf("%7s %7s %7s %7s %7s %7s %7s %7s %7s %7s\n", "|  No. |", " Nama Club |", "Jumlah Pertandingan |", "Jumlah Menang |", "Jumlah Seri |", "Jumlah Kalah |", "Jumlah Gol Masuk |", "Jumlah Gol Kemasukan |", "Selisih Gol |", "Jumlah Poin |")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------")
	for i = 0; i < *n; i++ {
		fmt.Printf("| %3d  | %6s     | %10d          | %7d       | %7d     | %7d      | %9d        | %10d           | %7d     | %7d     |\n", (i + 1), A[i].nama, A[i].nPertandingan, A[i].nKemenangan, A[i].nSeri, A[i].nKekalahan, A[i].nGolMasuk, A[i].nGolKemasukan, A[i].nSelisihGol, A[i].nPoint)
	}
}

func findData(A tabInt, n int, namaClubDicari string) int {
	var temp, i int
	temp = -1
	i = 0
	for temp == -1 && i < n {
		if A[i].nama == namaClubDicari {
			temp = i

		}
		i++

	}

	return temp
}

func findData2(A tabInt, n int, namaClubDicari2 string) int {
	var temp, i int
	temp = -1
	i = 0
	for temp == -1 && i < n {
		if A[i].nama == namaClubDicari2 {
			temp = i

		}
		i++

	}

	return temp
}

func hapusData(A *tabInt, n *int, indeks int) {
	var namaClubDicari string
	var i int
	fmt.Println()
	fmt.Println("Silahkan masukkan nama club yang ingin dihapus ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Silahkan masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		hapusData(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Data club", namaClubDicari, "berhasil dihapus")
			fmt.Println()
			for i = indeks; i < *n-1; i++ {
				A[i] = A[i+1]
			}
			*n = *n - 1
		}
		if indeks == -1 {
			fmt.Println()
			fmt.Println("Data tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			hapusData(A, n, indeks)
		}
	}
}

func tambahData(A *tabInt, n *int) {

	var nDataBaru, nilaiNTerakhir int
	var h, i, j, k, golmasuk, golkemasukan int
	var namaClubDicari, namaClubDicari2 string

	var tempindeks1, tempindeks2, penentupekan int
	nilaiNTerakhir = *n
	penentupekan = nilaiNTerakhir

	fmt.Println("Berapa banyak data club yang ingin anda tambahkan?")
	fmt.Println("0. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&nDataBaru)
	if nDataBaru == 0 {
		partOfMenuEdit(A, n)
	} else {
		*n = *n + nDataBaru

		if *n > NMAX {
			*n = NMAX
		}
		fmt.Println("Masukkan", nDataBaru, "nama club yang bertanding pada musim ini")
		for i = nilaiNTerakhir; i < *n; i++ {
			fmt.Print(i+1, ". ")
			fmt.Scan(&A[i].nama)
			for len(A[i].nama) != 3 {
				fmt.Println("----------------------------------------------------------------------")
				fmt.Println("!!! nama club tidak valid !!!")
				fmt.Println("Silahkan masukkan kembali nama club yang valid (3 karakter): ")
				fmt.Println("----------------------------------------------------------------------")
				fmt.Println()
				fmt.Println("Masukkan nama club yang bertanding pada musim ini")
				fmt.Scan(&A[i].nama)
			}
		}
		fmt.Println()
		fmt.Println("Berikut adalah nama club yang bertanding pada musim ini: ")

		for h = 0; h < *n; h++ {

			fmt.Print(h+1, ". ")
			fmt.Println(A[h].nama)

		}
		fmt.Println()

		for j = 0; j < ((nDataBaru - 1) * 2); j++ {

			penentupekan = penentupekan + 1
			fmt.Println("Pekan ke", penentupekan, ": ")

			for k = nilaiNTerakhir; k < *n; k++ {
				for {
					fmt.Println("Masukkan nama club yang bertanding pada pertandingan ke", ((k - nilaiNTerakhir) + 1), ": ")
					fmt.Scan(&namaClubDicari)
					tempindeks1 = findData(*A, *n, namaClubDicari)

					fmt.Println("Masukkan nama club yang menjadi lawan dari club", namaClubDicari, "pada pertandingan ke", ((k - nilaiNTerakhir) + 1), ": ")
					fmt.Scan(&namaClubDicari2)
					tempindeks2 = findData2(*A, *n, namaClubDicari2)

					if tempindeks1 > -1 && tempindeks2 > -1 {
						break
					}

					if tempindeks1 == -1 || tempindeks2 == -1 {
						fmt.Println("Nama club yang anda inputkan tidak ditemukan di dalam daftar nama club yang bermain musim ini")
						fmt.Println("Silahkan input nama club yang terdapat dalam daftar nama club yang bermain musim ini")
						fmt.Println()
						fmt.Println("Berikut adalah nama club yang bertanding pada musim ini: ")
						for h = 0; h < *n; h++ {
							fmt.Println(h+1, A[h].nama)
						}
					}
				}
				if tempindeks1 > -1 && tempindeks2 > -1 {
					A[tempindeks1].nPertandingan = A[tempindeks1].nPertandingan + 1

					A[tempindeks2].nPertandingan = A[tempindeks2].nPertandingan + 1
					fmt.Println()
					fmt.Println("Silahkan masukkan data score hasil pertandingan club pekan ke", penentupekan, "pertandingan ke", ((k - nilaiNTerakhir) + 1), ": ")
					fmt.Println()
					fmt.Println(A[tempindeks1].nama, "VS", A[tempindeks2].nama)
					fmt.Println()

					fmt.Println("Masukkan jumlah gol masuk: ")
					fmt.Println("Data jumlah gol berikut adalah yang dialami oleh club", A[tempindeks1].nama)
					fmt.Scan(&golmasuk)

					fmt.Println()
					fmt.Println("Masukkan jumlah gol kemasukan: ")
					fmt.Println("Data jumlah gol berikut adalah yang dialami oleh club", A[tempindeks1].nama)
					fmt.Scan(&golkemasukan)

					fmt.Println()

					if golmasuk > golkemasukan {
						A[tempindeks1].nKemenangan = A[tempindeks1].nKemenangan + 1
						A[tempindeks2].nKekalahan = A[tempindeks2].nKekalahan + 1

					} else if golmasuk < golkemasukan {
						A[tempindeks1].nKekalahan = A[tempindeks1].nKekalahan + 1
						A[tempindeks2].nKemenangan = A[tempindeks2].nKemenangan + 1

					} else if golmasuk == golkemasukan {
						A[tempindeks1].nSeri = A[tempindeks1].nSeri + 1
						A[tempindeks2].nSeri = A[tempindeks2].nSeri + 1
					}

					A[tempindeks1].nGolMasuk = A[tempindeks1].nGolMasuk + golmasuk
					A[tempindeks1].nGolKemasukan = A[tempindeks1].nGolKemasukan + golkemasukan

					A[tempindeks2].nGolMasuk = A[tempindeks2].nGolMasuk + golkemasukan
					A[tempindeks2].nGolKemasukan = A[tempindeks2].nGolKemasukan + golmasuk

					/*mencari selisih gol*/

					A[tempindeks1].nSelisihGol = A[tempindeks1].nGolMasuk - A[tempindeks1].nGolKemasukan
					if A[tempindeks1].nSelisihGol < 0 {
						A[tempindeks1].nSelisihGol = A[tempindeks1].nSelisihGol * -1
					}

					/*mencari npoint*/

					A[tempindeks1].nPoint = (A[tempindeks1].nKemenangan * 3) + (A[tempindeks1].nSeri * 1) + (A[tempindeks1].nKekalahan * 0)

					/*mencari selisih gol*/

					A[tempindeks2].nSelisihGol = A[tempindeks2].nGolMasuk - A[tempindeks2].nGolKemasukan
					if A[tempindeks2].nSelisihGol < 0 {
						A[tempindeks2].nSelisihGol = A[tempindeks2].nSelisihGol * -1
					}

					/*mencari npoint*/

					A[tempindeks2].nPoint = (A[tempindeks2].nKemenangan * 3) + (A[tempindeks2].nSeri * 1) + (A[tempindeks2].nKekalahan * 0)
				}

				fmt.Println()
			}
		}

		fmt.Println("Berhasil menambahkan data sebanyak", nDataBaru, "club")
		fmt.Println()
	}
}

func editSemuaDataPertandingan(A *tabInt, n *int, indeks int) {
	var namaClubDicari string
	fmt.Println("Silahkan masukkan nama club yang ingin dilakukan perubahan: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editSemuaDataPertandingan(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan masukan kembali dan perbarui data untuk club", A[indeks].nama, ": ")
			fmt.Println("___________________________________________________")
			fmt.Println("Masukkan jumlah pertandingannya: ")
			fmt.Scan(&A[indeks].nPertandingan)
			fmt.Println()
			fmt.Println("jumlah kemenangan pertandingan club: ")
			fmt.Scan(&A[indeks].nKemenangan)
			fmt.Println()
			fmt.Println("Masukkan total seri pertandingan club: ")
			fmt.Scan(&A[indeks].nSeri)
			fmt.Println()
			fmt.Println("Masukkan kekalahan pertandingan club: ")
			fmt.Scan(&A[indeks].nKekalahan)
			fmt.Println()
			fmt.Println("Masukkan jumlah gol masuk: ")
			fmt.Scan(&A[indeks].nGolMasuk)
			fmt.Println()
			fmt.Println("Masukkan jumlah gol kemasukan: ")
			fmt.Scan(&A[indeks].nGolKemasukan)
			fmt.Println()
			fmt.Println("Semua data pertandingan club", A[indeks].nama, "berhasil diperbaharui")
			/*mencari selisih gol*/

			A[indeks].nSelisihGol = A[indeks].nGolMasuk - A[indeks].nGolKemasukan
			if A[indeks].nSelisihGol < 0 {
				A[indeks].nSelisihGol = A[indeks].nSelisihGol * -1
			}

			/*mencari npoint*/

			A[indeks].nPoint = (A[indeks].nKemenangan * 3) + (A[indeks].nSeri * 1) + (A[indeks].nKekalahan * 0)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editSemuaDataPertandingan(A, n, indeks)
		}
	}
}

func editJumlahPertandingan(A *tabInt, n *int, indeks int) {
	var namaClubDicari string
	fmt.Println("Silahkan masukkan nama club yang ingin dirubah jumlah pertandingannya: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editJumlahPertandingan(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan perbarui jumlah pertandingan club", (A[indeks].nama))
			fmt.Scan(&A[indeks].nPertandingan)
			fmt.Println()
			fmt.Println("Jumlah pertandingan club", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nPertandingan)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editJumlahPertandingan(A, n, indeks)

		}
	}
}

func editTotalKemenanganPertandingan(A *tabInt, n *int, indeks int) {
	var namaClubDicari string

	fmt.Println("Silahkan masukkan nama club yang ingin dirubah total kemenangannya: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editTotalKemenanganPertandingan(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan perbarui total kemenangan club", (A[indeks].nama))
			fmt.Scan(&A[indeks].nKemenangan)
			fmt.Println()
			fmt.Println("Total kemenangan club", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nKemenangan)

			/*mencari npoint*/
			A[indeks].nPoint = (A[indeks].nKemenangan * 3) + (A[indeks].nSeri * 1) + (A[indeks].nKekalahan * 0)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editTotalSeriPertandingan(A, n, indeks)
		}
	}
}

func editTotalSeriPertandingan(A *tabInt, n *int, indeks int) {
	var namaClubDicari string

	fmt.Println("Silahkan masukkan nama club yang ingin dirubah total serinya: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editTotalSeriPertandingan(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan perbarui total seri club", (A[indeks].nama))
			fmt.Scan(&A[indeks].nSeri)
			fmt.Println()
			fmt.Println("Total seri club", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nSeri)

			/*mencari npoint*/
			A[indeks].nPoint = (A[indeks].nKemenangan * 3) + (A[indeks].nSeri * 1) + (A[indeks].nKekalahan * 0)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editTotalSeriPertandingan(A, n, indeks)
		}
	}
}

func editTotalKekalahanPertandingan(A *tabInt, n *int, indeks int) {
	var namaClubDicari string

	fmt.Println("Silahkan masukkan nama club yang ingin dirubah total kekalahannya: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editTotalKekalahanPertandingan(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan perbarui total kekalahan club", (A[indeks].nama))
			fmt.Scan(&A[indeks].nKekalahan)
			fmt.Println()
			fmt.Println("Total kekalahan club", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nKekalahan)

			/*mencari npoint*/
			A[indeks].nPoint = (A[indeks].nKemenangan * 3) + (A[indeks].nSeri * 1) + (A[indeks].nKekalahan * 0)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editTotalKekalahanPertandingan(A, n, indeks)
		}
	}
}

func editTotalGolMasuk(A *tabInt, n *int, indeks int) {
	var namaClubDicari string

	fmt.Println("Silahkan masukkan nama club yang ingin dirubah total gol masuknya: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editTotalGolMasuk(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan perbarui total gol masuk club", (A[indeks].nama))
			fmt.Scan(&A[indeks].nGolMasuk)
			fmt.Println()
			fmt.Println("Total gol masuk club", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nGolMasuk)

			/*mencari selisih gol*/

			A[indeks].nSelisihGol = A[indeks].nGolMasuk - A[indeks].nGolKemasukan
			if A[indeks].nSelisihGol < 0 {
				A[indeks].nSelisihGol = A[indeks].nSelisihGol * -1
			}

			/*mencari npoint*/
			A[indeks].nPoint = (A[indeks].nKemenangan * 3) + (A[indeks].nSeri * 1) + (A[indeks].nKekalahan * 0)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editTotalGolMasuk(A, n, indeks)
		}
	}
}

func editTotalGolKemasukan(A *tabInt, n *int, indeks int) {
	var namaClubDicari string

	fmt.Println("Silahkan masukkan nama club yang ingin dirubah total gol kemasukannya: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editTotalGolKemasukan(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)

		fmt.Println()
		fmt.Println("Silahkan perbarui total gol kemasukan club", (A[indeks].nama))
		fmt.Scan(&A[indeks].nGolKemasukan)
		fmt.Println()
		fmt.Println("Total gol kemasukan club", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nGolKemasukan)

		/*mencari selisih gol*/

		A[indeks].nSelisihGol = A[indeks].nGolMasuk - A[indeks].nGolKemasukan
		if A[indeks].nSelisihGol < 0 {
			A[indeks].nSelisihGol = A[indeks].nSelisihGol * -1
		}

		/*mencari npoint*/
		A[indeks].nPoint = (A[indeks].nKemenangan * 3) + (A[indeks].nSeri * 1) + (A[indeks].nKekalahan * 0)

		if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editTotalGolKemasukan(A, n, indeks)
		}
	}
}

func editNamaClub(A *tabInt, n *int, indeks int) {
	var namaClubDicari string
	fmt.Println("Silahkan masukkan nama club yang ingin dirubah: ")
	fmt.Println("Z. Cancel")
	fmt.Println()
	fmt.Print("Jawab: ")
	fmt.Scan(&namaClubDicari)
	if namaClubDicari == "Z" {
		partOfMenuEdit(A, n)
	} else if len(namaClubDicari) == 1 || len(namaClubDicari) == 2 && namaClubDicari != "Z" {
		fmt.Println()
		fmt.Println("!! ERROR !! Masukkan anda tidak valid")
		fmt.Println("Masukkan nama club (3 Karakter) atau huruf kapital (Z) untuk cancel")
		fmt.Println()
		editNamaClub(A, n, indeks)

	} else {
		indeks = findData(*A, *n, namaClubDicari)
		if indeks > -1 {
			fmt.Println()
			fmt.Println("Silahkan perbarui nama club ke-", (indeks + 1))
			fmt.Scan(&A[indeks].nama)
			fmt.Println()
			fmt.Println("Nama club ke-", A[indeks].nama, "berhasil diperbaharui menjadi", A[indeks].nama)

		} else if indeks == -1 {
			fmt.Println()
			fmt.Println("Nama tidak ditemukan")
			fmt.Println("Silahkan masukkan nama yang terdapat pada data")
			fmt.Println()
			editNamaClub(A, n, indeks)
		}
	}
}

func subMenuEdit(A *tabInt, n *int) {
	var jawabsubmenu string
	var indeksClub int

	for {

		fmt.Println()
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println("Apa yang ingin anda edit?")
		fmt.Println("A. Nama club")
		fmt.Println("B. Semua data pertandingan club")
		fmt.Println("C. Jumlah pertandingannya")
		fmt.Println("D. Total Kemenangan pertandingan")
		fmt.Println("E. Total seri pertandingan")
		fmt.Println("F. Total kekalahan pertandingan")
		fmt.Println("G. Jumlah gol masuk")
		fmt.Println("H. Jumlah gol kemasukan")
		fmt.Println("I. Back to Menu Ubah Data")
		fmt.Println("J. Cancel")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println()
		fmt.Println("Masukkan jawaban yang valid (A/B/C/D/E/F/G/H/I/J)")
		fmt.Print("Jawab: ")
		fmt.Scan(&jawabsubmenu)
		fmt.Println()
		if jawabsubmenu == "A" || jawabsubmenu == "B" || jawabsubmenu == "C" || jawabsubmenu == "D" || jawabsubmenu == "E" || jawabsubmenu == "F" || jawabsubmenu == "G" || jawabsubmenu == "H" || jawabsubmenu == "I" || jawabsubmenu == "J" {
			break
		}
	}
	if jawabsubmenu == "A" {
		editNamaClub(A, n, indeksClub)

	} else if jawabsubmenu == "B" {
		editSemuaDataPertandingan(A, n, indeksClub)

	} else if jawabsubmenu == "C" {
		editJumlahPertandingan(A, n, indeksClub)

	} else if jawabsubmenu == "D" {
		editTotalKemenanganPertandingan(A, n, indeksClub)

	} else if jawabsubmenu == "E" {
		editTotalSeriPertandingan(A, n, indeksClub)

	} else if jawabsubmenu == "F" {
		editTotalKekalahanPertandingan(A, n, indeksClub)

	} else if jawabsubmenu == "G" {
		editTotalGolMasuk(A, n, indeksClub)

	} else if jawabsubmenu == "H" {
		editTotalGolKemasukan(A, n, indeksClub)

	} else if jawabsubmenu == "I" {
		partOfMenuEdit(A, n)

	} else if jawabsubmenu == "J" {
		menuEdit(A, n)
	} else {
		fmt.Println()
		fmt.Println("!! ERROR !! : Masukkan anda tidak valid")
		fmt.Println("Silahkan masukkan kembali huruf kapital yang valid (A/B/C/D/E/F/G/H/I/J)")
		subMenuEdit(A, n)
	}
}

func descendingnPoint(A *tabInt, n *int) {
	var i, idx, pass int
	var temp dataklub

	for pass = 1; pass <= *n-1; pass++ {

		idx = pass - 1
		for i = pass; i <= *n-1; i++ {

			if (A[i].nPoint > A[idx].nPoint) || (A[i].nPoint == A[idx].nPoint && A[i].nSelisihGol > A[idx].nSelisihGol) {

				idx = i

			}

		}

		temp = A[idx]
		A[idx] = A[pass-1]
		A[pass-1] = temp
	}
	fmt.Println("Data club telah kami urutkan dari yang terbesar ke yang terkecil")
	fmt.Println("Silahlan cetak data terbaru untuk melihat ranking club")
	menuEdit(A, n)
}

func partOfMenuEdit(A *tabInt, n *int) {
	var jawabanmenu string
	var indeksClub int
	fmt.Println()
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("#---------------------------Menu ubah data---------------------------#")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println()
	fmt.Println("Apa yang ingin anda lakukan?")
	fmt.Println("A. Edit data")
	fmt.Println("B. Hapus data")
	fmt.Println("C. Tambah data")
	fmt.Println("D. Cancel")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println()
	fmt.Println("Masukkan jawaban yang valid (A/B/C/D)")
	fmt.Print("jawab: ")
	fmt.Scan(&jawabanmenu)
	if jawabanmenu == "A" {
		subMenuEdit(A, n)
	} else if jawabanmenu == "B" {
		hapusData(A, n, indeksClub)

	} else if jawabanmenu == "C" {
		tambahData(A, n)
	} else if jawabanmenu == "D" {
		menuEdit(A, n)
	} else {
		fmt.Println()
		fmt.Println("!! ERROR !! : Masukkan anda tidak valid")
		fmt.Println("Silahkan masukkan kembali huruf kapital yang valid (A/B/C)")
		partOfMenuEdit(A, n)
	}
}

func menuEdit(A *tabInt, n *int) {
	var jawaban string
	jawaban = "A"
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------------------------------------------")
	for {
		if jawaban == "B" {
			break
		}
		fmt.Println()
		fmt.Println("Apakah anda ingin melakukan perubahan dengan data tersebut?")
		fmt.Println("A. Iya")
		fmt.Println("B. Tidak")
		fmt.Println("C. Cetak data terbaru")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Println()
		fmt.Println("Masukkan jawaban yang valid (A/B/C)")
		fmt.Print("Jawab: ")
		fmt.Scan(&jawaban)
		if jawaban == "A" {
			partOfMenuEdit(A, n)
		} else if jawaban == "B" {
			menuDataFix(A, n)
		} else if jawaban == "C" {
			fmt.Println()
			fmt.Println("Menampilkan data terbaru yang kami simpan:")

			cetakData(A, n)
		} else {
			fmt.Println()
			fmt.Println("ERROR : Masukkan anda tidak valid")
			fmt.Println("Silahkan masukkan kembali huruf kapital yang valid (A/B/C)")

		}

	}
}

func menuDataFix(A *tabInt, n *int) {
	var jawabMenuDataFix string
	fmt.Println()
	fmt.Println("Apa yang ingin anda lakukan selanjutnya?")
	fmt.Println("A. Urutkan ranking club berdasarkan jumlah point")
	fmt.Println("B. Keluar dari aplikasi")
	fmt.Println("C. Cancel")
	fmt.Println("----------------------------------------------------------------------")
	fmt.Print("jawab: ")
	fmt.Scan(&jawabMenuDataFix)
	if jawabMenuDataFix == "A" {
		descendingnPoint(A, n)
	} else if jawabMenuDataFix == "C" {
		menuEdit(A, n)
	} else if jawabMenuDataFix == "B" {
		fmt.Println("Aplikasi akan segera ditutup")
		fmt.Println("Semua data anda akan dihapus")
	}
}

func main() {
	var data tabInt
	var ndata int

	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("#----------------------- APLIKASI EPL MANAGER -----------------------#")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println("#--------------------------------------------------------------------#")
	fmt.Println()
	for {

		fmt.Println("Berapa banyak club yang ingin anda inputkan dalam liga ini?")
		fmt.Println("(Jawab dengan menginputkan bilangan bulat positif(MAXIMAL 19 TEAM))")
		fmt.Println("----------------------------------------------------------------------")
		fmt.Print("Jawab: ")
		fmt.Scan(&ndata)
		if ndata <= 19 {

			fmt.Println()
			break
		}
	}

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Silahkan input data club")
	fmt.Println()
	inputNamaClub(&data, ndata)
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("Data anda telah kami rekam")
	cetakData(&data, &ndata)
	menuEdit(&data, &ndata)

}
