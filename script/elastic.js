const connection = require('./db')
const elasticsearch = require('elasticsearch')

var client = new elasticsearch.Client({
    host: 'localhost:9200',
    log: 'trace'
});

client.ping({
    // ping usually has a 3000ms timeout
    requestTimeout: 1000
}, function (error) {
    if (error) {
        console.trace('elasticsearch cluster is down!');
    } else {
        console.log('All is well');
    }
});

connection.query('select * from etablishment', function (error, results) {
    if (error) throw error;
    const etablishments = JSON.parse(JSON.stringify(results))
    for (etablishment of etablishments) {
        console.log(etablishment)
    }
});
