package main

import (
	"fmt"
)

const NMAX = 3

type mahasiswa struct {
	NIM          int
	NamaLengkap  string
	Alamat       string
	TanggalLahir string
	JenisKelamin string
	Kontak       string
	Jurusan      [3]string
}

type tabJurusan struct {
	namaJurusan       string
	ketersediaanKursi int
}

var daftarMahasiswa [NMAX]mahasiswa
var jumlahMahasiswa int
var nimMahasiswa int

var daftarJurusan = [3]tabJurusan{
	{namaJurusan: "Teknik Informatika", ketersediaanKursi: 1},
	{namaJurusan: "Sistem Informasi", ketersediaanKursi: 3},
	{namaJurusan: "Rekayasa Perangkat Lunak", ketersediaanKursi: 2},
}

func main() {
	menuUtama()
}

func menuUtama() {
	fmt.Println("=================================")
	fmt.Println("   Sistem Pendaftaran Mahasiswa  ")
	fmt.Println("=================================")
	fmt.Println("1. Pendaftaran Mahasiswa")
	fmt.Println("2. Tampilkan Data Mahasiswa")
	fmt.Println("3. Cari Data Mahasiswa")
	fmt.Println("4. Ubah Data Mahasiswa")
	fmt.Println("5. Hapus Data Mahasiswa")
	fmt.Println("6. Keluar")
	fmt.Println()
	fmt.Println("Aslam Rizky Fadillah 103022330152")
	fmt.Println("Riyanda Wiesya Bella 103022300001")
	fmt.Println("=================================")
	fmt.Print("Pilih menu: ")

	var pilihan int
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		pendaftaranMahasiswa()
	} else if pilihan == 2 {
		tampilkanDataMahasiswa()
	} else if pilihan == 3 {
		cariDataMahasiswa()
	} else if pilihan == 4 {
		ubahDataMahasiswa()
	} else if pilihan == 5 {
		hapusDataMahasiswa()
	} else if pilihan == 6 {
		fmt.Println("Keluar dari program...")
		return
	} else {
		fmt.Println("Pilihan tidak valid!")
		menuUtama()
	}
}

func pendaftaranMahasiswa() {
	fmt.Println("Pendaftaran Mahasiswa Baru")
	fmt.Println("=================================")

	var namaLengkap, alamat, tanggalLahir, jenisKelamin, kontak string
	var jurusanPilihan, banyakJurusan int
	var jurusanPilihanArr [3]string
	var i, j int

	if jumlahMahasiswa >= NMAX {
		fmt.Println("Kuota mahasiswa sudah penuh!")
		menuUtama()
	}

	fmt.Println("NIM Anda:", nimMahasiswa+1)
	fmt.Print("Nama Lengkap: ")
	fmt.Scan(&namaLengkap)
	fmt.Print("Alamat: ")
	fmt.Scan(&alamat)
	fmt.Print("Tanggal Lahir (YYYY-MM-DD): ")
	fmt.Scan(&tanggalLahir)
	fmt.Print("Informasi Kontak: ")
	fmt.Scan(&kontak)
	fmt.Print("Jenis Kelamin (L/P): ")
	fmt.Scan(&jenisKelamin)

	for jenisKelamin != "L" && jenisKelamin != "P" {
		fmt.Println("Masukkan Tidak Valid!")
		fmt.Print("Jenis Kelamin (L/P): ")
		fmt.Scan(&jenisKelamin)
	}

	fmt.Println("Daftar Jurusan:")
	for i = 0; i < 3; i++ {
		fmt.Printf("%d. %s (Ketersediaan Kursi: %d)\n", i+1, daftarJurusan[i].namaJurusan, daftarJurusan[i].ketersediaanKursi)
	}

	fmt.Println("Ingin Milih Berapa Jurusan: ")
	fmt.Scan(&banyakJurusan)

	for i := 0; i < banyakJurusan; i++ {
		fmt.Print("Pilih Jurusan (nomor): ")
		fmt.Scan(&jurusanPilihan)

		for j = 0; j < 3; j++ {
			for jurusanPilihanArr[j] == daftarJurusan[jurusanPilihan-1].namaJurusan {
				fmt.Println("Anda sudah memilih jurusan ini!")
				fmt.Print("Pilih Jurusan (nomor): ")
				fmt.Scan(&jurusanPilihan)
			}
		}

		if jurusanPilihan < 1 || jurusanPilihan > 3 {
			fmt.Println("Jurusan tidak valid, silakan pilih kembali!")
			i--
		} else if daftarJurusan[jurusanPilihan-1].ketersediaanKursi <= 0 {
			fmt.Println("Maaf, jurusan yang dipilih sudah penuh!")
			i--
		} else {
			jurusanPilihanArr[i] = daftarJurusan[jurusanPilihan-1].namaJurusan
			daftarJurusan[jurusanPilihan-1].ketersediaanKursi--
			fmt.Println("Anda telah memilih jurusan ", daftarJurusan[jurusanPilihan-1].namaJurusan)
		}
	}

	daftarMahasiswa[jumlahMahasiswa].NIM = nimMahasiswa + 1
	daftarMahasiswa[jumlahMahasiswa].NamaLengkap = namaLengkap
	daftarMahasiswa[jumlahMahasiswa].Alamat = alamat
	daftarMahasiswa[jumlahMahasiswa].TanggalLahir = tanggalLahir
	daftarMahasiswa[jumlahMahasiswa].JenisKelamin = jenisKelamin
	daftarMahasiswa[jumlahMahasiswa].Kontak = kontak
	daftarMahasiswa[jumlahMahasiswa].Jurusan = jurusanPilihanArr

	if jumlahMahasiswa < NMAX {
		jumlahMahasiswa++
		nimMahasiswa++
		fmt.Println("Mahasiswa berhasil didaftarkan!")
	}

	menuUtama()
}

