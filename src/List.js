import React, { useState, useEffect } from 'react';

export default ({ getReviews }) => {
  const [reviews, setReviews] = useState({
    1: {
      data: "The place have a great atmosphere. \nThe service was good and kind. \nThe menu is not versatile enough for me. \nThe I took a Philadelphia sandwich and my friend took a 3-cheese burger, both were great!",
      author: "Nate"
    },
    2: {
      data: "American dinner place opened almost throughout the night. Great choice to eat after a long party or a dance. The service is good and the food is always tasty. The most recommended dish is their chicken-wings, covered with a batter based sweet chili sauce. These are also great sweet option here. There is plenty of parking places close by and children can also find their way here.",
      author: "Sergey"
    }
  });

  useEffect(() => {
    getReviews()
      .then(({ outputArguments }) => {
        const intermediate = JSON.parse(outputArguments[0].value);
        return Object.keys(intermediate).reduce((acc, id) => {
          acc[id] = JSON.parse(intermediate[id]);
          return acc;
        }, {...reviews});
      }).then(setReviews);
  }, [getReviews]);

  return (
    <>
      <h3>Reviews</h3>
      {Object.keys(reviews).map(id => <div key={id}>
        <h5>From: {reviews[id].author}</h5>
        <p>{reviews[id].data}</p>
      </div>)}
    </>
  )
};