# Documentación
#### Mike Leonel Molina García 201212535

### Inserción de Modulos
`$ sudo insmod <modulo>.ko`

### Eliminación de Modulos
`$ sudo rmmod <modulo>.ko`

## Contenedor Frontend

Contenedor que aloja el dashboard, donde podemos visualizar los procesos e hijos, así como el consumo del procesador y RAM.

## Contenedor API
Contenedor que hace de puente entre la base de datos MySQL y la parte del frontend

## Contenedor Backend

Contenedor que aloja una aplicación desarrollada en go la cual es la encargada de almacenar la información de los procesos, uso de cpu y ram en una base de datos MySQL

## Contenedor MySQL
Contenedor que aloja la información de los procesos, uso de cpu y ram.

## Entidad Relacion BD

#### CPU
| Key | cpu  |
| --- | --- |
|PK  |id_cpu int  |
|  | porc_uso varchar(100)  |
|  | running varchar(100) |
|  | suspended varchar(100)  |
|  | stopped varchar(100)  |
|  | zombie varchar(100)  |
|  | procesos longtext  |

#### RAM
| Key | ram  |
| --- | --- |
|PK  |id_ram int  |
|  | total varchar(100)  |
|  | consumida varchar(100)  |
|  | porc_utilizado varchar(100)  |
