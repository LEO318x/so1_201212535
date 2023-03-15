import React from "react";
import "../styles.css";
import { Link } from "react-router-dom";

export default function Nav(){

  return(
        <div className="navbar">
          <div className="logo">SO1</div>
           <ul className="nav-links">
              <Link to="/">Home</Link>
              <Link to="/ram">Ram</Link>
              <Link to="/cpu">Cpu</Link>
           </ul>
        </div>
  );

}