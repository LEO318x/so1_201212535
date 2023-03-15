const { Router } = require("express");
const router = Router();

const ram = require("../controllers/ram.controller")

router.get('/ram', async (req, res) =>{
    const result = await ram.raminfo()
    res.status(200).json({msg: "Info RAM", data: result})
});

module.exports = router;