import { useEffect, useState } from "react";
import Modal from "./components/incomeModal";

const app = () => {
    const [incomes, setIncomes] = useState([]);

    useEffect(() => {
      fetch("http://localhost:4000/incomes")
        .then((response) => response.json())
        .then((data) => {
          setIncomes(data.incomes)
          console.log(data)
        })
        .catch((error) => console.log(error));
    }, []);

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
              {incomes && incomes.length > 0 ? (
                <ul>
                  {incomes.map((income) => (
                    <li key={income.id}>
                      {income.title}: {income.price}
                    </li>
                  ))}
                </ul>
              ) : (
                <p>Loading...</p>
              )}
            </div>
          </article>

          <Modal />
        </div>
      </main>
    );
};

export default app;
