# GORM-PRACTICAS
 Repositorio de Golang para trabajar con funcionalidades de GORM

 * Conexion a base de datos
 * Creacion de modelos
 * Creacion de migraciones
 * Creacion de modelos con relacion
---
---
 # Rutas Personas

 ## /Personas
  ### Tipo de metodo: GET

  #### Esta ruta retorna todas las personas que se encuentran en la base de datos
  ### Respuesta de ejemplo
  ```
  [
    {
        "ID": 1,
        "CreatedAt": "2024-11-06T17:24:56.084603-03:00",
        "UpdatedAt": "2024-11-06T19:55:54.110512-03:00",
        "DeletedAt": null,
        "nombre": "Manuxxrsb",
        "apellido": "Back Mas Vio",
        "Objetos": null
    },
    {
        "ID": 2,
        "CreatedAt": "2024-11-06T18:31:02.072508-03:00",
        "UpdatedAt": "2024-11-06T18:31:02.072508-03:00",
        "DeletedAt": null,
        "nombre": "Miguel",
        "apellido": "Solis",
        "Objetos": null
    }
  ]
  ```
---
## GET /Persona/:id
 #### Esta ruta retorna una Persona por id y sus relaciones
  ```
  http://localhost:8080/Persona/:id 
  ```
  ### Respuesta de ejemplo
  ### HTTP RESPONSE ( 200 OK )
  ```
  {
    "ID": 1,
    "CreatedAt": "2024-11-06T17:24:56.084603-03:00",
    "UpdatedAt": "2024-11-06T19:55:54.110512-03:00",
    "DeletedAt": null,
    "nombre": "Manuxxrsb",
    "apellido": "Back Mas Vio",
    "Objetos": [
        {
            "ID": 2,
            "CreatedAt": "2024-11-06T19:18:37.843244-03:00",
            "UpdatedAt": "2024-11-06T20:01:54.921498-03:00",
            "DeletedAt": null,
            "nombre": "Cepillo de dientes",
            "persona_id": 1
        },
        {
            "ID": 3,
            "CreatedAt": "2024-11-06T19:19:10.060903-03:00",
            "UpdatedAt": "2024-11-06T19:19:10.060903-03:00",
            "DeletedAt": null,
            "nombre": "Escritorio",
            "persona_id": 1
        }
    ]
  }
  ```
---
 

 ## POST /Persona 
 #### Esta ruta crea un nuevo Persona, para crear un nuevo Persona debes ingresar un JSON con la siguiente estructura
 ```
  {
    "nombre":"Miguel",
    "apellido":"Solis"
  }
  ```
  ### Respuesta de ejemplo
  ### HTTP RESPONSE ( 200 OK )
  ```
  {
    "ID": 2,
    "CreatedAt": "2024-11-06T18:31:02.0725084-03:00",
    "UpdatedAt": "2024-11-06T18:31:02.0725084-03:00",
    "DeletedAt": null,
    "Nombre": "Miguel",
    "Apellido": "Solis",
    "Objetos": null
}
  ```

 - DELETE /Persona/:id: elimina un Persona
  ```
  ```

 # Rutas Objeto

 - POST /Objeto
 ### JSON de ejemplo Objeto con  relacion
  ```
    {
    "nombre":"Computadora",
    "persona_id":"1"
    }
  ```
  - POST /Objeto
 ### JSON de ejemplo Objeto sin  relacion
  ```
    {
    "nombre":"Computadora",
    }
  ```

