export default class HttpRequestHandler {
  constructor() {
    this.headers = new Headers();

    this.headers.append("Content-Type", "application/json");
    this.headers.append("Accept", "application/json");
    this.headers.append("Access-Control-Allow-Origin", "*");
    this.headers.append("GET", "POST", "OPTIONS");
  }
  fetchLoanInfo = async () => {
    try{
        const response = await fetch(
            "http://localhost:5000/loandetails",
            {
              headers: this.headers
            });
          const data = await response.json();
          return data;
    } catch (error)
    {   
        console.log(error.message);
        return []
    }
  };
}
