import React, { useState } from 'react';

export default ({ onSubmit }) => {
  const [review, setReview] = useState('');
  const submitHandler = async () => {
    await onSubmit({text: review})
    setReview('');
  };

  return (
    <>
      <h3>Write a review</h3>
      <div>
        <p>
          <label>Review: </label> <br />
          <textarea value={review} rows="30" cols="80" onChange={ev => setReview(ev.target.value)}></textarea>
        </p>
        <p>
          <button onClick={submitHandler}>Submit</button>
        </p>
      </div>
    </>
  )
};