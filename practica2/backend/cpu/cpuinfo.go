package cpu

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

// Porcentaje utilizado del cpu
type CPUPorcentaje struct {
	Porcentaje float64 `json:"Porcentaje"`
}

// struct que guarda la informacion de los procesos que luego son utilizadas en el front
type DataCpu struct {
	Nombre  string    `json:"Nombre"`
	UID     int       `json:"UID"`
	PID     int       `json:"PID"`
	Proceso string    `json:"Proceso"`
	Estado  string    `json:"Estado"`
	VmRSS   int       `json:"VmRSS"`
	Hijos   []DataCpu `json:"Hijos"`
}

type InfoGeneral struct {
	ProcesoEjecucion  int `json:"ProcesoEjecucion"`
	ProcesoSuspendido int `json:"ProcesoSuspendido"`
	ProcesosDetenidos int `json:"ProcesosDetenidos"`
	ProcesosZombies   int `json:"ProcesosZombies"`
	Total             int `json:"Total"`
}

// ///////////CPU
// Retornamos el username por un uid
func returnUsername(uid int) string {
	cmd := exec.Command("sh", "-c", "getent passwd "+strconv.Itoa(uid)+" | cut -d: -f1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])
	output = strings.Replace(output, "\n", "", -1)
	return output
}

// Obtenemos el porcentaje de uso (% de uso)
func CpuUso() CPUPorcentaje {
	cmd := exec.Command("sh", "-c", "top -bn 1 -i -c | head -n 3 | tail -1 | awk {'print $8'}")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])
	//quitamos espacios en blanco
	output2 := strings.Replace(output, "\n", "", -1)
	//eliminando saltos de linea y lo pasamos en un array
	var porcentajeUso float64
	porcentajeUso, err = strconv.ParseFloat(output2, 64)
	if err != nil {
		fmt.Println("error de conversion", err)
		porcentajeUso = 0
	}
	porcentajeUso = 100 - porcentajeUso
	var porcentaje CPUPorcentaje
	porcentaje.Porcentaje = porcentajeUso
	return porcentaje
}

// Retornamos información del modulo del CPU
func CpuTablaInfo() (cpu []DataCpu, resumen InfoGeneral) {

	var procesoEjecucion int = 0
	var procesoSuspendido int = 0
	var procesoDetenido int = 0
	var procesoZombie int = 0

	cmd := exec.Command("sh", "-c", "cat /proc/cpu_201212535")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	output := string(out[:])
	//convirtiendo json a struct
	moduloCpu := []DataCpu{}
	err2 := json.Unmarshal([]byte(output), &moduloCpu)
	if err2 != nil {
		fmt.Println("Error al leer json modulo de cpu", err2)
		//return moduloCpu
	}
	moduloCpu = moduloCpu[1:]
	//ciclo que formatea el json con el nombre del usuario
	for i, s := range moduloCpu {
		_ = i
		moduloCpu[i].Nombre = returnUsername(s.UID)
		moduloCpu[i].Estado = estados(moduloCpu[i].Estado)

		if moduloCpu[i].Estado == "Running" {
			procesoEjecucion += 1
		} else if moduloCpu[i].Estado == "Sleeping" {
			procesoSuspendido += 1
		} else if moduloCpu[i].Estado == "Zombie" {
			procesoZombie += 1
		} else if moduloCpu[i].Estado == "Stopped" {
			procesoDetenido += 1
		}

		moduloCpu[i].Hijos = moduloCpu[i].Hijos[1:]
		for j, hj := range moduloCpu[i].Hijos {
			_ = j
			nomH := returnUsername(hj.UID)
			moduloCpu[i].Hijos[j].Nombre = nomH
			moduloCpu[i].Hijos[j].Estado = estados(moduloCpu[i].Hijos[j].Estado)

			if moduloCpu[i].Hijos[j].Estado == "Running" {
				procesoEjecucion += 1
			} else if moduloCpu[i].Hijos[j].Estado == "Sleeping" {
				procesoSuspendido += 1
			} else if moduloCpu[i].Hijos[j].Estado == "Zombie" {
				procesoZombie += 1
			} else if moduloCpu[i].Hijos[j].Estado == "Stopped" {
				procesoDetenido += 1
			}
		}
	}

	var resultados InfoGeneral
	resultados.ProcesoEjecucion = procesoEjecucion
	resultados.ProcesoSuspendido = procesoSuspendido
	resultados.ProcesosZombies = procesoZombie
	resultados.ProcesosDetenidos = procesoDetenido
	resultados.Total = procesoEjecucion + procesoSuspendido + procesoZombie + procesoDetenido
	return moduloCpu, resultados
}

// funcion que evalua en que estado se encuentran los procesos
func estados(numEstado string) string {
	var tipoEstado string
	if numEstado == "0" {
		tipoEstado = "Running"
	} else if numEstado == "1" || numEstado == "1026" || numEstado == "2" {
		tipoEstado = "Sleeping"
	} else if numEstado == "4" {
		tipoEstado = "Zombie"
	} else if numEstado == "8" {
		tipoEstado = "Stopped"
	} else {
		tipoEstado = "No Reconocido"
	}
	return tipoEstado
}
