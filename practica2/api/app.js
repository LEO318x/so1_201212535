require('dotenv').config();
const express = require("express");
const morgan = require("morgan");
const cors = require("cors");

const app = express();
const port = process.env.PORT || 5000;

//Environment variables
require("dotenv").config({path: `.env.${process.env.NODE_ENV}`});

// Middlewares
app.use(morgan('dev'));
app.use(cors());
app.use(express.json({ limit: '50mb' }));

//Default Router
app.use('/', require('./routes/default.route'));

// Routes
app.use('/api/', require('./routes/ram.route'));
app.use('/api/', require('./routes/cpu.route'));
app.use('/api/', require('./routes/cpuram.route'));

// Iniciando servidor
app.listen(port, () => {
    console.log("Server nodejs in the port", port);
});

module.exports = app;