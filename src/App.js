import './App.css';
import React from 'react';
import {
  createAccount,
  Client,
  argString
} from 'orbs-client-sdk/dist/index.es';
import List from './List';
import NewReview from './NewReview';

function App() {
  const { publicKey, privateKey } = createAccount();
  const orbsClient = new Client('http://localhost:8080', 42, 'TEST_NET');

  const getReviews = () => {
    const query = orbsClient.createQuery(
      publicKey,
      'reviews',
      'getAll',
      []
    );
    return orbsClient.sendQuery(query);
  }

  const submitHandler = ({ text }) => {
    const [tx] = orbsClient.createTransaction(
      publicKey,
      privateKey,
      'reviews',
      'add',
      [argString(text)]
    );
    return orbsClient.sendTransaction(tx);
  };

  return (
    <>
      <h1>Transparent Reviews</h1>
      <div className="container">
        <article>
          <List getReviews={getReviews} />
        </article>
        <article>
          <NewReview onSubmit={submitHandler} />
        </article>
      </div>
    </>
  );
}

export default App;