func tampilkanDataMahasiswa() {
	var i, k int
	var mhs mahasiswa

	if jumlahMahasiswa > 0 {
		fmt.Println("Data Mahasiswa")
		fmt.Println("=================================")

		for i = 0; i < jumlahMahasiswa; i++ {
			mhs = daftarMahasiswa[i]
			fmt.Printf("NIM: %d\nNama Lengkap: %s\nAlamat: %s\nTanggal Lahir: %s\nInformasi Kontak: %s\nJenis Kelamin: %s\n",
				mhs.NIM, mhs.NamaLengkap, mhs.Alamat, mhs.TanggalLahir, mhs.Kontak, mhs.JenisKelamin)
			fmt.Println("Jurusan:")
			for k = 0; k < 3; k++ {
				fmt.Println(mhs.Jurusan[k])
			}
		}
		menuUtama()
	}
	fmt.Println("Data mahasiswa tidak ditemukan!")
	menuUtama()
}

func cariDataMahasiswa() {
	var k int
	var nim int
	var left, right, mid, found int

	fmt.Println("Cari Data Mahasiswa")
	fmt.Println("=================================")

	if jumlahMahasiswa > 0 {
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&nim)

		left = 0
		right = jumlahMahasiswa - 1
		found = -1

		for left <= right && found == -1 {
			mid = (left + right) / 2
			if nim < daftarMahasiswa[mid].NIM {
				right = mid - 1
			} else if nim > daftarMahasiswa[mid].NIM {
				left = mid + 1
			} else {
				found = mid
			}
		}
		if found != -1 {
			fmt.Printf("NIM: %d\nNama Lengkap: %s\nAlamat: %s\nTanggal Lahir: %s\nInformasi Kontak: %s\nJenis Kelamin: %s\n",
				daftarMahasiswa[found].NIM, daftarMahasiswa[found].NamaLengkap, daftarMahasiswa[found].Alamat, daftarMahasiswa[found].TanggalLahir, daftarMahasiswa[found].Kontak, daftarMahasiswa[found].JenisKelamin)
			fmt.Println("Jurusan: ")
			for k = 0; k < 3; k++ {
				fmt.Println(daftarMahasiswa[found].Jurusan[k])
			}
			menuUtama()
		}
	}
	fmt.Println("Data mahasiswa tidak ditemukan!")
	menuUtama()
}

