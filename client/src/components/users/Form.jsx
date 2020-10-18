import React, { useState, useEffect } from 'react';
import { Modal, Button, FormControl } from 'react-bootstrap';

const Form = (props) => {
  const [show, setShow] = useState(false)
  const handleClose = () => setShow(false)
  const handleShow = () => setShow(true)
  const { name, familyName, givenName, email, admin } = props.user;

  return (
    <div>
      <Button variant="primary" onClick={ handleShow } className="float-right">Add New User</Button>

      <Modal show={ show } onHide={ handleClose }>
        <form className="form" onSubmit={ props.onSubmitHandler }>
          <Modal.Header closeButton>
            {/* <img src={ picture } alt={ name }/> */}
            <Modal.Title>{ props.title ? props.title : `Edit ${name}` }</Modal.Title>
          </Modal.Header>
          
          <Modal.Body>
            <div className="form-group">
              <label htmlFor="givenName">Last Name</label>
              <FormControl
                name="givenName"
                id="givenName"
                value={ givenName }
                onChange={ props.onChangeHandler }
                placeholder="Enter Last name"
              />
            </div>

            <div className="form-group">
              <label htmlFor="familyName">First Name</label>
              <FormControl
                name="familyName"
                id="familyName"
                value={ familyName }
                onChange={ props.onChangeHandler }
                placeholder="Enter First name"
              />
            </div>

            <div className="form-group">
              <label htmlFor="email">Email</label>
              <FormControl
                name="email"
                id="email"
                value={ email }
                onChange={ props.onChangeHandler }
                placeholder="Enter Email"
              />
            </div>

            <div className="form-group">
              <label htmlFor="admin" className="mr-1">Admin</label>
              <input
                type="checkbox"
                name="admin"
                id="admin"
                value={ admin }
                onChange={ props.onChangeHandler }
                className="checkbox"
              />
            </div>
          </Modal.Body>
          
          <Modal.Footer>
            <Button variant="secondary" onClick={ handleClose }>Cancel</Button>
            <Button type="submit" variant="primary" onClick={ handleClose }>{ props.title ? "Add" : "Update" }</Button>
          </Modal.Footer>
        </form>
      </Modal>
    </div>
  );
}

export default Form;