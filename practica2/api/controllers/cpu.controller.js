const conn = require("../connection/conn");

async function cpuinfo() {
    const result =
    await conn.promise().query('SELECT * FROM cpu ORDER BY id_cpu DESC LIMIT 1')
    return result[0];
}

module.exports = {cpuinfo}