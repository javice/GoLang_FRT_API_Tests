# JSONPlaceholder API Testing Project

Este proyecto contiene pruebas automatizadas para la API de JSONPlaceholder utilizando Go y siguiendo el patrón Page Object Model (POM).

## Estructura del Proyecto

```
jsonplaceholder-api-tests/
├── cmd/
│   └── run_tests/
│       └── main.go
├── pkg/
│   ├── api/
│   │   ├── client.go
│   │   └── endpoints.go
│   ├── models/
│   │   ├── post.go
│   │   ├── comment.go
│   │   ├── album.go
│   │   ├── photo.go
│   │   ├── todo.go
│   │   └── user.go
│   └── utils/
│       └── http_helper.go
├── tests/
│   └── api/
│       ├── posts_test.go
│       ├── comments_test.go
│       ├── albums_test.go
│       ├── photos_test.go
│       ├── todos_test.go
│       └── users_test.go
├── go.mod
├── Makefile
└── README.md
```

## Requisitos Previos

- Go 1.21 o superior
- Make (opcional, para usar el Makefile)

## Instalación

1. Clona el repositorio:

```bash
git clone https://github.com/your-username/jsonplaceholder-api-tests.git
cd jsonplaceholder-api-tests
```

2. Instala las dependencias:

```bash
make deps
# o
go mod download
```

## Ejecución de Tests

### Usando Make

```bash
# Ejecutar todos los tests
make test

# Ejecutar tests con cobertura
make test-coverage

# Limpiar archivos generados
make clean

# Ver todos los comandos disponibles
make help
```

### Usando Go directamente

```bash
# Ejecutar todos los tests
go test ./tests/api/...

# Ejecutar tests con cobertura
go test -coverprofile=coverage.out ./tests/api/...
go tool cover -html=coverage.out -o coverage.html
```

## Endpoints Testeados

La API de JSONPlaceholder proporciona los siguientes endpoints:

- `/posts` - Gestión de posts
- `/comments` - Gestión de comentarios
- `/albums` - Gestión de álbumes
- `/photos` - Gestión de fotos
- `/todos` - Gestión de tareas
- `/users` - Gestión de usuarios

### Métodos HTTP Soportados

- GET `/posts` - Obtener todos los posts
- GET `/posts/1` - Obtener un post específico
- GET `/posts/1/comments` - Obtener comentarios de un post
- GET `/comments?postId=1` - Obtener comentarios filtrados
- POST `/posts` - Crear un nuevo post
- PUT `/posts/1` - Actualizar un post completo
- PATCH `/posts/1` - Actualizar parcialmente un post
- DELETE `/posts/1` - Eliminar un post

## Estructura de los Tests

Los tests están organizados siguiendo el patrón Page Object Model:

- `pkg/api/client.go` - Cliente base para las peticiones HTTP
- `pkg/models/*.go` - Modelos de datos para cada recurso
- `tests/api/*_test.go` - Tests para cada endpoint

## Contribución

1. Fork el repositorio
2. Crea una rama para tu feature (`git checkout -b feature/amazing-feature`)
3. Commit tus cambios (`git commit -m 'Add some amazing feature'`)
4. Push a la rama (`git push origin feature/amazing-feature`)
5. Abre un Pull Request

## Licencia

Este proyecto está bajo la Licencia MIT - ver el archivo [LICENSE](LICENSE) para más detalles.
