package engine

import (
	"fmt"
	"poo_taller15/model"
)

func GenerarReporteJustificado(e model.Evaluacion, reglas []string) string {
	reporte := "=== REPORTE DE EVALUACIÓN ===\n\n"
	reporte += fmt.Sprintf("Evaluación: %s\n", e.ID)
	reporte += fmt.Sprintf("Puntaje Base: %.2f\n", e.PuntajeBase)
	reporte += fmt.Sprintf("Puntaje Final: %.2f\n\n", e.PuntajeFinal)

	reporte += "Reglas aplicadas:\n"
	if len(reglas) == 0 {
		reporte += "  Ninguna\n"
	} else {
		for _, r := range reglas {
			reporte += fmt.Sprintf("  - %s\n", r)
		}
	}

	return reporte
}
