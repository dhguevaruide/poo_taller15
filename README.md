Sistema de Evaluación de Competencias con Strategy

Justificación de Diseño

1. ¿Por qué elegiste el patrón Strategy en lugar de Chain of Responsibility o Visitor?
Elegí el patrón Strategy porque cada regla es un cálculo diferente y es más fácil agregar o cambiar reglas en el futuro. Chain of Responsibility me pareció más complicado porque las reglas no necesariamente se tienen que pasar de una a otra. Visitor no tenía mucho sentido para este caso.

2. ¿Cómo manejaste el caso donde dos reglas tienen la misma prioridad?
Le asigné a cada regla un número de prioridad diferente (del 1 al 5). Antes de aplicarlas las ordeno por prioridad para que no haya conflicto.

3. ¿Por qué usaste composición en lugar de agregación para ciertas relaciones?
Usé composición en el Contexto con el Docente, Competencia y Evaluacion porque el contexto necesita esos datos para funcionar. En cambio, el Calculador tiene una agregación con las estrategias porque las estrategias pueden existir por separado.

4. ¿Cómo extenderías el sistema para soportar reglas que dependen del historial del docente?
Podría agregar al Contexto una lista con las evaluaciones anteriores del docente. Después las reglas podrían revisar ese historial para decidir si se aplican o no.

5. ¿Qué impacto tiene tu diseño en la performance si hay 10,000 evaluaciones con 50 reglas cada una?
Creo que no habría mucho problema. Go es rápido y 50 reglas no son tantas. De todas formas, se podría mejorar guardando las estrategias ordenadas una sola vez al inicio.

---