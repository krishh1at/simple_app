import React from 'react';
import Axios from 'axios';
import { Button } from 'react-bootstrap';
import UserRow from './UserRow';
import Form from './Form';

class Users extends React.Component {
  constructor(params) {
    super(params)

    this.state = {
      user: {},
      users: [],
      userCount: 0,
      modal: false
    }

    this.onChangeHandler = this.onChangeHandler.bind(this);
    this.onSubmitHandler = this.onSubmitHandler.bind(this);
    this.showModal = this.showModal.bind(this);
    this.hideModal = this.hideModal.bind(this);
    this.deleteUser = this.deleteUser.bind(this);
  }

  showModal(id) {
    if(id) {
      Axios.get(`/api/users/${id}`)
      .then(resp => {
        this.setState({ user: resp.data, modal: true });
      })
      .catch(resp => {
        console.log(resp);
      });
    } else {
      this.setState({ user: {}, modal: true });
    }
  }

  hideModal() {
    this.setState({ modal: false });
  }

  componentDidMount() {
    Axios.get(`/api/users`)
    .then(resp => {
      console.log(resp.data);
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
    const user = this.state.user;

    if(user.ID == null) {
      Axios.post(`/api/users`, this.state.user)
      .then(resp => {
        const users = this.state.users;
        const userCount = users.unshift(resp.data);
        this.setState({ users: users, userCount: userCount, user: {} });
      })
      .catch(resp => {
        console.log(resp);
      });
    } else {
      Axios.put(`/api/users/${user.ID}`, this.state.user)
      .then(resp => {
        const users = this.state.users.map(function(user){
          if(user.ID == resp.data.ID) {
            user = resp.data;
          }

          return user
        });

        this.setState({ users: users, user: {} });
      })
      .catch(resp => {
        console.log(resp);
      });
    }
  }

  deleteUser(id) {
    const confirmation = window.confirm("Are you sure?");
    
    if(confirmation) {
      Axios.delete(`/api/users/${id}`)
      .then(resp => {
        const users = this.state.users.filter((user) => user.ID != id);
        this.setState({ users: users });
      })
      .catch(resp => {
        console.log(resp.data);
      });
    }
  }

  render() {
    const users = this.state.users.map(user => {
      return <UserRow
               showModal={ this.showModal }
               getUserData={ this.getUserData }
               deleteUser={ this.deleteUser }
               user={user}
               id={ user.ID }
               key={ user.ID }
             />
    });

    return (
      <div>
        <div className="section-title clearfix">
          <h2 className="h2 float-left">Users: { this.state.userCount }</h2>
          <Button variant="primary" onClick={ () => this.showModal() } className="float-right">Add New User</Button>
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

        <Form
            onChangeHandler={ this.onChangeHandler }
            onSubmitHandler={ this.onSubmitHandler }
            showModal={ this.showModal }
            hideModal={ this.hideModal }
            modal={ this.state.modal }
            title="Add New User"
            user={ this.state.user }
          />
      </div>
    );
  }
}

export default Users;