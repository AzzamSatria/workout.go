package main
import "fmt"

type Workout struct {
	Jenis, Tanggal string
	Durasi, Kalori int
}

const maxData = 500
var dataWorkout [maxData]Workout
var n int

func main() {
	repeat := true
	for repeat {
		fmt.Println("\nAplikasi Manajemen dan Pencatatan Workout Harian")
		fmt.Println("1. Tambah Data Workout")
		fmt.Println("2. Ubah Data Workout")
		fmt.Println("3. Hapus Data Workout")
		fmt.Println("4. Rekomendasi Workout")
		fmt.Println("5. Cari Workout")
		fmt.Println("6. Urutkan Workout")
		fmt.Println("7. Tampilkan Laporan")
		fmt.Println("8. Keluar")

		var pilihan int
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahWorkout()
		} else if pilihan == 2 {
			ubahWorkout()
		} else if pilihan == 3 {
			hapusWorkout()
		} else if pilihan == 4 {
			rekomendasiWorkout()
		} else if pilihan == 5 {
			cariWorkout()
		} else if pilihan == 6 {
			urutkanWorkout()
		} else if pilihan == 7 {
			tampilkanLaporan()
		} else if pilihan == 8 {
			fmt.Println("Terima kasih telah menggunakan aplikasi ini.")
			repeat = false
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahWorkout() {
	if n < maxData {
		var jenis, tanggal string
		var durasi, kalori int
		fmt.Print("Masukkan jenis workout: ")
		fmt.Scan(&jenis)
		fmt.Print("Masukkan tanggal (tanggal-bulan-tahun): ")
		fmt.Scan(&tanggal)
		fmt.Print("Masukkan durasi (menit): ")
		fmt.Scan(&durasi)
		fmt.Print("Masukkan kalori yang terbakar: ")
		fmt.Scan(&kalori)
		dataWorkout[n] = Workout{jenis, tanggal, durasi, kalori}
		n++
	} else {
		fmt.Println("Data sudah penuh")
	}
	fmt.Println("Data berhasil ditambahkan")
}

func ubahWorkout() {
	if n == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}
	fmt.Println("Daftar Workout:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s - %d menit - %d kalori - %s\n", i+1, dataWorkout[i].Jenis, dataWorkout[i].Durasi, dataWorkout[i].Kalori, dataWorkout[i].Tanggal)
	}
	var indeks int
	fmt.Print("Pilih nomor workout yang ingin diubah: ")
	fmt.Scan(&indeks)
	indeks--
	if indeks >= 0 && indeks < n {
		fmt.Print("Masukkan jenis workout baru: ")
		fmt.Scan(&dataWorkout[indeks].Jenis)
		fmt.Print("Masukkan tanggal baru (tanggal-bulan-tahun): ")
		fmt.Scan(&dataWorkout[indeks].Tanggal)
		fmt.Print("Masukkan durasi baru (menit): ")
		fmt.Scan(&dataWorkout[indeks].Durasi)
		fmt.Print("Masukkan kalori baru: ")
		fmt.Scan(&dataWorkout[indeks].Kalori)
	} else {
		fmt.Println("Nomor tidak valid")
	}
	fmt.Println("Data berhasil diubah")
}

func hapusWorkout() {
	if n == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}
	fmt.Println("Daftar Workout:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s - %d menit - %d kalori - %s\n", i+1, dataWorkout[i].Jenis, dataWorkout[i].Durasi, dataWorkout[i].Kalori, dataWorkout[i].Tanggal)
	}
	var indeks int
	fmt.Print("Pilih nomor workout yang ingin dihapus: ")
	fmt.Scan(&indeks)
	indeks--
	if indeks >= 0 && indeks < n {
		for i := indeks; i < n-1; i++ {
			dataWorkout[i] = dataWorkout[i+1]
		}
		n--
		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Nomor tidak valid")
	}
}

