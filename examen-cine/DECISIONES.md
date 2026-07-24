# DECISIONES (CP1) — máximo 10 líneas

Responda: ¿qué entidades identificó en las pantallas y por qué se relacionan
así? Sea concreto: qué vio en cada pantalla que lo llevó a cada campo y a
cada clave foránea.

1. **Función** (Entidad A): Pantalla 01 muestra catálogo con nombre, precio, stock y estado. Cada función tiene múltiples compras.
2. **Cliente**: Pantalla 02 selector de cliente. Pantalla 03 muestra nombre + cédula. Un cliente hace múltiples compras.
3. **Compra**: Pantalla 02 formulario crea compra. Pantalla 03 lista con cliente, función, cantidad, total, estado. Compra relaciona Función (FK) + Cliente (FK) + campos de cantidad, estado y total.

