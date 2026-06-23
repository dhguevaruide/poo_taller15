package engine

import (
	"fmt"
	"strings"
)

func GenerarReporteJustificado(
	puntajeBase float64,
	puntajeFinal float64,
	reglas []string,
) string {

	var sb strings.Builder

	sb.WriteString("===== REPORTE =====\n")
	sb.WriteString(
		fmt.Sprintf(
			"Puntaje Base: %.2f\n",
			puntajeBase,
		),
	)

	sb.WriteString(
		fmt.Sprintf(
			"Puntaje Final: %.2f\n",
			puntajeFinal,
		),
	)

	sb.WriteString("\nReglas Aplicadas:\n")

	for _, r := range reglas {
		sb.WriteString("- " + r + "\n")
	}

	return sb.String()
}