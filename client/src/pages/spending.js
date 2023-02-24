import Modal from "./components/spendingModal";
import { useEffect, useState } from "react";
import axios from "axios";

const app = () => {
  const [spendings, setSpendings] = useState([]);

    useEffect(() => {
      axios
        .get("http://localhost:4000/spendings")
        .then((response) => {
          console.log(response.data); // check the data being returned by the API
          setSpendings(response.data["spendings:"]); // update the state with the incomes array
        })
        .catch((error) => console.log(error));
    }, []);
  
    
    const deleteSpending = (id) => {
      axios
        .delete(`http://localhost:4000/spendings/${id}`)
        .then(() => {
          console.log("succesfully deleted spending :D"); 
          setSpendings(prevSpendings => prevSpendings.filter(spending => spending.ID !== id));
        })
        .catch((error) => console.log(error));
    }

    return (
      <main>
      <div class="header">
          <h1><a href="/">Finance tracker app</a></h1>
          <h3>
            <a href='dashboard'>Dashboard</a>
          </h3>
          <h3>
            <a href='income'>Income</a>
          </h3>
      </div>

      <div class="income">
        <article>
          <h2>Spendings:</h2>
          <div class="card">
            <h2><b>History:</b></h2>
            <div>
            {spendings && spendings.map((spending) => (
              <div key={spending.ID} class="div">
                <p>Title: {spending.title}</p>
                <p>Price: -{spending.price}</p>
                <p>Description: {spending.description}</p>
                <button class="button1" onClick={() => deleteSpending(spending.ID)}>Delete spending</button>
              </div>
            ))}
            </div>
          </div>
        </article>
        <Modal/>
        </div>
    </main>
    )
}
export default app;