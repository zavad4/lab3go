const http = require('../common/http.js');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listForums: () => client.get('/forums'),
        createUser: (name, topics) => client.post('/users', 
            {
                name: name, 
                topics: topics
            }
        )
    }
};

module.exports = { Client };