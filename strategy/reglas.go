package strategy

// R1: Ciencias + Investigación (+10%)
type ReglaCienciasInvestigacion struct{}

func (r ReglaCienciasInvestigacion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Departamento == "Ciencias" &&
		ctx.Competencia.Tipo == "investigacion"
}

func (r ReglaCienciasInvestigacion) Calcular(ctx Contexto) float64 {
	return ctx.PuntajeActual * 0.10
}

func (r ReglaCienciasInvestigacion) Nombre() string {
	return "R1 Bonificación Ciencias Investigación"
}

func (r ReglaCienciasInvestigacion) Prioridad() int {
	return 1
}

// R2: Director + Gestión (-10%)
type ReglaDirectorGestion struct{}

func (r ReglaDirectorGestion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Cargo == "Director" &&
		ctx.Competencia.Tipo == "gestion"
}

func (r ReglaDirectorGestion) Calcular(ctx Contexto) float64 {
	return ctx.PuntajeActual * -0.10
}

func (r ReglaDirectorGestion) Nombre() string {
	return "R2 Reducción Director Gestión"
}

func (r ReglaDirectorGestion) Prioridad() int {
	return 2
}

// R4: Coordinador + Académica (+5%)
type ReglaCoordinadorAcademica struct{}

func (r ReglaCoordinadorAcademica) Aplica(ctx Contexto) bool {
	return ctx.Docente.Cargo == "Coordinador" &&
		ctx.Competencia.Tipo == "academica"
}

func (r ReglaCoordinadorAcademica) Calcular(ctx Contexto) float64 {
	return ctx.PuntajeActual * 0.05
}

func (r ReglaCoordinadorAcademica) Nombre() string {
	return "R4 Bonificación Coordinador Académica"
}

func (r ReglaCoordinadorAcademica) Prioridad() int {
	return 3
}

// R5: Administración + Gestión (+15%)
type ReglaAdministracionGestion struct{}

func (r ReglaAdministracionGestion) Aplica(ctx Contexto) bool {
	return ctx.Docente.Departamento == "Administración" &&
		ctx.Competencia.Tipo == "gestion"
}

func (r ReglaAdministracionGestion) Calcular(ctx Contexto) float64 {
	return ctx.PuntajeActual * 0.15
}

func (r ReglaAdministracionGestion) Nombre() string {
	return "R5 Bonificación Administración Gestión"
}

func (r ReglaAdministracionGestion) Prioridad() int {
	return 4
}

// ==========================
// R3: Antigüedad +10 años Y R1 aplicado
// ==========================
type ReglaAntiguedad struct{}

func (r ReglaAntiguedad) Aplica(ctx Contexto) bool {
	return ctx.Docente.AniosAntiguedad > 10 &&
		ctx.ReglasAplicadas["R1 Bonificación Ciencias Investigación"]
}

func (r ReglaAntiguedad) Calcular(ctx Contexto) float64 {
	return ctx.PuntajeActual * 0.05
}

func (r ReglaAntiguedad) Nombre() string {
	return "R3 Bonificación Antigüedad"
}

func (r ReglaAntiguedad) Prioridad() int {
	return 5
}
