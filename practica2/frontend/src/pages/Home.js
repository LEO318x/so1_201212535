import React from "react";
import { useState, useEffect } from "react";

export default function Home() {
    const [isLoading, setIsLoading] = useState(true);
    const [data, setData] = useState([]);

    useEffect(() => {        
        fetch("http://localhost:5000/api/cpu")
        .then((response) => response.json())
        .then((data) => {
        setData(data.data[0]);
        setIsLoading(false);
      });
    }, []);


    return <div>
        <p>Total de procesos: {data.total}</p>
        <p>Procesos en ejecución: {data.running}</p>
        <p>Procesos suspendidos: {data.suspended}</p>
        <p>Procesos detenidos: {data.stopped}</p>
        <p>Procesos zombies: {data.zombie}</p>
    </div>;
  }