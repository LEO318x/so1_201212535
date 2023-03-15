import React from "react";
import { useState, useEffect } from "react";
import "../styles.css";

import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);

export default function Cpu() {
  const [isLoading, setIsLoading] = useState(true);
  const [conCpu, setConCpu] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setIsLoading(true);
      fetch("http://localhost:5000/api/cpu")
      .then((response) => response.json())
      .then((data) => {
        setConCpu(data.data[0].porc_uso);
        console.log(data.data[0].porc_uso);
        setIsLoading(false);
      });
    }, 5000);
    return () => clearInterval(interval);
  }, []);
 
 const data = {
    labels: ['Uso', 'Libre'],
    datasets: [
      {
        label: 'Uso de CPU',
        data: [conCpu, 100-conCpu],
        backgroundColor: [
          'rgba(255, 99, 132, 0.2)',
          'rgba(54, 162, 235, 0.2)',
        ],
        borderColor: [
          'rgba(255, 99, 132, 1)',
          'rgba(54, 162, 235, 1)',
        ],
        borderWidth: 1,
      },
    ],
  };

  const config = {
    type: 'doughnut',
    data: data,
  };

  if (isLoading) {
    return (
      <div className="App">
        <h1>Cargando...</h1>
      </div>
    );
  }

  return (
    <div style={{height:"60vh", with: "80vw",position:"relative", marginBottom:"1%", padding:"1%"}}>
       <Doughnut
       height="200px"
       width="200px"
       options={{ maintainAspectRatio: false }}
        data={data}
       />
    </div>
  );
}