package main

import (
	"encoding/json"
	"fmt"
	"log"
	"pra2/cpu"
	"pra2/db"
	"pra2/ram"
	"time"
)

type dataProyecto struct {
	TablaProcesos []cpu.DataCpu `json:"TablaProcesos"`
}

func main() {
	bd, err := db.Conexion()
	if err != nil {
		log.Printf("Error con la base de datos" + err.Error())
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error al conectarse a la base de datos. Por favor verifica las credenciales, el error es: " + err.Error())
		}
	}

	for true {
		msg := dataProyecto{}
		raminfo := ram.ModuloRAM()
		cpup := cpu.CpuUso()
		cpuinfo1, cpuinfo2 := cpu.CpuTablaInfo()
		msg.TablaProcesos = cpuinfo1
		//j, _ := json.MarshalIndent(msg, "", "")
		j, _ := json.Marshal(msg)

		fmt.Println("########################################")
		fmt.Println("Guardando en BD...")
		_, err = bd.Exec("INSERT INTO ram (total, consumida, porc_utilizado) VALUES (?, ?, ?)", raminfo.MemTotal, raminfo.MemConsumida, raminfo.RamUsada)
		if err != nil {
			log.Printf("Hubo un error al insertar el registro en la base de datos | RAM")
		}

		_, err = bd.Exec("INSERT INTO cpu (porc_uso, running, suspended, stopped, zombie, total, procesos) VALUES (?, ?, ?, ?, ?, ?, ?)", cpup.Porcentaje, cpuinfo2.ProcesoEjecucion, cpuinfo2.ProcesoSuspendido, cpuinfo2.ProcesosDetenidos, cpuinfo2.ProcesosZombies, cpuinfo2.Total, string(j))
		if err != nil {
			log.Printf("Hubo un error al insertar el registro en la base de datos | CPU")
		}
		fmt.Println("########################################")

		time.Sleep(10 * time.Second)
	}

}
