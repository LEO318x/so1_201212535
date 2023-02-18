#!/bin/bash

content=$(curl -L http://192.168.3.7:8000/historial)
content2=$(curl -L http://192.168.3.7:8000/historialhoy)

operaciones=$(echo $content | jq '.. | .operador? //empty')
operacioneserror=$(echo $content | jq '.. | .resultado? //empty')
cantsuma=$(echo $operaciones | grep -o "+" | wc -l)
cantresta=$(echo $operaciones | grep -o "-" | wc -l)
cantmulti=$(echo $operaciones | grep -o "*" | wc -l)
cantdiv=$(echo $operaciones | grep -o "/" | wc -l)
cantdiverr=$(echo $operacioneserror | grep -oP -- '-[1]' | wc -l)
totaloperaciones=$(($cantsuma+$cantresta+$cantmulti+$cantdiv))

echo "###################################################"
echo "# Operaciones realizadas: " $totaloperaciones
echo "###################################################"
echo "# Sumas totales realizadas: " $cantsuma
echo "# Restas totales realizadas: " $cantresta
echo "# Multiplicaciones totales realizadas: " $cantmulti
echo "# Divisiones totales realizadas: " $cantdiv
echo "###################################################"
echo "# Divisiones con error: " $cantdiverr
echo "###################################################"
echo "#Operaciones del día"
echo "###################################################"
echo $content2
echo "###################################################"
echo "#Historial completo operaciones"
echo "###################################################"
echo $content



