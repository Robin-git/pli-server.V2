const mysql = require('./db.js')
const https = require('https')

const KEY = `&key=AIzaSyA8sitcs9Rb43DfwWt775u9Pz_F2gYFZas`
const DEFAULT = `query=bars+in+Ivry${KEY}`
const URL = `https://maps.googleapis.com/maps/api/place/textsearch/json?`
const DEFAULT_URL = `${URL}${DEFAULT}`

let nbFunc = 0

function recurseCall(start = DEFAULT_URL) {
    nbFunc++
    console.log("Call func : " + nbFunc)
    https.get(start, (res) => {
        res.setEncoding('utf8')
        let body = '';
        res.on('data', (data) => {
            body += data
        })
        res.on('end', () => {
            let encode = JSON.parse(JSON.stringify(body))
            let decode = JSON.parse(encode)
            if (decode['results'] && decode['status'] == 'OK') {
                let record = []
                for (bar of decode.results) {
                    let tmp = []
                    tmp.push(bar['name'])
                    tmp.push(bar['geometry']['location']['lat'])
                    tmp.push(bar['geometry']['location']['lng'])
                    // Quand on arrivent en ville 
                    let address = bar['formatted_address']
                    let res = address.split(',');
                    const street = res.slice(0, res.length - 2).toString()
                    let city = res[res.length - 2].split(' ');
                    tmp.push(city[1])
                    tmp.push(city[2])
                    tmp.push(street)
                    record.push(tmp)
                }

                // INSERT 
                var sql = `INSERT INTO etablishment (name, x, y, postal_code, city, street) VALUES ?`;
                mysql.query(sql, [record], (err, result) => {
                    if (err) console.log(err);
                });

                if (decode['next_page_token']) {
                    setTimeout(recurseCall, 5000, `${URL}pagetoken=${decode['next_page_token']}${KEY}`)
                } else {
                    console.log("FINISH")
                    process.exit()
                }
            } else if (decode['status'] == "OVER_QUERY_LIMIT") {
                console.log("OVER_QUERY_LIMIT")
                setTimeout(recurseCall, 5000, start)
            } else if (decode['status'] == "INVALID_REQUEST") {
                console.log(decode['status'])
                console.log("INVALID_REQUEST")
            } else {
                console.log("FINISH")
            }
        }).on('error', (e) => {
            console.error(e);
        })
    })
}

// Start script
recurseCall()