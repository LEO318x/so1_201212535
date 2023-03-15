package ram

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type ModuloMemoria struct {
	MemTotal  int `json:"MemTotal"`
	MemLibre  int `json:"MemLibre"`
	MemBuffer int `json:"MemBuffer"`
}

type DataMemoriaRam struct {
	MemTotal     int     `json:"MemTotal"`
	MemConsumida int     `json:"MemConsumida"`
	RamUsada     float64 `json:"RamUsada"`
}

// ////////////MEMORIA
// Retornamos la memoria cache
func returnCacheMemory() int {
	cmd := exec.Command("sh", "-c", "free -m |head -n 2 |tail -1 |awk {'print $6'}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := strings.Trim(string(out[:]), "\n")
	intVar, err := strconv.Atoi(output)
	if err != nil {
		fmt.Println("Error al retornar memoria cache")
		return 0
	}
	return intVar
}

// Retornamos información del modulo de RAM
func ModuloRAM() DataMemoriaRam {
	var memoriaTotal int
	var memoriaConsumida int
	var memoriaLibre int
	var infMemoria DataMemoriaRam

	cmd := exec.Command("sh", "-c", "cat /proc/ram_201212535")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])

	modMemoria := ModuloMemoria{}

	err2 := json.Unmarshal([]byte(output), &modMemoria)
	if err2 != nil {
		fmt.Println("Error al leer json modulo de memoria")
		return infMemoria
	}
	memoriaTotal = modMemoria.MemTotal

	memoriaLibre = modMemoria.MemLibre + returnCacheMemory()
	memoriaConsumida = memoriaTotal - memoriaLibre
	porcentajeRam := (float64(memoriaConsumida) / float64(memoriaTotal)) * 100

	infMemoria.MemTotal = memoriaTotal
	infMemoria.MemConsumida = memoriaConsumida
	infMemoria.RamUsada = porcentajeRam

	return infMemoria
}