func ubahDataMahasiswa() {
	var i, k, l, m, n int
	var nim int
	var namaLengkap, alamat, tanggalLahir, kontak, jenisKelamin string
	var jurusanPilihan, banyakJurusan int
	var jurusanPilihanArr [3]string
	var left, right, mid, found int

	fmt.Println("Ubah Data Mahasiswa")
	fmt.Println("=================================")

	if jumlahMahasiswa > 0 {
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&nim)

		left = 0
		right = jumlahMahasiswa - 1
		found = -1

		for left <= right && found == -1 {
			mid = (left + right) / 2
			if nim < daftarMahasiswa[mid].NIM {
				right = mid - 1
			} else if nim > daftarMahasiswa[mid].NIM {
				left = mid + 1
			} else {
				found = mid
			}
		}

		if found != -1 {
			for k = 0; k < 3; k++ {
				for l = 0; l < 3; l++ {
					if daftarJurusan[l].namaJurusan == daftarMahasiswa[found].Jurusan[k] {
						daftarJurusan[l].ketersediaanKursi++
					}
				}
			}

			fmt.Print("Nama Lengkap: ")
			fmt.Scan(&namaLengkap)
			fmt.Print("Alamat: ")
			fmt.Scan(&alamat)
			fmt.Print("Tanggal Lahir (YYYY-MM-DD): ")
			fmt.Scan(&tanggalLahir)
			fmt.Print("Informasi Kontak: ")
			fmt.Scan(&kontak)
			fmt.Print("Jenis Kelamin (L/P): ")
			fmt.Scan(&jenisKelamin)

			for jenisKelamin != "L" && jenisKelamin != "P" {
				fmt.Println("Masukkan Tidak Valid!")
				fmt.Print("Jenis Kelamin (L/P): ")
				fmt.Scan(&jenisKelamin)
			}

			fmt.Println("Daftar Jurusan:")

			for m = 0; m < 3; m++ {
				fmt.Printf("%d. %s (Ketersediaan Kursi: %d)\n", m+1, daftarJurusan[m].namaJurusan, daftarJurusan[m].ketersediaanKursi)
			}

			fmt.Println("Ingin Milih Berapa Jurusan: ")
			fmt.Scan(&banyakJurusan)

			for i = 0; i < banyakJurusan; i++ {
				fmt.Print("Pilih Jurusan (nomor): ")
				fmt.Scan(&jurusanPilihan)

				for n = 0; n < 3; n++ {
					for jurusanPilihanArr[n] == daftarJurusan[jurusanPilihan-1].namaJurusan {
						fmt.Println("Anda sudah memilih jurusan ini!")
						fmt.Print("Pilih Jurusan (nomor): ")
						fmt.Scan(&jurusanPilihan)
					}
				}

				if jurusanPilihan < 1 || jurusanPilihan > 3 {
					fmt.Println("Jurusan tidak valid, silakan pilih kembali!")
					n--
				} else if daftarJurusan[jurusanPilihan-1].ketersediaanKursi <= 0 {
					fmt.Println("Maaf, jurusan yang dipilih sudah penuh!")
					n--
				} else {
					jurusanPilihanArr[i] = daftarJurusan[jurusanPilihan-1].namaJurusan
					daftarJurusan[jurusanPilihan-1].ketersediaanKursi--
					fmt.Println("Anda telah memilih jurusan ", daftarJurusan[jurusanPilihan-1].namaJurusan)
				}
			}

			daftarMahasiswa[found].NamaLengkap = namaLengkap
			daftarMahasiswa[found].Alamat = alamat
			daftarMahasiswa[found].TanggalLahir = tanggalLahir
			daftarMahasiswa[found].JenisKelamin = jenisKelamin
			daftarMahasiswa[found].Kontak = kontak
			daftarMahasiswa[found].Jurusan = jurusanPilihanArr

			fmt.Println("Data mahasiswa berhasil diubah!")
			menuUtama()
		}
	}
	fmt.Println("Data mahasiswa tidak ditemukan!")
	menuUtama()
}

func hapusDataMahasiswa() {
	var i, j, k int
	var nim int
	var left, right, mid, found int

	fmt.Println("Hapus Data Mahasiswa")
	fmt.Println("=================================")

	if jumlahMahasiswa > 0 {
		fmt.Print("Masukkan NIM: ")
		fmt.Scan(&nim)

		left = 0
		right = jumlahMahasiswa - 1
		found = -1

		for left <= right && found == -1 {
			mid = (left + right) / 2
			if nim < daftarMahasiswa[mid].NIM {
				right = mid - 1
			} else if nim > daftarMahasiswa[mid].NIM {
				left = mid + 1
			} else {
				found = mid
			}
		}

		if found != -1 {

			for i = found; i < jumlahMahasiswa-1; i++ {
				daftarMahasiswa[i] = daftarMahasiswa[i+1]
			}

			daftarMahasiswa[jumlahMahasiswa-1] = mahasiswa{}

			for k = 0; k < 3; k++ {
				for j = 0; j < 3; j++ {
					if daftarJurusan[k].namaJurusan == daftarMahasiswa[found].Jurusan[j] {
						daftarJurusan[k].ketersediaanKursi++
					}
				}
			}

			jumlahMahasiswa--

			fmt.Println("Data mahasiswa berhasil dihapus!")
			menuUtama()
		}
	}
	fmt.Println("Data mahasiswa tidak ditemukan!")
	menuUtama()
}
