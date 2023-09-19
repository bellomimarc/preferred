const { createClient } = require('redis')

const getConnectedClient = async () => {
    const client = createClient({
        url: 'redis://localhost:6666'
    });

    client.on('error', err => console.log('Redis Client Error', err));

    await client.connect();

    return client
}

module.exports = { getConnectedClient }