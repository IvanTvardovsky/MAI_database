import React from "react";
import { Link } from "react-router-dom";
import logo from "./../css/source/mai_db_logo.png";
import Image from "./Image";

class Navbar extends React.Component {
   render() {
      return (
         <div className="navbar">
            <div className="navbar-wrapper">
               <div className="navbar-logo">
                  <Link to="/">
                     <Image image={logo} />
                  </Link>
               </div>
               <nav role="navigation" className="navbar-menu">
                  <ul role="list" className="navbar-menu-list">
                     <li>
                        <Link to="/main" className="highlight-link">Университеты</Link>
                     </li>
                     <li>
                        <Link to="/admin" className="highlight-link">Админ панель</Link>
                     </li>
                  </ul>
               </nav>
            </div>
         </div>
      );
   }
}

export default Navbar;
