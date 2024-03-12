import path from 'path';
import morgan from 'morgan';
import dotenv from 'dotenv';
import express from 'express';

import indexRouter from './routes'

import './db';

// Initialize environment variables.
dotenv.config();

var app = express();

app.use(express.json());
app.use(morgan('tiny'));
app.use(express.urlencoded({ extended: false }));
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);

const port = process.env.PORT || 3000;
app.listen(port, '0.0.0.0', () => {
  console.log(`Server started at ${port}`);
});

export default app;
