import Modal from "./components/spendingModal";
const app = () => {
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
            <p id="history"></p>
          </div>
        </article>
        <Modal/>
        </div>
    </main>
    )
}
export default app;