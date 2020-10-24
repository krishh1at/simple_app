import React from 'react';
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckCircle, faTimesCircle, faEdit, faTrash } from '@fortawesome/free-solid-svg-icons';

const UserRow = (props) => {
  const { ID, name, email, admin } = props.user;

  return(
    <tr>
      <td className="font-weight-bold">
        <Link to={`users/${ID}`} >{ ID }</Link>
      </td>
      <td>{ name }</td>
      <td>{ email }</td>
      <td className="text-center">
        { admin ? 
            <FontAwesomeIcon icon={ faCheckCircle } className="text-success" /> :
            <FontAwesomeIcon icon={ faTimesCircle } className="text-danger" />
        }
      </td>
      <td>
        <Link to="#" onClick={ () => props.showModal(ID) } className="mx-1" >
          <FontAwesomeIcon icon={ faEdit } />
        </Link>

        <Link to="#" onClick={ () => props.deleteUser(ID) } className="mx-1 text-danger">
          <FontAwesomeIcon icon={ faTrash } />
        </Link>
      </td>
    </tr>
  );
}

export default UserRow;