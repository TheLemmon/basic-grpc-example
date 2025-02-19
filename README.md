# basic-grpc-example

## Description

Este repositorio contiene un ejemplo sencillo que integra dos servicios:

- Un servidor gRPC implementado en Golang, el cual consulta de manera asíncrona una API externa de chistes ("jokes") y devuelve 25 chistes de forma aleatoria. Es importante destacar que cada chiste es único, es decir, no se repiten en cada respuesta.

- Un servidor HTTP desarrollado en Python utilizando FastAPI. Este servidor actúa como cliente del servidor gRPC de chistes y expone el siguiente endpoint:

## Instalacion

### Cliente Grpc
1. Instalar Python en su version 3.12
2. Crear un entorno virtual y activarlo
```
cd client
python -m venv venv
source venv/bin/activate
```
3. Instalar las dependencias del proyecto
```
pip install -r requirements.txt
```
4. Ejecutar el servidor HTTP
```
fastapi run main.py
```

### Servidor Grpc
1. Instalar Go en su version 1.23.5
2. Instalar dependencias del proyecto
```
cd server
go mod tidy
```
3. Ejecutar el servidor gRPC
```
go run main.go
```

### Endpoints disponibles
- **GET localhost:8000/jokes**
    - Al consumir este endpoint se realiza una petición al servidor gRPC para obtener un listado de chistes. El response es un objeto JSON que contiene:

```json
{
  "result": [
    {
      "id": "1",
      "url": "https://example.com/chiste/1",
      "value": "Aquí va un chiste."
    },
    {
      "id": "2",
      "url": "https://example.com/chiste/2",
      "value": "Aquí va otro chiste."
    }
  ],
  "total": 25
}
```

Esta arquitectura permite separar la lógica de negocio (servidor gRPC en Golang) de la capa de exposición y consumo (servidor HTTP en Python), facilitando la escalabilidad y el mantenimiento del sistema.
