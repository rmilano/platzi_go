# platzi_go

Repositorio de prácticas del curso básico de Go (Platzi).

## Estructura

- `HolaMundo/main.go`: ejemplo inicial "Hola Mundo".
- `Makefile`: comandos convenientes para build, run, test y clean.
- `bin/`: carpeta donde se generan los binarios (ignorados por Git).
- `go.mod`: definición del módulo Go `github.com/rmilano/platzi_go`.

## Requisitos

- Go 1.20+ (recomendado 1.21)
- `make` (opcional para usar el Makefile)

## Uso rápido

### Ejecutar sin compilar
```bash
go run ./HolaMundo
```

### Compilar binario
```bash
mkdir -p bin
go build -o bin/HolaMundo ./HolaMundo
./bin/HolaMundo
```

### Usando Makefile
```bash
make build    # compila a bin/HolaMundo
make run      # ejecuta directamente
make fmt      # formatea código
make vet      # analiza posibles problemas
make test     # ejecuta tests (cuando existan)
make clean    # limpia la carpeta bin
```

## Flujo Git básico
```bash
git add -A
git commit -m "mensaje descriptivo"
git push
```

## Próximos pasos sugeridos
- Añadir pruebas unitarias (`go test`).
- Crear más ejemplos (paquetes, variables, valores, condiciones, bucles, etc.).

## Licencia
Pendiente de definir. Puedes añadir una licencia (MIT recomendada para proyectos de ejemplo).
