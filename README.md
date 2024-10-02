# Hetmo Technical Test

## Descripción

Este proyecto es un MVP desarrollado como parte de una prueba técnica para la empresa Hetmo. La aplicación está implementada en Golang utilizando el framework [Fiber](https://gofiber.io/). La arquitectura del proyecto sigue el enfoque **hexagonal**, lo que permite una alta escalabilidad y mantenibilidad en el tiempo.

### Arquitectura del Proyecto

El proyecto está organizado en base a la arquitectura hexagonal, separando la lógica por capas para garantizar una estructura desacoplada y fácilmente escalable. Las capas principales incluyen:

- **Domain**: Define las entidades centrales y sus reglas de negocio.
- **Infrastructure**: Contiene los controladores HTTP y las configuraciones de conexión con otras dependencias, como bases de datos.
- **Repository**: Implementación de la persistencia de datos, que interactúa con la base de datos PostgreSQL.

Además, se utiliza **vertical slicing** y **screaming architecture**, facilitando la organización del código y permitiendo un desarrollo más intuitivo. Estas estrategias permiten:

- **Vertical Slicing**: Agrupar cada nueva funcionalidad dentro de su propio módulo, manteniendo las capas relacionadas en un solo lugar.
- **Screaming Architecture**: Asegurar que la estructura del proyecto comunique claramente las intenciones del software y cada parte de la lógica esté relacionada a su correspondiente entidad.

Cada entidad dentro del proyecto sigue este esquema organizativo, lo que permite agregar nuevas funcionalidades sin afectar el código existente.

### Contenedores Docker

La aplicación está contenida en dos contenedores:

1. **Backend**: Usa una imagen basada en Golang.
2. **Base de datos**: Usa una imagen de PostgreSQL para almacenar la información de la aplicación.

Todas las variables sensibles como contraseñas y configuraciones de base de datos están gestionadas mediante variables de entorno.

## Requisitos Previos

- Docker y Docker Compose instalados en tu sistema.
- Golang 1.20+ para desarrollo local (opcional).

## Cómo Ejecutar el Proyecto

1. Clonar el repositorio:
   ```bash
   git clone https://github.com/FrancoRutigliano/hetmo_prueba_tecnica
   cd hetmo_prueba_tecnica
   ```

2. Setear variables de entorno **.env**:
     ```bash
     PORT=":8080"
    POSTGRES_USER="admin"
    POSTGRES_PASSWORD="1234"
    POSTGRES_DB="mydb"
    POSTGRES_HOST="container_db"
    ROUTE_MIGRATE="file://cmd/migrate/migrations"
    CONNECTION="postgres://admin:1234@container_db/mydb?sslmode=disable"
    SECRET_JWT="adminHetmo"
    JWT_EXP="604800"
     ```
3. Levantá contenedores y construí las imagenes:
```bash
   docker-compose up -d --build
   ```

## Colección de Postman
Puedes consultar y probar todas las peticiones disponibles en la colección de Postman a través del siguiente enlace:

[Enlace a la Colección de Postman](https://documenter.getpostman.com/view/29283926/2sAXxLAtmV)


