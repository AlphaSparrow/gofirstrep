const express = require('express');
const app = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
    res.status(200).send("Welcome Bhaijan");
})

app.get('/', (req, res) => {
    res.status(200).json({message: "Hello Bhai tamne"});
})

app.post('/post', (req, res) => {
    let myJson = req.body;

    res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})

app.listen(port, () => {
    console.log(`Listening bhai tere ko on ${port}`)
})