import React from 'react';
import { Modal, Button, FormControl } from 'react-bootstrap';

const Form = (props) => {
  const { ID, name, familyName, givenName, email, admin } = props.user;
  const showModal = props.showModal;
  const hideModal = props.hideModal;

  return (
    <div>
      <Modal show={ props.modal } onHide={ hideModal }>
        <form className="form" onSubmit={ props.onSubmitHandler }>
          <Modal.Header closeButton>
            {/* <img src={ picture } alt={ name }/> */}
            <Modal.Title className="text-primary font-weight-bold">{ ID ? `${name}` : props.title }</Modal.Title>
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
                checked={ admin }
                onChange={ props.onChangeHandler }
                className="checkbox"
              />
            </div>
          </Modal.Body>
          
          <Modal.Footer>
            <Button variant="secondary" onClick={ hideModal }>Cancel</Button>
            <Button type="submit" variant="primary" onClick={ hideModal }>{ ID ? "Update" : "Add" }</Button>
          </Modal.Footer>
        </form>
      </Modal>
    </div>
  );
}

export default Form;