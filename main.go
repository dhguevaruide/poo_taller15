package main

import (
	"fmt"
	"poo_taller15/engine"
	"poo_taller15/model"
)

func main() {
	fmt.Println("=== SISTEMA DE EVALUACIÓN DOCENTE - PATRÓN STRATEGY ===\n")

	calculador := engine.NewCalculador()

	// Escenario 1
	fmt.Println("--- ESCENARIO 1 ---")
	e1 := model.Evaluacion{ID: "E1", DocenteID: "D1", CompetenciaID: "C1", PuntajeBase: 80}
	d1 := model.Docente{ID: "D1", Nombre: "Ana López", Departamento: "Ciencias", Cargo: "Profesor"}
	c1 := model.Competencia{ID: "C1", Nombre: "Investigación Avanzada", Tipo: "investigacion"}

	p1, r1 := calculador.CalcularPuntajeConReglas(e1, d1, c1, 8)
	e1.PuntajeFinal = p1
	fmt.Println(engine.GenerarReporteJustificado(e1, r1))
	fmt.Println()

	// Escenario 2
	fmt.Println("--- ESCENARIO 2 ---")
	e2 := model.Evaluacion{ID: "E2", DocenteID: "D2", CompetenciaID: "C2", PuntajeBase: 75}
	d2 := model.Docente{ID: "D2", Nombre: "Carlos Mendoza", Departamento: "Ingeniería", Cargo: "Director"}
	c2 := model.Competencia{ID: "C2", Nombre: "Gestión de Proyectos", Tipo: "gestion"}

	p2, r2 := calculador.CalcularPuntajeConReglas(e2, d2, c2, 12)
	e2.PuntajeFinal = p2
	fmt.Println(engine.GenerarReporteJustificado(e2, r2))
	fmt.Println()

	// Escenario 3
	fmt.Println("--- ESCENARIO 3 ---")
	e3 := model.Evaluacion{ID: "E3", DocenteID: "D3", CompetenciaID: "C3", PuntajeBase: 85}
	d3 := model.Docente{ID: "D3", Nombre: "Elena Vargas", Departamento: "Ciencias", Cargo: "Director"}
	c3 := model.Competencia{ID: "C3", Nombre: "Gestión Académica", Tipo: "gestion"}

	p3, r3 := calculador.CalcularPuntajeConReglas(e3, d3, c3, 15)
	e3.PuntajeFinal = p3
	fmt.Println(engine.GenerarReporteJustificado(e3, r3))
	fmt.Println()

	// Escenario 4
	fmt.Println("--- ESCENARIO 4 ---")
	e4 := model.Evaluacion{ID: "E4", DocenteID: "D4", CompetenciaID: "C4", PuntajeBase: 90}
	d4 := model.Docente{ID: "D4", Nombre: "Miguel Torres", Departamento: "Ciencias", Cargo: "Profesor"}
	c4 := model.Competencia{ID: "C4", Nombre: "Investigación Aplicada", Tipo: "investigacion"}

	p4, r4 := calculador.CalcularPuntajeConReglas(e4, d4, c4, 12)
	e4.PuntajeFinal = p4
	fmt.Println(engine.GenerarReporteJustificado(e4, r4))
	fmt.Println()
}
