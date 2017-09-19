const https = require('https');

const KEY = `AIzaSyBUuXyU38XzoNvD3v8ax9gREhs9PT7mvIc`
const URL = `https://maps.googleapis.com/maps/api/place/textsearch/json?query=bar+in+Chartres&key=${KEY}`
https.get(URL, (res) => {
    res.on('data', (data) => {
        // const allBar = JSON.parse(data.toString())
        // let json = JSON.stringify(data)
        // let bars = bufferOriginal.toString('utf8')
        // console.log(data)
        let buf = Buffer.concat([data], data.length);
        let json = JSON.stringify(buf)
        console.log(json)
        // console.log(allBar['results'])
    })
})