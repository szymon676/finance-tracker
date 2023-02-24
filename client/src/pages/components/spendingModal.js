import React, { useState } from "react";

export default function Modal() {
  const [modal, setModal] = useState(false);

  const toggleModal = () => {
    setModal(!modal);
  };

  function PostSpending() {
    const title = document.querySelector("#title").value;
    const price = document.querySelector("#price").value;
    const description = document.querySelector("#description").value;
    const data = { title: title, price: price, description: description };
    
    fetch("http://localhost:4000/spendings", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((data) => {
        location.reload()
        toggleModal(); // close the modal after successful post
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }

  return (
    <>
      {!modal && (
        
        <button onClick={toggleModal} className="btn-modal">
        Add spending
        </button>

      )}
      {modal && (
        <div className="modal">
          <div onClick={toggleModal} className="overlay"></div>
          <div className="modal-content">
            <h2>Add your spending</h2>
            <form>
              <label>
                <h3>Title</h3>
                <input type="text" name="name" id="title" />
              </label>

              <label>
                <h3>Price</h3>
                <input type="text" name="name" id="price" />
              </label>

              <label>
                <h3>Description</h3>
                <input type="text" name="name" id="description" />
              </label>
            </form>
            <hr />
            <button className="close-modal" onClick={PostSpending}>
              Add
            </button>
            <button className="close-modal" onClick={toggleModal}>
              Cancel
            </button>
          </div>
        </div>
      )}
    </>
  );
}
