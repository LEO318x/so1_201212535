import React from "react";
import { useState, useEffect } from "react";
import { Col, Container, Row, Table } from "react-bootstrap";

export default function Home() {
    const [isLoading, setIsLoading] = useState(true);
    const [datos, setData] = useState([]);

    useEffect(() => {        
        fetch(`${process.env.REACT_APP_URL}/api/cpu`)
        .then((response) => response.json())
        .then((data) => {
        setData(data.data[0]);
        setIsLoading(false);
      });
    }, []);

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
        <div>
          <p>Total de procesos: {datos.total}</p>
          <p>Procesos en ejecución: {datos.running}</p>
          <p>Procesos suspendidos: {datos.suspended}</p>
          <p>Procesos detenidos: {datos.stopped}</p>
          <p>Procesos zombies: {datos.zombie}</p>
        </div>
        </Col>
      </Row>
      <Row>
      <Table striped bordered hover>
      <thead>
        <tr>
          <th>Nombre</th>
          <th>UID</th>
          <th>PID</th>
          <th>Proceso</th>
          <th>Estado</th>
          <th>RAM MB</th>
          <th>Hijos</th>
        </tr>
      </thead>
      <tbody>
        
        {JSON.parse(datos.procesos).TablaProcesos.map(((key, index) =>
        <tr key={`${key.Nombre}${index}`}>
          <td>{key.Nombre}</td>
          <td>{key.UID}</td>
          <td>{key.PID}</td>
          <td>{key.Proceso}</td>
          <td>{key.Estado}</td>
          <td>{key.VmRSS}MB</td>
          <td>
            <Table striped bordered hover>
      <thead>
        <tr>
          <th>Nombre</th>
          <th>UID</th>
          <th>PID</th>
          <th>Proceso</th>
          <th>Estado</th>
          <th>RAM MB</th>
        </tr>
      </thead>
      <tbody>
      {
        key.Hijos.map(d => (
          <tr>
          <td>{d.Nombre}</td>
          <td>{d.UID}</td>
          <td>{d.PID}</td>
          <td>{d.Proceso}</td>
          <td>{d.Estado}</td>
          <td>{d.VmRSS}MB</td>
          </tr>
        ))
      }
      </tbody>
      </Table>
      </td>
        </tr>
      ))}
        
      </tbody>
    </Table>
      
      </Row>
    </Container>
    );
  }
