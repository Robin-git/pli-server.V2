const mysql = require('mysql')

// LOCALHOST
// const con = mysql.createConnection({
//     host: "localhost",
//     user: "root",
//     password: ""
//     database: "gloo_rec"
// })

// AZURE

function createConnection() {
    const connection = mysql.createConnection({
        host: "localhost",
        user: "root",
        password: "",
        port: 3306,
        database: "gloo_dev"
    })
    connection.connect((err) => {
        if (err) throw err
        console.log("Connected to BDD!")
    })
    return connection
}
const connection = createConnection()
module.exports = connection