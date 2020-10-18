import React from 'react';
import { Link } from "react-router-dom";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckCircle, faTimesCircle } from '@fortawesome/free-solid-svg-icons';

const UserRow = (props) => {
  const { ID, name, Email, admin } = props.user;

  return(
    <tr>
      <td className="font-weight-bold">
        <Link to={`users/${ID}`} >{ ID }</Link>
      </td>
      <td>{ name }</td>
      <td>{ Email }</td>
      <td className="text-center">
        { admin ? 
            <FontAwesomeIcon icon={ faCheckCircle } className="text-success" /> :
            <FontAwesomeIcon icon={ faTimesCircle } className="text-danger" />
        }
      </td>
      <td></td>
    </tr>
  );
}

export default UserRow;