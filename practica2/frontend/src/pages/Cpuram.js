import React from "react";
import { useState, useEffect } from "react";
import "../styles.css";

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js';
import { Doughnut } from 'react-chartjs-2';

ChartJS.register(ArcElement, Tooltip, Legend);

export default function Cpuram() {
  const [isLoading, setIsLoading] = useState(true);
  const [conCpu, setConCpu] = useState(0);
  const [conRam, setConRam] = useState(0);

  useEffect(() => {
    const interval = setInterval(() => {
      setIsLoading(true);
      fetch(`${process.env.REACT_APP_URL}/api/cpuram`)
      .then((response) => response.json())
      .then((data) => {
        setConCpu(data.data[0].porc_uso);
        setConRam(data.data2[0].porc_utilizado)
        setIsLoading(false);
      });
    }, 10000);
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

  const data2 = {
    labels: ['Uso', 'Libre'],
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

  const config2 = {
    type: 'doughnut',
    data: data2,
  };

  if (isLoading) {
    return (
      <div className="App">
        <h1>Cargando...</h1>
      </div>
    );
  }

  return (
    <Container>
        <Row>
            <Col>
            <div>Uso de CPU</div>
                <div style={{height: "400px", width: "400px"}}>
                    <Doughnut
                    options={{responsive:true, maintainAspectRatio: false }}
                    data={data}
                    />
                </div> 
                <div>
                <p>Uso de CPU: {conCpu}%</p>  
                </div>               
            </Col>
            <Col>
            <div>Uso de RAM</div>
                <div style={{height: "400px", width: "400px"}}>
                    <Doughnut
                    options={{responsive:true, maintainAspectRatio: false }}
                    data={data2}
                    />
                </div> 
                <div>
                <p>Uso de RAM: {conRam}%</p>
                <p>Ram Disponible: {100-conRam}%</p>
                </div>                
            </Col>
        </Row>
        <Row>
            <div></div>
        </Row>
    </Container>
    
  );
}
