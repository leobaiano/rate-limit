const axios = require('axios');

const url = "http://localhost:8080/hello";

let numRequests = 0;
let numErrors = 0;

async function sendRequest() {
    try {
        const response = await axios.get(url);

        if (response.status === 200) {
            numRequests++;
            console.log(`Request ${numRequests}: Success`);
            sendRequest();
        }
    } catch (error) {
        if (error.response.status === 429 && numErrors < 5) {
            numErrors++;
            numRequests++;
            console.log(`Request ${numRequests}: Error 429 - Too Many Requests`);

            sendRequest();
        } else {
            console.log("Maximum number of errors reached. Closing the client.");
        }
    }
}

sendRequest();