package main

import (
	"fmt"

	"poo_taller15/engine"
	"poo_taller15/model"
	"poo_taller15/strategy"
)

func main() {

	// =========================
	// REGISTRO DE REGLAS
	// =========================
	reglas := []strategy.EstrategiaPonderacion{
		strategy.ReglaCienciasInvestigacion{},
		strategy.ReglaDirectorGestion{},
		strategy.ReglaCoordinadorAcademica{},
		strategy.ReglaAdministracionGestion{},
		strategy.ReglaAntiguedad{}, // 👈 R3 AGREGADA AQUÍ
	}

	// =========================
	// ESCENARIO 1
	// Ciencias + Investigación
	// =========================
	doc1 := model.Docente{
		Nombre:       "Docente 1",
		Departamento: "Ciencias",
		Cargo:        "Profesor",
	}

	comp1 := model.Competencia{
		Nombre: "Investigación",
		Tipo:   "investigacion",
	}

	eval1 := model.Evaluacion{
		PuntajeBase: 100,
	}

	p1, r1 := engine.CalcularPuntajeConReglas(eval1, doc1, comp1, reglas)

	fmt.Println("\n=== ESCENARIO 1 ===")
	fmt.Println("Puntaje final:", p1)
	fmt.Println("Reglas aplicadas:", r1)

	// =========================
	// ESCENARIO 2
	// Director + Gestión
	// =========================
	doc2 := model.Docente{
		Nombre:       "Docente 2",
		Departamento: "Humanidades",
		Cargo:        "Director",
	}

	comp2 := model.Competencia{
		Nombre: "Gestión",
		Tipo:   "gestion",
	}

	eval2 := model.Evaluacion{
		PuntajeBase: 100,
	}

	p2, r2 := engine.CalcularPuntajeConReglas(eval2, doc2, comp2, reglas)

	fmt.Println("\n=== ESCENARIO 2 ===")
	fmt.Println("Puntaje final:", p2)
	fmt.Println("Reglas aplicadas:", r2)

	// =========================
	// ESCENARIO 3
	// Ciencias + Director + Gestión
	// =========================
	doc3 := model.Docente{
		Nombre:       "Docente 3",
		Departamento: "Ciencias",
		Cargo:        "Director",
	}

	comp3 := model.Competencia{
		Nombre: "Gestión",
		Tipo:   "gestion",
	}

	eval3 := model.Evaluacion{
		PuntajeBase: 100,
	}

	p3, r3 := engine.CalcularPuntajeConReglas(eval3, doc3, comp3, reglas)

	fmt.Println("\n=== ESCENARIO 3 ===")
	fmt.Println("Puntaje final:", p3)
	fmt.Println("Reglas aplicadas:", r3)

	// =========================
	// ESCENARIO 4
	// Ciencias + Investigación + Antigüedad
	// =========================
	doc4 := model.Docente{
		Nombre:          "Docente 4",
		Departamento:    "Ciencias",
		Cargo:           "Profesor",
		AniosAntiguedad: 15,
	}

	comp4 := model.Competencia{
		Nombre: "Investigación",
		Tipo:   "investigacion",
	}

	eval4 := model.Evaluacion{
		PuntajeBase: 100,
	}

	p4, r4 := engine.CalcularPuntajeConReglas(eval4, doc4, comp4, reglas)

	fmt.Println("\n=== ESCENARIO 4 ===")
	fmt.Println("Puntaje final:", p4)
	fmt.Println("Reglas aplicadas:", r4)
}