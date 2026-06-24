package strategy

type EstrategiaPonderacion interface {
	Aplica(ctx Contexto) bool
	Calcular(ctx Contexto) float64
	Nombre() string
	Prioridad() int
}
