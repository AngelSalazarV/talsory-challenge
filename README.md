# Desafío Técnico – División TI

Sistema compuesto por un frontend y dos APIs que procesan matrices.

## Tecnologías

* **Frontend:** Next.js, React, TypeScript, Tailwind CSS
* **API principal:** Go + Fiber
* **API secundaria:** Node.js + Express + TypeScript
* **Seguridad:** JWT
* **Pruebas:** Go testing + Jest
* **Contenedores:** Docker + Docker Compose
* **Despliegue:** Vercel + Render

## Arquitectura

```text
Frontend
   ↓
Go API
   ↓ HTTP
Node API
```

La API Go recibe la matriz, valida el JWT, realiza la factorización QR y envía la matriz a Node.js.

La API Node.js calcula:

* Máximo
* Mínimo
* Promedio
* Suma total
* Si la matriz es diagonal

## Ejemplo

```json
{
  "matrix": [
    [1, 2],
    [3, 4],
    [5, 6]
  ]
}
```

La API Go devuelve las matrices `Q` y `R`, además de las estadísticas calculadas por Node.js.

La factorización utiliza el proceso de **Gram-Schmidt modificado**.

## Autenticación

El endpoint `/api/qr` está protegido mediante JWT.

Login:

```text
POST /auth/login
```

Usuario de prueba:

```text
admin / admin123
```

## Ejecución local

```bash
docker compose up --build
```

* Go API: `http://localhost:3000`
* Node API: `http://localhost:3001`

Para ejecutar las pruebas:

```bash
cd go-api
go test ./...
```

```bash
cd node-api
npm test
```

## Despliegue

**Frontend:** https://talsory-challenge.vercel.app

## Nota

El enunciado menciona inicialmente una rotación de matrices, pero posteriormente especifica la **factorización QR** como funcionalidad requerida. Se implementó QR como operación principal.
