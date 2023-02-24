import {useState,useEffect} from 'react'
import axios from 'axios';

const deleteSpending = (id) => {
  axios.delete(`http://localhost:4000/${id}`)
    .then(() => alert('Deleted Spending'));
}

const dashboard = () => {
  const [incomes, setIncomes] = useState([]);

    useEffect(() => {
      axios
        .get("http://localhost:4000/incomes")
        .then((response) => {
          console.log(response.data); // check the data being returned by the API
          setIncomes(response.data["incomes:"]); // update the state with the incomes array
        })
        .catch((error) => console.log(error));
    }, []);

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
  return (
        <main>

        <div class="header">
            <h1><a href="/">Finance tracker app</a></h1>
            <h3>
              <a href='income'>Income</a>
            </h3>
            <h3>
              <a href='spending'>Spending</a>
            </h3>
        </div>

        <div class="income">
          <article>
            <h2>Income:</h2>
            <div class="card">
              <h2><b>History:</b></h2>
              <div>
            {incomes && incomes.map((income) => (
              <div key={income.ID} class="div">
                <p>Title: {income.title}</p>
                <p>Price: +{income.price}</p>
                <p>Description: {income.description}</p>
              </div>
            ))}
            </div>
            </div>
          </article>
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
              </div>
            ))}
            </div>
            </div>
          </article>
          
          </div>

    </main>
    )
}
export default dashboard;