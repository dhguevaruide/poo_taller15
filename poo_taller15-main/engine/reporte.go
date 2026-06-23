package engine

import (
	"fmt"
	"poo_taller15/model"
)

func GenerarReporteJustificado(
	evaluacion model.Evaluacion,
	docente model.Docente,
	comp model.Competencia,
) string {

	puntaje, reglas :=
		CalcularPuntajeConReglas(
			evaluacion,
			docente,
			comp,
		)

	reporte := "===== REPORTE =====\n"

	reporte += fmt.Sprintf(
		"Docente: %s\n",
		docente.Nombre)

	reporte += fmt.Sprintf(
		"Puntaje Base: %.2f\n",
		evaluacion.PuntajeBase)

	reporte += "Reglas Aplicadas:\n"

	for _, r := range reglas {
		reporte += "- " + r + "\n"
	}

	reporte += fmt.Sprintf(
		"Puntaje Final: %.2f\n",
		puntaje)

	return reporte
}