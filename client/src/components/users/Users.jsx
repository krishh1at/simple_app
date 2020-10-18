import React from 'react';
import Axios from 'axios';
import UserRow from './UserRow';
import Form from './Form';

class Users extends React.Component {
  constructor(params) {
    super(params)

    this.state = {
      user: {},
      users: [],
      userCount: 0
    }

    this.onChangeHandler = this.onChangeHandler.bind(this);
    this.onSubmitHandler = this.onSubmitHandler.bind(this);
  }

  componentDidMount() {
    Axios.get(`/api/users`)
    .then(resp => {
      this.setState({ users: resp.data, userCount: resp.data.length });
    })
    .catch(resp => {
      console.log(resp);
    });
  }

  onChangeHandler = (e) => {
    var newUser = Object.assign(
                    {},
                    this.state.user,
                    {[e.target.name]: (e.target.type == "checkbox" ? e.target.checked : e.target.value)
                  });
       
    this.setState({ user: newUser });
  }

  onSubmitHandler = (e) => {
    e.preventDefault();

    Axios.post(`/api/users`, this.state.user)
    .then(resp => {
      const users = this.state.users;
      const userCount = users.unshift(resp.data);
      this.setState({ users: users, userCount: userCount, user: {} });
    })
    .catch(resp => {
      console.log(resp);
    })
  }

  render() {
    const users = this.state.users.map(user => {
      return <UserRow user={user} id={ user.ID } key={ user.ID } />
    });

    return (
      <div>
        <div className="section-title clearfix">
          <h2 className="h2 float-left">Users: { this.state.userCount }</h2>
          
          <Form
            onChangeHandler={ this.onChangeHandler }
            onSubmitHandler={ this.onSubmitHandler }
            title="Add New User"
            user={ this.state.user }
          />
        </div>
        <div className="table-responsive">
          <table className="table table-hover">
            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Email</th>
                <th className="text-center">Admin</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody>
              { users }
            </tbody>
          </table>
        </div>
      </div>
    );
  }
}

export default Users;