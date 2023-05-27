const express = require('express');
const { exec } = require('child_process');

const app = express();

app.get('/', (req, res) => {
    // Execute the docker-compose up command
    exec('docker-compose up -d', (error, stdout, stderr) => {
        if (error) {
            console.error(`Error starting containers: ${error}`);
            res.status(500).send('Error starting containers');
            return;
        }
        console.log(`Containers started: ${stdout}`);
        res.send('Containers started successfully');
    });
});

app.listen(3000, () => {
    console.log('Server started on port 3000');
});
