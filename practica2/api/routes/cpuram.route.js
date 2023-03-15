const { Router } = require("express");
const router = Router();

const cpu = require("../controllers/cpu.controller")
const ram = require("../controllers/ram.controller")

router.get('/cpuram', async (req, res) =>{
    const result = await cpu.cpuinfo();
    const result2 = await ram.raminfo();
    //console.log(result[0].procesos);
    res.status(200).json({msg: "Info CPU y RAM", data: result, data2: result2})
});

module.exports = router;