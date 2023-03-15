const { Router } = require("express");
const router = Router();

const cpu = require("../controllers/cpu.controller")

router.get('/cpu', async (req, res) =>{
    const result = await cpu.cpuinfo()
    //console.log(result[0].procesos);
    res.status(200).json({msg: "Info CPU", data: result})
});

module.exports = router;