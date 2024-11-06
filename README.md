# GORM-PRACTICAS
 Repositorio de Golang para trabajar con funcionalidades de GORM

 * Conexion a base de datos
 * Creacion de modelos
 * Creacion de migraciones
 * Creacion de modelos con relacion

 # Rutas Personas

 - GET /users: devuelve todos los usuarios
  ## JSON de ejemplo
  ```
    {
    "nombre" : "nombre persona",
    "apellido" :  "apellido persona",

    }
  ```
 - GET /users/:id: devuelve un usuario por id
  ```
  ```
 - POST /users: crea un nuevo usuario
 ```
  ```
 - DELETE /users/:id: elimina un usuario
  ```
  ```

 # Rutas Objeto

 - POST /Objeto
 ## JSON de ejemplo Objeto con  relacion
  ```
    {
    "nombre":"Computadora",
    "persona_id":"1"
    }
  ```
  - POST /Objeto
 ## JSON de ejemplo Objeto sin  relacion
  ```
    {
    "nombre":"Computadora",
    }
  ```

