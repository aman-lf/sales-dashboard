import express from 'express';

const router = express.Router();

/* GET home page. */
router.get('/', function (req, res) {
  res.json({ title: 'Sales Dashbaord App' }).status(200);
});

export default router;
