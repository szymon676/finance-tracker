import { useEffect, useState } from "react";
import Modal from "./components/incomeModal";
import axios from "axios";

const App = () => {
  
  const [incomes, setIncomes] = useState([]);

    useEffect(() => {
      axios
        .get("http://localhost:4000/incomes")
        .then((response) => {
          console.log(response.data); 
          setIncomes(response.data["incomes:"]); 
        })
        .catch((error) => console.log(error));
    }, []);

    const deleteIncome = (id) => {
      axios
        .delete(`http://localhost:4000/incomes/${id}`)
        .then(() => {
          console.log("succesfully deleted income :D"); 
          setIncomes(prevIncomes => prevIncomes.filter(income => income.ID !== id));
        })
        .catch((error) => console.log(error));
    } 
    
    return (
      <main>
        <div className="header">
          <h1>
            <a href="/">Finance tracker app</a>
          </h1>
          <h3>
            <a href="dashboard">Dashboard</a>
          </h3>
          <h3>
            <a href="spending">Spending</a>
          </h3>
        </div>

        <div className="income">
          <article>
            <h2>Total income:</h2>
            <div className="card">
              <h2 id="history">
                <b>History:</b>
              </h2>
            <div>
            {incomes && incomes.map((income) => (
              <div key={income.ID} class="div">
                <p>Title: {income.title}</p>
                <p>Price: {income.price}</p>
                <p>Description: {income.description}</p>
                <button class="button1" onClick={() => deleteIncome(income.ID)}>Delete Income</button>
              </div>
            ))}
            </div>
          </div>
          </article>

          <Modal />
        </div>
      </main>
    );
};

export default App;
