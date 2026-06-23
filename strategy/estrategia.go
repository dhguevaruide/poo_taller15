package strategy

type EstrategiaPonderacion interface {
    Calcular(ctx Contexto) float64
    Aplica(ctx Contexto) bool
    Nombre() string
    Prioridad() int
}