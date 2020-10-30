import React from 'react';

const formatTime = (datetime) => {
  const time = new Date(datetime);

  const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric' };

  return time.toLocaleDateString("en-IN", options);
}

const Post = (props) => {
  const { ID, title, body, created_at } = props.post;
  const backgroundColor = (ID % 2) ? "" : "bg-light";

  return(
    <div className={ backgroundColor + " p-4 rounded my-4" }>
      <div className="section-title">
        <h2 className="h2"> { title } </h2>
      </div>
      
      <div className="section-body">
        <p dangerouslySetInnerHTML={{ __html: body }} />
      </div>
      
      <div className="small">
        <span className="float-right">{ formatTime(created_at) }</span>
      </div>
    </div>
  );
}

export default Post;