package main

import (
	"encoding/json"
	"log"
	"math"
)

type Actividad struct {
	Nombre   string  `json:"nombre"`
	Duracion float64 `json:"duracion"`

	// Tiempo más próximo, izquierda | derecha
	ProximoL float64 `json:"proximoL"`
	ProximoR float64 `json:"proximoR"`

	// Tiempo más lejano, izquierda | derecha
	LejanoL float64 `json:"lejanoL"`
	LejanoR float64 `json:"lejanoR"`
}

func NuevaActividad(n string, p float64) *Actividad {
	return &Actividad{n, p, NInf, NInf, Inf, Inf}
}

type RespuestaCPM struct {
	Actividades   []*Actividad `json:"actividades"`
	RutaCritica   []string     `json:"rutaCritica"`
	DuracionTotal float64      `json:"duracionTotal"`
}

func ResuelveCPM(p *[]Vertice) ([]byte, RespuestaCPM) {
	anterior := VerticesToAdjList(p, true)

	actividades := make(map[string]*Actividad, len(*p))
	actividades["-"] = &Actividad{"-", 0, 0, 0, 0, 0} // Inicio

	s := set()
	sn := set()

	for idx := 0; idx < len(*p); idx += 1 {
		actividades[(*p)[idx].Origen] = NuevaActividad((*p)[idx].Origen, (*p)[idx].Peso)

		s.Add((*p)[idx].Origen)
		// Invertimos porque el formato de la tabla es distinto
		(*p)[idx].Origen, (*p)[idx].Destino = (*p)[idx].Destino, (*p)[idx].Origen
		s.Add((*p)[idx].Origen)
		sn.Add((*p)[idx].Origen)
	}

	recorridoIda("-", VerticesToAdjList(p, true), actividades)

	duracionTotal := NInf
	for k := range s.SymmetricDifference(sn).m {
		duracionTotal = math.Max(duracionTotal, actividades[k].ProximoR)
	}

	// Actividades terminales
	for t := range s.SymmetricDifference(sn).m {
		actividades[t].LejanoR = duracionTotal
		actividades[t].LejanoL = duracionTotal - actividades[t].Duracion
		recorridoRegreso(t, anterior, actividades)
	}

	// Ruta crítica
	ruta := make([]string, 0, len(actividades))
	acti := make([]*Actividad, 0, len(actividades))

	for a, v := range actividades {
		acti = append(acti, v)

		if v.ProximoR == v.LejanoR {
			ruta = append(ruta, a)
		}
	}

	r := RespuestaCPM{acti, ruta, duracionTotal}
	resp, err := json.Marshal(r)

	if err != nil {
		log.Println(err)
		return nil, r
	} else {
		return resp, r
	}
}

func recorridoIda(anterior string, s map[string]map[string]float64, n map[string]*Actividad) {
	duracion := n[anterior].ProximoR

	for ve := range s[anterior] {
		if n[ve].ProximoL < duracion {
			n[ve].ProximoL = duracion
			n[ve].ProximoR = n[ve].Duracion + duracion
			recorridoIda(ve, s, n)
		}
	}
}

func recorridoRegreso(anterior string, s map[string]map[string]float64, n map[string]*Actividad) {
	duracion := n[anterior].LejanoL

	for ve := range s[anterior] {
		if n[ve].LejanoR > duracion {
			n[ve].LejanoR = duracion
			n[ve].LejanoL = duracion - n[ve].Duracion
			recorridoRegreso(ve, s, n)
		}
	}
}
