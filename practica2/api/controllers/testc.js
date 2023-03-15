const conn = require("../connection/conn");

function tesq(){
const query = conn.query(
    'SELECT * FROM ram',
    function(err, results, fields) {
        console.log(results); // results contains rows returned by server
      }
    );
}

module.exports = {tesq};