import React from 'react';
import { NavLink } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faBars, faHome, faUsers, faServer, faSignOutAlt, faPhone } from '@fortawesome/free-solid-svg-icons';

const Navigation = (props) => {
  return (
    <div className = "bg-seconadry">
      <button type="button" className="mobile-nav-toggle d-xl-none">
        <FontAwesomeIcon icon={faBars} />
      </button>

      <header id="header">
        <div className="d-flex flex-column">
          <div className="profile">
          <img src={props.picture} className="img-fluid rounded-circle" />

            <div className="social-links mt-3 text-center">
              {/* <a href="#" className="twitter"><FontAwesomeIcon icon={faTwitter} /></a>
              <a href="#" className="facebook"><FontAwesomeIcon icon={faFacebook} /></a>
              <a href="#" className="instagram"><FontAwesomeIcon icon={faInstagram} /></a>
              <a href="#" className="google-plus"><FontAwesomeIcon icon={faGooglePlus} /></a>
              <a href="#" className="linkedin"><FontAwesomeIcon icon={faLinkedin} /></a> */}
            </div>
          </div>

          <nav className="nav-menu">
            <ul>
              <li className="nav-item">
                <NavLink exact to = "/" className="nav-link">
                  <FontAwesomeIcon icon={faHome} /> &nbsp;Home
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink exact to = "/users" className="nav-link">
                  <FontAwesomeIcon icon={faUsers} /> &nbsp;Users
                </NavLink>
              </li>
              <li className="nav-item">
                <a href="#services">
                <FontAwesomeIcon icon={faServer} /> &nbsp;Services
                </a>
              </li>
              <li className="nav-item">
                <a href="#contact">
                  <FontAwesomeIcon icon={faPhone} /> &nbsp;Contact
                </a>
              </li>
              <li className="nav-item">
                <a href="/logout">
                  <FontAwesomeIcon icon={faSignOutAlt} /> &nbsp;Log Out
                </a>
              </li>
            </ul>
          </nav>

          <button type="button" className="mobile-nav-toggle d-xl-none"><i className="icofont-navigation-menu"></i></button>
        </div>
      </header>
    </div>
  );
}

export default Navigation;