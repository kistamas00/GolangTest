import * as React from 'react';

const NavBar = (props: any) => (
  <nav className="navbar fixed-bottom navbar-dark">
    <a className="navbar-brand" href="#"/>
    <button type="button" className="btn btn-success" data-toggle="modal" data-target={'#'+props.target}>
      <i className="fas fa-plus"/>
    </button>
    <a className="navbar-brand" href="#"/>
  </nav>
);

export default NavBar;