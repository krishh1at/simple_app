import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navigation from './components/layouts/Navigation';
import Users from './components/users/Users';
import User from './components/users/User';

function App() {
  var PICTURE_URL = '';

  return (
    <Router>
      <Navigation />

      <div className="main">
        <div className="about">
          <div className="container">
            <Switch>
              <Route exact path="/users" component={Users} />
              <Route exact path="/users/:id" component={User} />
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
