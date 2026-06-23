package main

import (
	"fmt"
	"poo_taller15/engine"
	"poo_taller15/model"
)

func main() {

	calculador := engine.NuevoCalculador()

	fmt.Println("ESCENARIO 1")

	docente1 := model.Docente{
		Nombre: "Juan",
		Departamento: "Ciencias",
		Cargo: "Profesor",
		Antiguedad: 5,
	}

	comp1 := model.Competencia{
		Tipo: "investigacion",
	}

	eval1 := model.Evaluacion{
		PuntajeBase: 80,
	}

	final1, reglas1 :=
		calculador.CalcularPuntajeConReglas(
			docente1,
			comp1,
			eval1,
		)

	fmt.Println(
		engine.GenerarReporteJustificado(
			80,
			final1,
			reglas1,
		),
	)

	fmt.Println("\nESCENARIO 2")

	docente2 := model.Docente{
		Cargo: "Director",
	}

	comp2 := model.Competencia{
		Tipo: "gestion",
	}

	eval2 := model.Evaluacion{
		PuntajeBase: 80,
	}

	final2, reglas2 :=
		calculador.CalcularPuntajeConReglas(
			docente2,
			comp2,
			eval2,
		)

	fmt.Println(
		engine.GenerarReporteJustificado(
			80,
			final2,
			reglas2,
		),
	)

	fmt.Println("\nESCENARIO 3")

	docente3 := model.Docente{
		Departamento: "Ciencias",
		Cargo: "Director",
	}

	comp3 := model.Competencia{
		Tipo: "gestion",
	}

	eval3 := model.Evaluacion{
		PuntajeBase: 80,
	}

	final3, reglas3 :=
		calculador.CalcularPuntajeConReglas(
			docente3,
			comp3,
			eval3,
		)

	fmt.Println(
		engine.GenerarReporteJustificado(
			80,
			final3,
			reglas3,
		),
	)

	fmt.Println("\nESCENARIO 4")

	docente4 := model.Docente{
		Departamento: "Ciencias",
		Cargo: "Profesor",
		Antiguedad: 15,
	}

	comp4 := model.Competencia{
		Tipo: "investigacion",
	}

	eval4 := model.Evaluacion{
		PuntajeBase: 80,
	}

	final4, reglas4 :=
		calculador.CalcularPuntajeConReglas(
			docente4,
			comp4,
			eval4,
		)

	fmt.Println(
		engine.GenerarReporteJustificado(
			80,
			final4,
			reglas4,
		),
	)
}