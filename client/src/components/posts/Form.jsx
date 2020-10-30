import React from 'react';
import { Button, FormControl } from 'react-bootstrap';

const Form = (props) => {
  const { title, body } = props.post;
  const show = props.show ? 'block' : 'd-none';

  return(
    <form onSubmit={ props.onSubmitHandler } className={ show + " mb-5"}>
      <FormControl
        name="title"
        value={ title }
        onChange={ props.onChangeHandler }
        placeholder="Add a suitable title for your content..."
        className="my-1"
      />
      
      <FormControl
        as="textarea"
        rows={ 5 }
        name="body"
        value={ body }
        onChange={ props.onChangeHandler }
        placeholder="Write your content..."
        className="my-1"
      />

      <Button type="submit" variant="primary" className="my-1 mr-2">post</Button>
      <Button variant="secondary" className="my-1 ml-2" onClick={ props.hideHandler }>cancel</Button>
    </form>
  );
}

export default Form;