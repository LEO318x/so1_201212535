const { Router } = require("express");
const router = Router();

const testc = require("../controllers/testc")

router.get('/', async (req, res) =>{
    testc.tesq();
    res.status(200).json({msg: "Servidor Up", data: []})
});

module.exports = router;