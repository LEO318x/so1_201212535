import React,{useState,useEffect} from 'react';


function Logs(){
    const [data,setData]=useState([]);
    const getData=()=>{
      fetch('http://localhost:8000/historial'
      ,{
        headers : { 
          'Content-Type': 'application/json',
          'Accept': 'application/json'
         }
      }
      )
        .then(function(response){
          console.log(response)
          return response.json();
        })
        .then(function(myJson) {
          console.log(myJson);
          setData(myJson)
        });
    }
    useEffect(()=>{
      getData()
    },[])
    return (
      <div className="Logs">
        <table align="center" className='tabla'>
	<thead>
		<tr>
			<th>ID</th>
			<th>Numero 1</th>
			<th>Numero 2</th>
			<th>Operador</th>
			<th>Resultado</th>
            <th>Fecha y Hora</th>
		</tr>
	</thead>
	<tbody>
		
       {
         data && data.length>0 && data.map((item)=><tr><td>{item.id}</td><td>{item.num1}</td><td>{item.num2}</td><td>{item.operador}</td><td>{item.resultado}</td><td>{item.fecha_hora}</td></tr>)
       }
	</tbody>
</table>
      </div>
    );
}    


export default Logs;