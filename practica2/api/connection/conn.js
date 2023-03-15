const mysql = require("mysql2");

const { HOSTDB, USERDB, PASSDB, NAMEDB, PORTDB} = process.env;

const connection = mysql.createConnection({
    host: HOSTDB,
    user: USERDB,
    password: PASSDB,
    database: NAMEDB,
    port: PORTDB
});

module.exports = connection;
