const conn = require("../connection/conn");

async function raminfo() {
    const result =
    await conn.promise().query('SELECT * FROM ram ORDER BY id_ram DESC LIMIT 1')
    return result[0];
}

module.exports = {raminfo}