package main

import (
	"fmt"
	"time"

	"poo_taller15-main/engine"
	"poo_taller15-main/model"
)

func main() {

	evaluacion := model.Evaluacion{
		ID: "E1",
		Fecha: time.Now(),
		PuntajeBase: 100,
	}

	// Escenario 1
	doc1 := model.Docente{
		ID: "D1",
		Nombre: "Juan",
		Departamento: "Ciencias",
		Cargo: "Profesor",
		AniosAntiguedad: 5,
	}

	comp1 := model.Competencia{
		ID: "C1",
		Tipo: "investigación",
	}

	fmt.Println(
		engine.GenerarReporteJustificado(
			evaluacion,
			doc1,
			comp1,
		),
	)

	// Escenario 2
	doc2 := model.Docente{
		ID: "D2",
		Nombre: "Ana",
		Cargo: "Director",
		AniosAntiguedad: 8,
	}

	comp2 := model.Competencia{
		ID: "C2",
		Tipo: "gestión",
	}

	fmt.Println(
		engine.GenerarReporteJustificado(
			evaluacion,
			doc2,
			comp2,
		),
	)

	// Escenario 3
	doc3 := model.Docente{
		ID: "D3",
		Nombre: "Pedro",
		Departamento: "Ciencias",
		Cargo: "Director",
		AniosAntiguedad: 6,
	}

	comp3 := model.Competencia{
		ID: "C3",
		Tipo: "gestión",
	}

	fmt.Println(
		engine.GenerarReporteJustificado(
			evaluacion,
			doc3,
			comp3,
		),
	)

	// Escenario 4
	doc4 := model.Docente{
		ID: "D4",
		Nombre: "María",
		Departamento: "Ciencias",
		Cargo: "Profesor",
		AniosAntiguedad: 15,
	}

	comp4 := model.Competencia{
		ID: "C4",
		Tipo: "investigación",
	}

	fmt.Println(
		engine.GenerarReporteJustificado(
			evaluacion,
			doc4,
			comp4,
		),
	)
}