func rekomendasiWorkout() {
	var hitung [maxData]int
	var totalKalori [maxData]int
	var jenisList [maxData]string
	var dataunik int

	for i := 0; i < n; i++ {
		found := false
		for j := 0; j < dataunik; j++ {
			if jenisList[j] == dataWorkout[i].Jenis {
				hitung[j]++
				totalKalori[j] += dataWorkout[i].Kalori
				found = true
			}
		}
		if !found {
			jenisList[dataunik] = dataWorkout[i].Jenis
			hitung[dataunik] = 1
			totalKalori[dataunik] = dataWorkout[i].Kalori
			dataunik++
		}
	}
	indeksFrekuensiMax := 0
	indeksKaloriMax := 0
	for i := 1; i < dataunik; i++ {
		if hitung[i] > hitung[indeksFrekuensiMax] {
			indeksFrekuensiMax = i
		}
		if totalKalori[i] > totalKalori[indeksKaloriMax] {
			indeksKaloriMax = i
		}
	}
	fmt.Println("Rekomendasi workout: ", jenisList[indeksFrekuensiMax], "dan", jenisList[indeksKaloriMax])
}

func cariWorkout() {
	if n == 0 {
		fmt.Println("Belum ada data workout.")
		return
	}
	var jenis string
	var metode int
	fmt.Println("Metode Pencarian:")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&metode)
	fmt.Print("Masukkan jenis workout yang dicari: ")
	fmt.Scan(&jenis)

	var indeks int = -1
	if metode == 1 {
		indeks = sequentialSearch(jenis)
	} else if metode == 2 {
		indeks = binarySearchJenis(jenis)
	} else {
		fmt.Println("Metode tidak valid.")
		return
	}
	if indeks != -1 {
		fmt.Printf("%d. %s - %d menit - %d kalori - %s\n", indeks+1, dataWorkout[indeks].Jenis, dataWorkout[indeks].Durasi, dataWorkout[indeks].Kalori, dataWorkout[indeks].Tanggal)
	} else {
		fmt.Println("Workout tidak ditemukan.")
	}
}

func urutkanWorkout() {
	var kategori, urutan int
	fmt.Println("Urutkan berdasarkan:")
	fmt.Println("1. Durasi")
	fmt.Println("2. Kalori")
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&kategori)
	fmt.Println("Urutan:")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
	fmt.Print("Pilih urutan: ")
	fmt.Scan(&urutan)

	if kategori == 1 {
		insertionSortDurasi(urutan == 1)
	} else {
		selectionSortKalori(urutan == 1)
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s - %d menit - %d kalori - %s\n", i+1, dataWorkout[i].Jenis, dataWorkout[i].Durasi, dataWorkout[i].Kalori, dataWorkout[i].Tanggal)
	}
}

func tampilkanLaporan() {
	fmt.Println("10 aktivitas terakhir:")
	laporan := 0
	if n > 10 {
		laporan = n - 10
	}
	totalKalori := 0
	for i := laporan; i < n; i++ {
		fmt.Printf("%d. %s - %d menit - %d kalori - %s\n", i+1, dataWorkout[i].Jenis, dataWorkout[i].Durasi, dataWorkout[i].Kalori, dataWorkout[i].Tanggal)
		totalKalori += dataWorkout[i].Kalori
	}
	fmt.Println("Total kalori yang dihasilkan: ", totalKalori , "kalori")
}

func sequentialSearch(jenis string) int {
	for i := 0; i < n; i++ {
		if dataWorkout[i].Jenis == jenis {
			return i
		}
	}
	return -1
}

func binarySearchJenis(jenis string) int {
	insertionSortJenis() 
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if dataWorkout[mid].Jenis == jenis {
			return mid
		} else if dataWorkout[mid].Jenis < jenis {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func insertionSortJenis() {
	for i := 1; i < n; i++ {
		temp := dataWorkout[i]
		j := i - 1
		for j >= 0 && dataWorkout[j].Jenis > temp.Jenis {
			dataWorkout[j+1] = dataWorkout[j]
			j--
		}
		dataWorkout[j+1] = temp
	}
}


func insertionSortDurasi(asc bool) {
	for i := 1; i < n; i++ {
		temp := dataWorkout[i]
		j := i - 1
		for j >= 0 && ((asc && dataWorkout[j].Durasi > temp.Durasi) || (!asc && dataWorkout[j].Durasi < temp.Durasi)) {
			dataWorkout[j+1] = dataWorkout[j]
			j--
		}
		dataWorkout[j+1] = temp
	}
}

func selectionSortKalori(asc bool) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (asc && dataWorkout[j].Kalori < dataWorkout[idx].Kalori) || (!asc && dataWorkout[j].Kalori > dataWorkout[idx].Kalori) {
				idx = j
			}
		}
		temp := dataWorkout[i]
		dataWorkout[i] = dataWorkout[idx]
		dataWorkout[idx] = temp
	}
}