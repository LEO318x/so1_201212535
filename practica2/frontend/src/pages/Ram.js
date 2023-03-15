import React from "react";
import { useState, useEffect } from "react";
import "../styles.css";

import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);


export default function Ram() {
  const [isLoading, setIsLoading] = useState(true);
  const [conRam, setConRam] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setIsLoading(true);
      fetch(`${process.env.REACT_APP_URL}/api/ram`)
      .then((response) => response.json())
      .then((data) => {
        setConRam(data.data[0].porc_utilizado);
        console.log(data.data[0].porc_utilizado);
        setIsLoading(false);
      });
    }, 5000);
    return () => clearInterval(interval);
  }, []);
 
 const data = {
    labels: ['Consumido', 'Libre'],
    datasets: [
      {
        label: 'Uso de RAM',
        data: [conRam, 100-conRam],
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