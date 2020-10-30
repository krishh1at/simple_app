import React from 'react';
import Axios from 'axios';
import { Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEdit } from '@fortawesome/free-solid-svg-icons';
import Form from './Form';
import Post from './Post';

class Posts extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      post: {},
      posts: [],
      postCount: 0,
      show: false
    }

    this.onChangeHandler = this.onChangeHandler.bind(this);
    this.onSubmitHandler = this.onSubmitHandler.bind(this);
    this.showHandler = this.showHandler.bind(this);
    this.hideHandler = this.hideHandler.bind(this);
  }

  componentDidMount() {
    Axios.get(`/api/trending_posts`)
    .then(resp => {
      console.log(resp.data);
      this.setState({ posts: resp.data, postCount: resp.data.length });
    })
    .catch(resp => {
      console.log(resp);
    });
  }

  showHandler() {
    this.setState({ show: true });
  }

  hideHandler() {
    this.setState({ show: false });
  }


  onChangeHandler = (e) => {
    var newPost = Object.assign({}, this.state.post, { [e.target.name]: e.target.value });
    this.setState({ post: newPost });
  }

  onSubmitHandler = (e) => {
    e.preventDefault();

    Axios.post(`/api/user/1/posts`, this.state.post)
    .then(resp => {
      const posts = this.state.posts;
      const postCount = posts.unshift(resp.data);
      this.setState({ posts: posts, postCount: postCount, post: { title: '', body: '' }, show: false });
    })
    .catch(resp => {
      console.log(resp);
    });
  }

  render() {
    const posts = this.state.posts.map(post => {
      return <Post post={ post } key={ post.ID } />
    });

    const show = this.state.show ? "d-none" : "";
    const hide = this.state.show ? "" : "d-none";

    return (
      <div className="section">
        <div className="section-title">
          <h1 className="h1">Trending Posts: { this.state.postCount }</h1>
          
          <h2 className="h2">
          <FontAwesomeIcon icon={ faEdit } className={ hide } /> Write your own post:
            
            <Button onClick={ this.showHandler } className={ show + " ml-2" }>
              <FontAwesomeIcon icon={ faEdit } />
            </Button>
          </h2>
        </div>
        
        <Form
          onChangeHandler={ this.onChangeHandler }
          onSubmitHandler={ this.onSubmitHandler }
          hideHandler={ this.hideHandler }
          post={ this.state.post }
          show={ this.state.show }
        />
        
        <div className="section">
          { posts }
        </div>
      </div>
    );
  }
}

export default Posts;