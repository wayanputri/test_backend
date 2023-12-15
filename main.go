package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pemain struct {
	ID   int
	Dadu []int
	Poin int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var jmlPemain int
	fmt.Print("masukkan jumlah pemain: ")
	_, errP := fmt.Scan(&jmlPemain)
	if errP != nil {
		fmt.Println("gagal input jumlah pemain")
		return
	}
	fmt.Println()

	var jmlDadu int
	fmt.Print("masukkan jumlah dadu: ")
	_, errD := fmt.Scan(&jmlDadu)
	if errD != nil {
		fmt.Println("gagal input jumlah dadu")
		return
	}

	pemainDadu := make([]*Pemain, jmlPemain)
	for i := 0; i < jmlPemain; i++ {
		pemainDadu[i] = &Pemain{ID: i + 1, Dadu: make([]int, jmlDadu)}
	}

	giliran := 1
	for {
		fmt.Printf("===============\ngiliran ke %d lempar dadu\n", giliran)

		status := Game(pemainDadu)
		if status {
			break
		}
		giliran++
	}

	maksimalPoin := 0
	var menang *Pemain
	for _, v := range pemainDadu {
		if v.Poin > maksimalPoin {
			menang = v
			maksimalPoin = v.Poin
		}
	}
	idTerakhir := pemainDadu[0].PemainTerakhir().ID
	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", idTerakhir)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.", menang.ID)

}

func (pemain *Pemain) PemainTerakhir() *Pemain {
	return &Pemain{
		ID:   pemain.ID,
		Dadu: pemain.Dadu,
		Poin: pemain.Poin,
	}
}
func Game(pemain []*Pemain) bool {
	for _, v := range pemain {
		for i := 0; i < len(v.Dadu); i++ {
			v.Dadu[i] = rand.Intn(6) + 1
		}
		fmt.Printf("Pemain #%d (%d): %v\n", v.ID, v.Poin, v.Dadu)
	}

	fmt.Println("setelah evaluasi:")
	for i, v := range pemain {
		if len(v.Dadu) > 0 {
			var daduBaru []int
			for _, dadu := range v.Dadu {
				switch dadu {
				case 1:
					if i == len(pemain)-1 {
						pemain[0].Dadu = append(pemain[0].Dadu, 1)
					} else {
						pemain[i+1].Dadu = append(pemain[i+1].Dadu, 1)
					}
				case 6:
					v.Poin++
				default:
					daduBaru = append(daduBaru, dadu)

				}
			}
			v.Dadu = daduBaru
		}
		fmt.Printf("Pemain #%d (%d): %v\n", v.ID, v.Poin, v.Dadu)
	}

	mainUlang := 0
	var pemainTerakhir *Pemain
	for _, v := range pemain {
		if len(v.Dadu) > 0 {
			mainUlang++
			pemainTerakhir = v
		}
	}
	pemainTerakhir.PemainTerakhir()
	return mainUlang == 1

}
