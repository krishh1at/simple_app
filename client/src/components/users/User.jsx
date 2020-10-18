import React from 'react';
import Axios from 'axios';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCheckCircle, faTimesCircle } from '@fortawesome/free-solid-svg-icons';

class User extends React.Component {
  constructor(props) {
    super(props);
    this.state = { user: {} };
  }

  componentDidMount() {
    Axios.get(`/api/users/${this.props.match.params.id}`)
    .then(resp => {
      this.setState({ user: resp.data })
    })
  }

  render() {
    const { name, Email, email_verified, admin, picture } = this.state.user;
    
    return(
      <div>
        <div className="section-title clearfix">
          <img src={ picture } alt={ name } className="img-fluid rounded-circle" />
          <h2 className="h2 mt-2">{ name }</h2>
        </div>

        <div className="section-body">
          <div className="email">
            <label className="text-dark font-weight-bold">Email:</label>
            <label className="mx-2">
              { Email }
              {
                email_verified ?
                <FontAwesomeIcon icon={ faCheckCircle } className="text-success" /> :
                <FontAwesomeIcon icon={ faTimesCircle } className="text-danger" />
              }
            </label>
          </div>

          <div className="admin">
            <label className="text-dark font-weight-bold">Admin:</label>
            <label className="mx-2">
              {
                admin ?
                <FontAwesomeIcon icon={ faCheckCircle } className="text-success" /> :
                <FontAwesomeIcon icon={ faTimesCircle } className="text-danger" />
              }
            </label>
          </div>
        </div>
      </div>
    );
  }
}

export default User;