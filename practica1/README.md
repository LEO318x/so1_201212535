# Documentación
#### Mike Leonel Molina García 201212535

## Contenedor Frontend

Contenedor que aloja una aplicación de una calculadora simple, realizada en el framework de react.js

## Contenedor Backend

Contenedor que aloja una aplicación desarrollada en go, la cual provee la api que es utilizada en el contendor frontend para la realización de las operaciones básicas de la calculadora.

## Contenedor MySQL
Contenedor que aloja todos los registros de las operaciones realizadas desde el frontend.

## Contenedor Bash
Contenedor que aloja un bash script la cual se encarga de analizar las operaciones que fueron realizadas a traves del frontend, la cual nos indica el total de sumas, restas, divisiones, multiplicaciones, errores en división entre otros, para el funcionamiento del script se utiliza la libreria curl y jq para su respectivo analisis, el bash se conecta a un api que es consumida y devuelve un resultado para su respectivo analisis.

## Entidad Relación BD
| Key | Log  |
| --- | --- |
|PK  |ID int  |
|  | numero1 int  |
|  | numero2 int |
|  | operador varchar(1)  |
|  | resultado int  |
|  | fecha_hora datetime  |
